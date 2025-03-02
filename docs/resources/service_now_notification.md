---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dynatrace_service_now_notification Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
  
---

# dynatrace_service_now_notification (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `active` (Boolean) The configuration is enabled (`true`) or disabled (`false`)
- `incidents` (Boolean) Send incidents into ServiceNow ITSM
- `message` (String) The content of the ServiceNow description. You can use the following placeholders:  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas
- `name` (String) The name of the notification configuration
- `profile` (String) The ID of the associated alerting profile
- `username` (String) The username of the ServiceNow account.   Make sure that your user account has the `rest_service`, `web_request_admin`, and `x_dynat_ruxit.Integration` roles

### Optional

- `events` (Boolean) Send events into ServiceNow ITOM
- `instance` (String) The ServiceNow instance identifier. It refers to the first part of your own ServiceNow URL. This field is mutually exclusive with the **url** field. You can only use one of them
- `password` (String, Sensitive) The password to the ServiceNow account
- `url` (String) The URL of the on-premise ServiceNow installation. This field is mutually exclusive with the **instance** field. You can only use one of them

### Read-Only

- `id` (String) The ID of this resource.


