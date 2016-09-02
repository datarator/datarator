JSON API
========
 ::

        {
                "template": "<template_name>",
                "count": <count>,
                "columns": [ <column> , <column> , ...],
                "payload": <payload>
        }

Legend:

* ``<template_name>`` - name of the output template (see :doc:`json_api_template`)
* ``<count>`` - generated rows count
* ``<column>`` - column to generate (see :doc:`json_api_column`)
* ``<payload>`` - additional options for generation (see :ref:`payload`)

.. _payload:

Payload
-------

Holds :doc:`json_api_template` (if present in root node) or :doc:`json_api_column` specific options (if present in column node).

Syntax:
::

    	{"<name>":"<value>"}}

Legend:

* ``<name>`` - name of the option
* ``<value>`` - value of the option