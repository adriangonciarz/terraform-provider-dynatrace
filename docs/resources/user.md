---
layout: ""
page_title: dynatrace_user Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_user` covers configuration for users
---

# dynatrace_user (Resource)

## Dynatrace Documentation

- User management and SSO - https://www.dynatrace.com/support/help/how-to-use-dynatrace/user-management-and-sso

- User management API - https://www.dynatrace.com/support/help/dynatrace-api/account-management-api/user-management-api

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) User's email address
- `first_name` (String) User's first name
- `last_name` (String) User's last name
- `user_name` (String) The User Name

### Optional

- `groups` (Set of String) List of user's user group IDs

### Read-Only

- `id` (String) The ID of this resource.
 