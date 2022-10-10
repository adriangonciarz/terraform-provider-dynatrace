---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dynatrace_pager_duty_notification Resource - terraform-provider-dynatrace"
subcategory: ""
description: |-
  
---

# dynatrace_pager_duty_notification (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account` (String) The name of the PagerDuty account
- `active` (Boolean) The configuration is enabled (`true`) or disabled (`false`)
- `name` (String) The name of the notification configuration
- `profile` (String) The ID of the associated alerting profile
- `service` (String) The name of the PagerDuty Service

### Optional

- `api_key` (String, Sensitive) The API key to access PagerDuty

### Read-Only

- `id` (String) The ID of this resource.

