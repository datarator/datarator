#!/bin/bash
set -e

# don't rerun
[ -d dist ] && { echo "dist/ dir already exists => no re-run supported!"; exit 1; }

#####################
# build cross-arch binaries (using gox)
#####################
go get github.com/mitchellh/gox
gox -output="dist/{{.Dir}}-${VERSION}-{{.OS}}_{{.Arch}}" 

pushd dist

####################
# generate man page
####################
mkdir -p usr/share/man/man1/
./datarator-${VERSION}-linux_amd64 -m | gzip -vc > usr/share/man/man1/datarator.1.gz

#####################
# build linux distro / osx packages (using fpm)
#####################
FPM_ARGS="-s dir -n datarator -v ${VERSION} --prefix '/' --license MIT --vendor 'http://github.com/datarator/datarator' --rpm-summary 'Stateless data generator with HTTP based JSON API' --description 'Stateless data generator with HTTP based JSON API' -m 'Peter Butkovic <butkovic@gmail.com>' --url 'https://github.com/datarator/datarator'"
gem install --no-ri --no-rdoc fpm

mkdir -p usr/local/bin/

# Linux

# 32 bit
cp -f datarator-${VERSION}-linux_386 usr/local/bin/datarator 
for target in deb rpm apk pacman
do 
    echo fpm -t ${target} "${FPM_ARGS}" -a 386 usr/ | /bin/bash
done
rm -rf usr/local/bin/* 

# 64 bit
cp -f datarator-${VERSION}-linux_amd64 usr/local/bin/datarator 
for target in deb rpm apk pacman
do 
    echo fpm -t ${target} "${FPM_ARGS}" -a amd64 usr/ | /bin/bash
done
rm -rf usr/local/bin/*

# prevent deployment problems afterwards, as we deploy `dist/*`` and `dist/usr` can't be deployed (see: https://travis-ci.org/datarator/datarator/builds/166543614)
rm -rf usr

# TODO
# # MacOS

# # 32 bit
# cp -f datarator-${VERSION}-darwin_386 usr/local/bin/datarator 
# for target in osxpkg
# do 
#     echo fpm -t ${target} "${FPM_ARGS}" -a 386 usr/ | /bin/bash
# done
# rm -rf usr/local/bin/*

# # 64 bit
# cp -f datarator-${VERSION}-darwin_amd64 usr/local/bin/datarator 
# for target in osxpkg
# do 
#     echo fpm -t ${target} "${FPM_ARGS}" -a amd64 usr/ | /bin/bash
# done
# rm -rf usr

#####################
# pack binaries using (g)zip
#####################
# non-windows
for dir in `ls -F | grep -v "/" | grep -v ".exe" | grep -v ".rpm" | grep -v ".deb" | grep -v ".apk" | grep -v ".pkg" | grep -v ".pkg.tar.xz"`
do 
    mv ${dir} datarator
    tar -cvzf ${dir}.tgz datarator
done

# windows
for dir in `ls -F | grep -v "/" | grep ".exe"`
do 
    mv ${dir} datarator.exe
    zip `echo ${dir} | sed 's/\.exe$//'`.zip datarator.exe
    rm -rf datarator.exe
done

popd

#####################
# sync github releases changelog with local one (using chandler)
#####################
gem install --no-ri --no-rdoc chandler

# setup credentials
cp -f travis/.netrc ~/.netrc
chmod 600 ~/.netrc
sed -i.bak "s/###OAUTH_LOGIN###/${OAUTH_LOGIN}/" ~/.netrc
sed -i.bak "s/###OAUTH_TOKEN###/${OAUTH_TOKEN}/" ~/.netrc
 
chandler --changelog=$(pwd)/docs/CHANGELOG.md --github=datarator/datarator push

#####################
# bintray 
#####################
# keep bintray.json in sync with current version
sed -i.bak "s/###VERSION###/${VERSION}/g" travis/bintray.json
