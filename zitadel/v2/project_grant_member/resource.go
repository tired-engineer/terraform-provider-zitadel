package project_grant_member

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper"
)

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing the membership of a user on an granted project, defined with the given role.",
		Schema: map[string]*schema.Schema{
			helper.OrgIDVar: helper.OrgIDResourceField,
			ProjectIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the project",
				ForceNew:    true,
			},
			GrantIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the grant",
				ForceNew:    true,
			},
			UserIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the user",
				ForceNew:    true,
			},
			rolesVar: {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "List of roles granted",
			},
		},
		DeleteContext: delete,
		CreateContext: create,
		UpdateContext: update,
		ReadContext:   read,
		Importer: &schema.ResourceImporter{StateContext: helper.ImportWithIDAndOptionalOrgV5(
			helper.ResourceIDVar,
			helper.ImportAttribute{
				Key:             ProjectIDVar,
				ValueFromString: helper.ConvertID,
			},
			helper.ImportAttribute{
				Key:             GrantIDVar,
				ValueFromString: helper.ConvertID,
			},
			helper.ImportAttribute{
				Key:             UserIDVar,
				ValueFromString: helper.ConvertID,
			},
		)},
	}
}
