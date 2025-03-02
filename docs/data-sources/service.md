---
layout: ""
page_title: "dynatrace_service Data Source - terraform-provider-dynatrace"
description: |-
  The data source `dynatrace_service` covers queries for the ID of a service based on name and tags / tag-value pairs
---

# dynatrace_service (Data Source)

The service data source allows the service ID to be retrieved by its name and optionally tags / tag-value pairs.

- `name` queries for all services with the specified name
- `tags` (optional) refers to the tags that need to be present for the service (inclusive)

If multiple services match the given criteria, the first result will be retrieved.

## Example Usage

```terraform
data "dynatrace_service" "Test" {
  name = "Example"
  tags = ["TerraformKeyTest","TerraformKeyValueTest=TestValue"]
}

resource "dynatrace_key_requests" "#name#" {
  service = data.dynatrace_service.Test.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `tags` (Set of String) Required tags of the service to find

### Read-Only

- `id` (String) The ID of this resource.