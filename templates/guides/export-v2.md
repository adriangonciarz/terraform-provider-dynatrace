---
layout: ""
page_title: "Export Utility"
description: |-
  The export utility queries the Dynatrace Environment specified and fetches all currently supported configuration
---

## Export Utility

### Command Line Syntax
Invoking the export functionality requires
* The environment variable `DYNATRACE_ENV_URL` as the URL of your Dynatrace environment
* The environment variable `DYNATRACE_API_TOKEN` as the API Token of your Dynatrace environment
* Optionally the environment variable `DYNATRACE_TARGET_FOLDER`. If it's not set, the output folder `./configuration` is assumed

Windows: `terraform-provider-dynatrace.exe -export [-v] [-ref] [-id] [-migrate] [-exclude] [<resourcename>[=<id>]]`

Linux: `./terraform-provider-dynatrace -export [-v] [-ref] [-id] [-migrate] [-exclude] [<resourcename>[=<id>]]`
### Options
* `-v` Enable verbose logging
* `-ref` Enable resources with data sources and dependencies
* `-id` Enable commented id output in resource files
* `-migrate` Enable output specific to environment migration
    -  Removes node IDs from private synthetic locations
* `-exclude` Exclude specified resource(s) from export

**NOTE:** Dashboards (because there could be thousands of them) are currently excluded from the export unless the resource is directly specified in the command line arguments.

### Usage Examples
* `./terraform-provider-dynatrace -export` downloads all available configuration settings without data sources and dependency references (export similar to previous version)
* `./terraform-provider-dynatrace -export -ref -id` downloads all available configuration settings with data sources / dependency references and adds commented ids in resource output
* `./terraform-provider-dynatrace -export -ref dynatrace_dashboard dynatrace_web_application` downloads all available dashboards, web applications and resource dependencies with references
* `./terraform-provider-dynatrace -export -ref dynatrace_alerting=4f5942d4-3450-40a8-818f-c5faeb3563d0 dynatrace_alerting=9c4b75f1-9a64-4b44-a8e4-149154fd5325` downloads the alerting profiles with the ids `4f5942d4-3450-40a8-818f-c5faeb3563d0` and `9c4b75f1-9a64-4b44-a8e4-149154fd5325`, includes all resource dependencies with references
* `./terraform-provider-dynatrace -export -ref dynatrace_calculated_service_metric dynatrace_alerting=4f5942d4-3450-40a8-818f-c5faeb3563d0` downloads all available calculated service metrics and also the alerting profile with the id `4f5942d4-3450-40a8-818f-c5faeb3563d0`, includes all resource dependencies with references
* `./terraform-provider-dynatrace -export -ref -exclude dynatrace_calculated_service_metric dynatrace_alerting` download all available configuration settings except `dynatrace_calculated_service_metric` and `dynatrace_alerting`, includes all resource dependencies with references

### Additional Information
* There may be instances where the exported configuration is deprecated and/or is unable to be used for a create/update. In these instances, the files will be moved into `.flawed` of the output folder and the explanation will be available as a commented output in the resource file. 
    -  E.g. A dashboard with no tiles can be created and can be retrieved via the export, but the subsequent `terraform apply` would fail without any tiles. 
* There are instances where the returned configuration does not contain all of the required information to run an `terraform apply` due to sensitive data or  instances where the files require additional attention. The files that apply to this scenario will be automatically moved to `.requires_attention`, the explanation will be available as a commented output in the resource file.
    -  E.g. `dynatrace_credentials` confidential strings are not available via the API.
