---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dynatrace_email_notification Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
  
---

# dynatrace_email_notification (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `body` (String) The template of the email notification.  You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsHTML}`: All problem event details, including root cause, as an HTML-formatted string.  * `{ProblemDetailsJSON}`: All problem event details, including root cause, as a JSON object.  * `{ProblemDetailsMarkdown}`: All problem event details, including root cause, as a [Markdown-formatted](https://www.markdownguide.org/cheat-sheet/) string.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://www.dynatrace.com/support/help/shortlink/impact-analysis) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://www.dynatrace.com/support/help/shortlink/event-types) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas
- `name` (String) The name of the notification configuration
- `profile` (String) The ID of the associated alerting profile
- `subject` (String) The subject of the email notifications
- `to` (Set of String) The list of the email recipients

### Optional

- `active` (Boolean) The configuration is enabled (`true`) or disabled (`false`)
- `bcc` (Set of String) The list of the email BCC-recipients
- `cc` (Set of String) The list of the email CC-recipients
- `notify_closed_problems` (Boolean) Send email if problem is closed

### Read-Only

- `id` (String) The ID of this resource.


