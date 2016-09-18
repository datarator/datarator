Installation
============

Binaries 
--------

`Download the binary <https://github.com/datarator/datarator/releases>`_, that targets your platform.


From source
-----------

* Make sure to have `go installed <https://golang.org/dl/>`_ (datarator has been tested with version 1.7) and directory with go binaries in your PATH environment variable
* afterwards run the following commands::

    go get github.com/datarator/datarator
    cd $GOPATH/src/github.com/datarator/datarator
    go generate
    go install
