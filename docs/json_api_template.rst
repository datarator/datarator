Output template
===============

Following output templates are available:

* :ref:`csv`
* :ref:`sql`
* :ref:`xml`

.. _csv:

Template: csv
-------------

Enabled via: `"template":"csv"`.

Optional :ref:`payload` available:

* ``"header":"true"`` / ``"header":"false"`` - whether names of the colums should included (as the 1.st row) or not. By default is ``false``.
.. * ``"empty_value":"<empty value>"``- empty value. By default is empty string.
* ``"separator":"<separator>"`` - the separator string to be used for joining values.

For **example**, input JSON:
::

    {
        "template": "csv",
        "count": 3,
        "columns": [{
            "name": "name1",
            "type": "const",
            "payload": {
                "value": "value1"
            }
        }, {
            "name": "name2",
            "type": "const",
            "payload": {
                "value": "value2"
            }
        }, {
            "name": "name3",
            "type": "const",
            "payload": {
                "value": "value3"
            }
        }],
        "payload": {
            "header": true,
            "separator": ","
        }
    }


results in:
::

    	name1,name2,name3
    	value1,value2,value3
    	value1,value2,value3
    	value1,value2,value3

.. _sql:

Template: sql
-------------

Enabled via: ``"template":"sql"``.

For **example**, input JSON:
::

   {
    "template": "sql",
    "count": 3,
    "columns": [{
        "name": "name1",
        "type": "const",
        "payload": {
            "value": "value1"
        }
    }, {
        "name": "name2",
        "type": "const",
        "payload": {
            "value": "value2"
        }
    }, {
        "name": "name3",
        "type": "const",
        "payload": {
            "value": "value3"
        }
    }]
}

results in:
::

    INSERT INTO foo ( name1, name2, name3 ) VALUES ( 'value1', 'value2', 'value3' );
    INSERT INTO foo ( name1, name2, name3 ) VALUES ( 'value1', 'value2', 'value3' );
    INSERT INTO foo ( name1, name2, name3 ) VALUES ( 'value1', 'value2', 'value3' );

.. _xml:

Template: xml
-------------

Enabled via: ``"template":"xml"``.

Optional :ref:`payload` available:

* ``"pretty_print":"true"`` / ``"pretty_print":"false"`` - whether pretty printing should be enabled or not. By default is ``false``.
* ``"pretty_print_tabs":"true"`` / ``"pretty_print_tabs":"false"`` - whether to use tabs (or spaces) for pretty print. By default is ``false`` (=> uses spaces).
* ``"pretty_print_spaces_count":<count>``- the count of spaces in case of pretty print enabled. By default is 4.

For **example**, input JSON:
::

    {
        "template": "xml",
        "count": 3,
        "columns": [
            {
                "name": "name1",
                "type": "const",
                "payload": {
                    "value": ""
                },
                "columns": [
                    {
                        "name": "name2",
                        "type": "const",
                        "payload": {
                            "value": "value2",
                            "xml": "attribute"
                        }
                    },
                    {
                        "name": "name3",
                        "type": "const",
                        "payload": {
                            "value": "value3"
                        }
                    }
                ]
            }
        ],
        "payload": {
            "pretty_print": true
        }
    }

results in:
::

    <name1 name2="value2">
        <name3>value3</name3>
    </name1>
    <name1 name2="value2">
        <name3>value3</name3>
    </name1>
    <name1 name2="value2">
        <name3>value3</name3>
    </name1>
