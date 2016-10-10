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
    -c, --chunk=   The count of generated data returned in one chunk  (default: 1000)
    -e, --embed    Use embedded rather than external static resources
    -p, --port=    Port to listen on (default: 9292)
    -t, --timeout= Timeout in [ms] for maximum request processing (default: 3000)

    Help Options:
    -h, --help      Show this help message

Generate
--------

Asuming datarator has been started with default command line options, to generate::

    Hello world

Pick your tool of choice for sending HTTP POST requests:

`curl`_ way
--------
::

    curl -H 'Accept-Encoding: gzip, deflate' -H 'Content-Type: application/json' -X POST -d '{"template":"csv","count":1,"columns":[{"name":"greeting","type":"const", "payload":{"value":"Hello world!"}}]}' http://127.0.0.1:9292/api/schemas/say_hello

`wget`_ way
--------
::

    wget -qO - --header='Accept-Encoding: gzip, deflate' --header='Content-Type: application/json' --post-data '{"template":"csv","count":1,"columns":[{"name":"greeting","type":"const", "payload":{"value":"Hello world!"}}]}' http://127.0.0.1:9292/api/schemas/say_hello

`httpie`_ way
--------
::

    echo '{"template":"csv","count":1,"columns":[{"name":"greeting","type":"const", "payload":{"value":"Hello world!"}}]}' | http http://127.0.0.1:9292/api/schemas/say_hello Accept-Encoding:gzip,deflate

Refer to :doc:`json_api` page for full reference.

.. _curl: http://github.com/curl/curl
.. _wget: https://www.gnu.org/software/wget/
.. _httpie: https://github.com/jkbrzt/httpie