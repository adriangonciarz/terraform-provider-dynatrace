---
layout: ""
page_title: dynatrace_http_monitor_performance Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_http_monitor_performance` covers configuration for HTTP monitor performance thresholds
---

# dynatrace_http_monitor_performance (Resource)

## Dynatrace Documentation

- Performance thresholds - https://www.dynatrace.com/support/help/platform-modules/digital-experience/synthetic-monitoring/http-monitors/configure-http-monitors#performance-thresholds

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:synthetic.http.performance-thresholds`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_http_monitor_performance` downloads all existing HTTP monitor performance thresholds configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_http_monitor_performance" "#name#" {
  enabled = true
  scope   = "HTTP_CHECK-1234567890000000"
  thresholds {
    threshold {
      event     = "HTTP_CHECK-1234567890000000"
      threshold = 10
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)
- `scope` (String) The scope of this setting (HTTP_CHECK)

### Optional

- `thresholds` (Block List, Max: 1) Performance thresholds (see [below for nested schema](#nestedblock--thresholds))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--thresholds"></a>
### Nested Schema for `thresholds`

Required:

- `threshold` (Block Set, Min: 1) (see [below for nested schema](#nestedblock--thresholds--threshold))

<a id="nestedblock--thresholds--threshold"></a>
### Nested Schema for `thresholds.threshold`

Required:

- `event` (String) Request
- `threshold` (Number) Threshold (in seconds)
 