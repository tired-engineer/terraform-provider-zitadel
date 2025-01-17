---
page_title: "zitadel_machine_user Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.
---

# zitadel_machine_user (Resource)

Resource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.

## Example Usage

```terraform
resource "zitadel_machine_user" "default" {
  org_id      = data.zitadel_org.default.id
  user_name   = "machine@example.com"
  name        = "name"
  description = "a machine user"
  with_secret = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the machine user
- `user_name` (String) Username

### Optional

- `access_token_type` (String) Access token type, supported values: ACCESS_TOKEN_TYPE_BEARER, ACCESS_TOKEN_TYPE_JWT
- `description` (String) Description of the user
- `org_id` (String) ID of the organization
- `with_secret` (Boolean) Generate machine secret, only applicable if creation or change from false

### Read-Only

- `client_id` (String, Sensitive) Value of the client ID if withSecret is true
- `client_secret` (String, Sensitive) Value of the client secret if withSecret is true
- `id` (String) The ID of this resource.
- `login_names` (List of String) Loginnames
- `preferred_login_name` (String) Preferred login name
- `state` (String) State of the user

## Import

```terraform
# The resource can be imported using the ID format `<id:has_secret[:org_id][:client_id][:client_secret]>`, e.g.
terraform import machine_user.imported '123456789012345678:123456789012345678:true:my-machine-user:j76mh34CHVrGGoXPQOg80lch67FIxwc2qIXjBkZoB6oMbf31eGMkB6bvRyaPjR2t'
```
