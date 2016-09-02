Usage
=====

First make sure to have datarator installed (see page: :doc:`installation`) and running.

To generate::

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

