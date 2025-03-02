---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dynatrace_webhook_notification Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
  
---

# dynatrace_webhook_notification (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `active` (Boolean) The configuration is enabled (`true`) or disabled (`false`)
- `name` (String) The name of the notification configuration
- `payload` (String) The content of the notification message. You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas
- `profile` (String) The ID of the associated alerting profile
- `url` (String) The URL of the WebHook endpoint

### Optional

- `headers` (Block List, Max: 1) A list of the additional HTTP headers (see [below for nested schema](#nestedblock--headers))
- `insecure` (Boolean) Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates
- `notify_closed_problems` (Boolean) Send email if problem is closed
- `notify_event_merges` (Boolean) Call webhook if new events merge into existing problems

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--headers"></a>
### Nested Schema for `headers`

Required:

- `header` (Block Set, Min: 1) An additional HTTP Header to include when sending requests (see [below for nested schema](#nestedblock--headers--header))

<a id="nestedblock--headers--header"></a>
### Nested Schema for `headers.header`

Required:

- `name` (String) The name of the HTTP header

Optional:

- `secret_value` (String, Sensitive) The value of the HTTP header as a sensitive property. May contain an empty value. `secret_value` and `value` are mutually exclusive. Only one of those two is allowed to be specified.
- `value` (String) The value of the HTTP header. May contain an empty value. `secret_value` and `value` are mutually exclusive. Only one of those two is allowed to be specified.


