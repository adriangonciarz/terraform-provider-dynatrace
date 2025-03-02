---
layout: ""
page_title: dynatrace_span_attribute Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_span_attribute` covers configuration for span attributes
---

# dynatrace_span_attribute (Resource)

## Dynatrace Documentation

- Span settings - https://www.dynatrace.com/support/help/extend-dynatrace/extend-tracing/span-settings

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:span-attribute`)

## Export Example Usage

- `terraform-provider-dynatrace export dynatrace_span_attribute` downloads all existing span attribute configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs#exporting-existing-configuration-from-a-dynatrace-environment).

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `key` (String) the key of the attribute to capture
- `masking` (String) granular control over the visibility of attribute values

### Optional

- `persistent` (Boolean) Prevents the Span Attribute from getting deleted when running `terraform destroy` - to be used for Span Attributes that are defined by default on every Dynatrace environment.

### Read-Only

- `id` (String) The ID of this resource.
 