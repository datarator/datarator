Column
======

Syntax:
::

	{
		"name": "<name>",
		"type": "<type>",
		"payload": <payload>
	}

Legend:

* `<name>` - name of the column
* `<type>` - type of the column
* `<payload>` - :ref:`payload`

Following column types are available:

* address:
    * :ref:`address.continent`
    * :ref:`address.country`
    * :ref:`address.city`
    * :ref:`address.phone`
    * :ref:`address.state`
    * :ref:`address.street`
    * :ref:`address.zip`
* color:
    * :ref:`color`
    * :ref:`color.hex`
* :ref:`const`
* :ref:`copy`
* credit card:
    * :ref:`credit_card.number`
    * :ref:`credit_card.type`
* currency:
    * :ref:`currency`
    * :ref:`currency.code`
* date:
    * :ref:`date.day.of_week`
    * :ref:`date.day.of_week.name`
    * :ref:`date.day.of_month`
    * :ref:`date.month`
    * :ref:`date.month.name`
    * :ref:`date.year`
    * :ref:`date.of_birth`
* :ref:`join`
* name:
    * :ref:`name.first`
    * :ref:`name.first.female`
    * :ref:`name.first.male`
    * :ref:`name.full`
    * :ref:`name.full.female`
    * :ref:`name.full.male`
    * :ref:`name.last`
    * :ref:`name.last.female`
    * :ref:`name.last.male`
* :ref:`regex`
* :ref:`row_index`

.. _address.continent:

Column: address.continent
-------------------------

Generates the random continent name.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "address.continent"
    }]

could result in value:
::

    	Europe

.. _address.country:

Column: address.country
-------------------------

Generates the random country name.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "address.country"
    }]

could result in value:
::

    	Slovakia

.. _address.city:

Column: address.city
-------------------------

Generates the random city name.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "address.city"
    }]

could result in value:
::

    	London

.. _address.phone:

Column: address.phone
-------------------------

Generates the random phone number.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "address.phone"
    }]

could result in value:
::

    	3-456-437-63-83

.. _address.state:

Column: address.state
-------------------------

Generates the random state name.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "address.state"
    }]

could result in value:
::

    	North Carolina

.. _address.street:

Column: address.street
-------------------------

Generates the random street name.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "address.street"
    }]

could result in value:
::

    	Eagle Crest Drive

.. _address.zip:

Column: address.zip
-------------------------

Generates the random zip name.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "address.zip"
    }]

could result in value:
::

    	9393157

.. _color:

Column: color
-------------------------

Generates the random color name.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "color"
    }]

could result in value:
::

    	Green

.. _color.hex:

Column: color.hex
-------------------------

Generates the random hexadecimal value of the color.

Optional :ref:`payload` available:

* ``"short":"true"`` / ``"short":"false"`` - whether short version of the hexadecimal value should be generated or not. By default is ``false``.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "color.hex",
        "payload": {
            "short": true
        }
    }]

could result in value:
::

    	390

.. _const:

Column: const
-------------------------

Generates constant value provided in payload.

Mandatory :ref:`payload` available:

* ``"value": <value>`` - the constant value to generate

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "const",
        "payload": {
            "value": "foo"
        }
    }]

results in value:
::

    	foo

.. _copy:

Column: copy
-------------------------

Generates the same value as the column referred.

Mandatory :ref:`payload` available:

* ``"from":"<column_name>"`` -  the column name whose value is to be copied.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "const",
        "options": {
            "value": "foo"
        }
    }, {
        "name": "name2",
        "type": "copy",
        "options": {
            "from": "name1"
        }
    }]

results (for columns: name1 as well as name2) in value:
::

    	foo

.. _credit_card.number:

Column: credit_card.number
-------------------------

Generates the random credit card number value.

Optional :ref:`payload` available:

* ``"type":"<column_name>"`` -  the type of credit card to generate number of. Valid values are: ``amex``, ``discover``, ``mastercard`` and ``visa``. 

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "credit_card.number",
        "payload": {
            "type": "amex"
        }
    }]

could result in value:
::

    	4771761587281649

.. _credit_card.type:

Column: credit_card.type
-------------------------

Generates the random credit card type value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "credit_card.type"
    }]

could result in value:
::

    	American Express

.. _currency:

Column: currency
-------------------------

Generates the random currency value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "currency"
    }]

could result in value:
::

    	New Zealand Dollars

.. _currency.code:

Column: currency.code
-------------------------

Generates the random currency code value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "currency.code"
    }]

could result in value:
::

    	GBP

.. _date.day.of_week:

Column: date.day.of_week
-------------------------

Generates the random weekday number value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "date.day.of_week"
    }]

could result in value:
::

    	2

.. _date.day.of_week.name:

Column: date.day.of_week
-------------------------

Generates the random weekday name value.

Optional :ref:`payload` available:

* ``"short":"true"`` / ``"short":"false"`` - whether short version of the weekday name should be generated or not. By default is ``false``.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "date.day.of_week.name",
        "payload": {
            "short": true
        }
    }]

could result in value:
::

    	Thu

.. _date.day.of_month:

Column: date.day.of_month
-------------------------

Generates the random day of month value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "date.day.of_month"
    }]

could result in value:
::

    	21

.. _date.month:

Column: date.month
-------------------------

Generates the random month value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "date.month"
    }]

could result in value:
::

    	11

.. _date.month.name:

Column: date.month.name
-------------------------

Generates the random month name value.

Optional :ref:`payload` available:

* ``"short":"true"`` / ``"short":"false"`` - whether short version of the month name should be generated or not. By default is ``false``.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "date.month.name",
        "payload": {
            "short": true
        }
    }]

could result in value:
::

    	Aug

.. _date.year:

Column: date.year
-------------------------

Generates the random year value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "date.year"
    }]

could result in value:
::

    	1448

.. _date.of_birth:

Column: date.of_birth
-------------------------

Generates the random date of birth value.

Optional :ref:`payload` available:

* ``"age":<age>`` - the age that date of birth should be generated for. If not specified, random age in interval 0-120 is used.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "date.of_birth",
        payload {
            "age": 18
        }
    }]

could result in value:
::

    	1998-02-22 22:08:28 +0100 CE

.. _join:

Column: join
-------------------------

Joins nested column values with the separator (optionaly) provided.

Optional :ref:`payload` available:

* ``"separator":<separator>`` -  the separator string to be used for joining values.

For **example** (without separator), input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "join",
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
        }]
    }]

would result in value:
::

    	value1value2

For **example** (with separator), input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "join",
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
        }]
        }, "payload": {
            "separator": ", "
    }]

would result in value:
::

    	value1,value2

.. _name.first:

Column: name.first
-------------------------

Generates the random first name value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "name.first"
    }]

could result in value:
::

    	Malcolm

.. _name.first.female:

Column: name.first.female
-------------------------

Generates the random female first name value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "name.first.female"
    }]

could result in value:
::

    	Sherly

.. _name.first.male:

Column: name.first.male
-------------------------

Generates the random male first name value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "name.first.male"
    }]

could result in value:
::

    	Brandon


.. _name.full:

Column: name.full
-------------------------

Generates the random full name value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "name.full"
    }]

could result in value:
::

    	Katrina Vanhamlin

.. _name.full.female:

Column: name.full.female
-------------------------

Generates the random female full name value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "name.full.female"
    }]

could result in value:
::

    	Katrina Vanhamlin

.. _name.full.male:

Column: name.full.male
-------------------------

Generates the random male full name value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "name.full.male"
    }]

could result in value:
::

    	Stephan Mciltrot


.. _name.last:

Column: name.last
-------------------------

Generates the random last name value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "name.last"
    }]

could result in value:
::

    	Vanhamlin


.. _name.last.female:

Column: name.last.female
-------------------------

Generates the random female last name value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "name.last.female"
    }]

could result in value:
::

    	Vanhamlin

.. _name.last.male:

Column: name.last.male
-------------------------

Generates the random male last name value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "name.last.male"
    }]

could result in value:
::

    	Mciltrot

.. _regex:

Column: regex
-------------------------

Generates the random string matching the specified regular expression (to examine full capabilities, refer to project: `lucasjones/reggen <https://github.com/lucasjones/reggen>`_ beeing used under the hood),

Mandatory :ref:`payload` available:

* ``"pattern":<pattern>`` - the pattern to match.

Optional :ref:`payload` available:

* ``"limit":<limit>`` - the maximum number of times `*`,`+` should repeat.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "regex",
        "payload": {
            "pattern": "z{1,3}",
            "limit": 10
        }
    }]

could result in value:
::

    	zzz

.. _row_index:

Column: row_index
-------------------------

Generates the current row index value.

For **example**, input JSON:
::

    "columns": [{
        "name": "name1",
        "type": "row_index"
    }]

results in values:
::

    0
    1
    2
    3
    ...
