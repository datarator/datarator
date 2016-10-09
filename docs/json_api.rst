JSON API
========
 ::

        {
                "template": "<template_name>",
                "emptyValue": "<empty_value>",
                "count": <count>,
                "columns": [ <column> , <column> , ...],
                "payload": <payload>
        }

Legend:

* ``<template_name>`` - :doc:`json_api_template` name
* ``<empty_value>`` - empty value. By default is empty string
* ``<count>`` - generated rows count
* ``<column>`` - see :doc:`json_api_column`
* ``<payload>`` - see :ref:`payload`

.. _payload:

Payload
-------

Holds :doc:`json_api_template` (if present in root node) or :doc:`json_api_column` specific options (if present in column node).

Syntax:
::

    	{
            "<name>":"<value>",
            "<name>":"<value>",
            ...
        }

Legend:

* ``<name>`` - name of the option
* ``<value>`` - value of the option

.. _payload_column_common:

:doc:`json_api_column` specific payload:
----------------------------------------------

All of these are optional:
::

        "emptyPercent": <empty_percent>,
        "xml": <attribute|element|cdata|comment|value>

Legend:

* ``<empty_percent>`` - indicates how much percent of the column values are to be empty. Valid values are: ``<0-100>``. Default value is ``0``.
* ``xml`` - considered for the :ref:`xml` only. Indicating the type of xml node to generate the column value to. (i.e.: ``xml: comment`` generates value wrapped in xml comment). Default value is ``element``.  