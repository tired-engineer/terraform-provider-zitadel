---
page_title: "zitadel_personal_access_token Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing a personal access token of a user
---

# zitadel_personal_access_token (Resource)

Resource representing a personal access token of a user

## Example Usage

```terraform
resource zitadel_personal_access_token pat {
  org_id          = zitadel_org.org.id
  user_id         = zitadel_machine_user.machine_user.id
  expiration_date = "2519-04-01T08:45:00Z"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `org_id` (String) ID of the organization
- `user_id` (String) ID of the user

### Optional

- `expiration_date` (String) Expiration date of the token in the RFC3339 format

### Read-Only

- `id` (String) The ID of this resource.
- `token` (String, Sensitive) Value of the token