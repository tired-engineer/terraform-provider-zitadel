---
page_title: "zitadel_human_user Resource - terraform-provider-zitadel"
subcategory: ""
description: |-
  Resource representing a human user situated under an organization, which then can be authorized through memberships or direct grants on other resources.
---

# zitadel_human_user (Resource)

**Caution: Email can only be set verified if a password is set for the user, either with initial_password or during runtime**

Resource representing a human user situated under an organization, which then can be authorized through memberships or direct grants on other resources.

## Example Usage

```terraform
resource zitadel_human_user human_user {
  depends_on = [zitadel_org.org]

  org_id             = zitadel_org.org.id
  user_name          = "human@localhost.com"
  first_name         = "firstname"
  last_name          = "lastname"
  nick_name          = "nickname"
  display_name       = "displayname"
  preferred_language = "de"
  gender             = "GENDER_MALE"
  phone              = "+41799999999"
  is_phone_verified  = "true"
  email              = "test@zitadel.com"
  is_email_verified  = "true"
  initial_password = "Password1!"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) Email of the user
- `first_name` (String) First name of the user
- `last_name` (String) Last name of the user
- `org_id` (String) ID of the organization
- `user_name` (String) Username

### Optional

- `display_name` (String) Display name of the user
- `gender` (String) Gender of the user
- `initial_password` (String, Sensitive) Initially set password for the user, not changeable after creation
- `is_email_verified` (Boolean) Is the email verified of the user, can only be true if password of the user is set
- `is_phone_verified` (Boolean) Is the phone verified of the user
- `nick_name` (String) Nick name of the user
- `phone` (String) Phone of the user
- `preferred_language` (String) Preferred language of the user

### Read-Only

- `id` (String) The ID of this resource.
- `login_names` (List of String) Loginnames
- `preferred_login_name` (String) Preferred login name
- `state` (String) State of the user