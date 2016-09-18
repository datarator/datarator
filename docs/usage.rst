Usage
=====

First make sure to have datarator installed (see page: :doc:`installation`).

Running
-------

Assuming datarator binary is in your path, simpy run::

    datarator

Command line options
--------------------

Following are available::

    Application Options:
    -e, --embed     Use embedded rather than external static resources
    -l, --parallel  Use parallel processing, rather than serial one.
    -p, --port=     Port to listen on (default: 9292)
    -t, --timeout=  Timeout in [ms] for maximum request processing (default: 3000)

    Help Options:
    -h, --help      Show this help message

Generate
--------

Asuming datarator has been started with default command line options, to generate::

    Hello world

Pick your tool of choice for sending HTTP POST requests:

curl way
--------
::

    curl -H 'Accept-Encoding: gzip, deflate' -H 'Content-Type: application/json' -X POST -d '{"template":"csv","count":1,"columns":[{"name":"greeting","type":"const", "payload":{"value":"Hello world!"}}]}' http://127.0.0.1:9292/api/schemas/say_hello

wget way
--------
::

    wget -qO - --header='Accept-Encoding: gzip, deflate' --header='Content-Type: application/json' --post-data '{"template":"csv","count":1,"columns":[{"name":"greeting","type":"const", "payload":{"value":"Hello world!"}}]}' http://127.0.0.1:9292/api/schemas/say_hello

Refer to :doc:`json_api` page for full reference.

