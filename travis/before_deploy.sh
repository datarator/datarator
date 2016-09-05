#!/bin/bash
set -e

# don't rerun
[ -d dist ] && exit 0

#####################
# build cross-arch binaries (using gox)
#####################
go get github.com/mitchellh/gox
gox -output="dist/{{.Dir}}-${VERSION}-{{.OS}}_{{.Arch}}" 

pushd dist

#####################
# build linux distro packages (using fpm)
#####################
FPM_ARGS="-s dir -n datarator -v ${VERSION} --prefix '/usr/local/bin' --license MIT --vendor 'http://github.com/datarator/datarator' --rpm-summary 'Stateless data generator with HTTP based JSON API' --description 'Stateless data generator with HTTP based JSON API' -m 'Peter Butkovic <butkovic@gmail.com>'"
gem install --no-ri --no-rdoc fpm

# 32 bit
cp datarator-${VERSION}-linux_386 datarator 
for target in deb rpm apk
do 
    echo fpm -t ${target} "${FPM_ARGS}" -a 386 datarator | /bin/bash
done
rm -f datarator 

# 64 bit
cp datarator-${VERSION}-linux_amd64 datarator 
for target in deb rpm apk
do 
    echo fpm -t ${target} "${FPM_ARGS}" -a amd64 datarator | /bin/bash
done
rm -f datarator

#####################
# pack binaries using (g)zip
#####################
# non-windows
for dir in `ls | grep -v ".exe" | grep -v ".rpm" | grep -v ".deb" | grep -v ".apk"`
do 
    mv ${dir} datarator
    tar -cvzf ${dir}.tgz datarator
done

# windows
for dir in `ls | grep ".exe"`
do 
    mv ${dir} datarator.exe
    zip `echo ${dir} | sed 's/\.exe$//'`.zip datarator.exe
    rm -rf datarator.exe
done

popd

#####################
# sync github releases changelog with local one (using chandler)
#####################

# chandler needs ruby version >= 2.2 => install it via rvm (while reuse travis installed rvm)
ruby -v 
rvm install 2.3.1 
rvm use --default 2.3.1 
ruby -v
/bin/bash --login
ruby -v

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
