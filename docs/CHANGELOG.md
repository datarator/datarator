# Change Log

All notable changes to this project will be documented in this file.

Datarator is in a pre-1.0 state. This means that its APIs and behavior are subject to breaking changes without deprecation notices. Until 1.0, version numbers will follow a [Semver][]-ish `0.y.z` format, where `y` is incremented when new features or breaking changes are introduced, and `z` is incremented for lesser changes or bug fixes.

## [Unreleased][]

## [0.2.1][] (2017-02-06)

Fixes:
* binaries release fixed

## [0.2.0][] (2016-10-10)

Known Issues:
* bianries release failed (no binaries provided)

Features:
* CLI:
    * options: confgurable chunk size
* API: 
    * column payload: `emptyPercent`
    * response gzip support
    * response header: `Content-Encoding`
    * timeout on data generation
    * removed api: GET /
    * template:`sql` removed whitespaces
    * json schema updated for usage with [jdorn/json-editor](http://github.com/jdorn/json-editor)

Fixes:
* using proper sql mime type

## 0.1.0 (2016-09-06)

* Initial release
 
[Semver]: http://semver.org
[Unreleased]: https://github.com/datarator/datarator/compare/v0.2.0...master
[0.2.1]: https://github.com/datarator/datarator/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/datarator/datarator/compare/0.1.0...v0.2.0