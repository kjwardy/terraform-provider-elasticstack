---
subcategory: "Ingest"
layout: ""
page_title: "Elasticstack: elasticstack_elasticsearch_ingest_processor_date Data Source"
description: |-
  Helper data source to create a processor which parses dates from fields, and then uses the date or timestamp as the timestamp for the document.
---

# Data Source: elasticstack_elasticsearch_ingest_processor_date

Helper data source to which can be used to parse dates from fields, and then uses the date or timestamp as the timestamp for the document. 
By default, the date processor adds the parsed date as a new field called `@timestamp`. You can specify a different field by setting the `target_field` configuration parameter. Multiple date formats are supported as part of the same date processor definition. They will be used sequentially to attempt parsing the date field, in the same order they were defined as part of the processor definition.

See: https://www.elastic.co/guide/en/elasticsearch/reference/current/date-processor.html

## Example Usage

Here is an example that adds the parsed date to the `timestamp` field based on the `initial_date` field:

```terraform
provider "elasticstack" {
  elasticsearch {}
}

data "elasticstack_elasticsearch_ingest_processor_date" "date" {
  field        = "initial_date"
  target_field = "timestamp"
  formats      = ["dd/MM/yyyy HH:mm:ss"]
  timezone     = "Europe/Amsterdam"
}

resource "elasticstack_elasticsearch_ingest_pipeline" "my_ingest_pipeline" {
  name = "date-ingest"

  processors = [
    data.elasticstack_elasticsearch_ingest_processor_date.date.json
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **field** (String) The field to get the date from.
- **formats** (List of String) An array of the expected date formats.

### Optional

- **description** (String) Description of the processor.
- **if** (String) Conditionally execute the processor
- **ignore_failure** (Boolean) Ignore failures for the processor.
- **locale** (String) The locale to use when parsing the date, relevant when parsing month names or week days.
- **on_failure** (List of String) Handle failures for the processor.
- **output_format** (String) The format to use when writing the date to `target_field`.
- **tag** (String) Identifier for the processor.
- **target_field** (String) The field that will hold the parsed date.
- **timezone** (String) The timezone to use when parsing the date.

### Read-Only

- **id** (String) Internal identifier of the resource
- **json** (String) JSON representation of this data source.