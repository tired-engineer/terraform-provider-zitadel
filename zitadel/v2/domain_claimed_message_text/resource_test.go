package domain_claimed_message_text_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/management"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper/test_utils"
)

func TestAccDomainClaimedMessageText(t *testing.T) {
	resourceName := "zitadel_domain_claimed_message_text"
	initialProperty := "initialtitle"
	updatedProperty := "updatedtitle"
	language := "en"
	frame, err := test_utils.NewOrgTestFrame(resourceName)
	if err != nil {
		t.Fatalf("setting up test context failed: %v", err)
	}
	test_utils.RunLifecyleTest(
		t,
		frame.BaseTestFrame,
		func(configProperty, _ interface{}) string {
			return fmt.Sprintf(`
resource "%s" "%s" {
  org_id      = "%s"
  language    = "%s"

  title       = "%s"
  pre_header  = "pre_header example"
  subject     = "subject example"
  greeting    = "greeting example"
  text        = "text example"
  button_text = "button_text example"
  footer_text = "footer_text example"
}`, resourceName, frame.UniqueResourcesID, frame.OrgID, language, configProperty)
		},
		initialProperty, updatedProperty,
		"", "",
		checkRemoteProperty(frame, language),
		regexp.MustCompile(`^\d{18}_en$`),
		// When deleted, the default should be returned
		checkRemoteProperty(frame, language)("Zitadel - Domain has been claimed"),
		nil, nil, "", "",
	)
}

func checkRemoteProperty(frame *test_utils.OrgTestFrame, lang string) func(interface{}) resource.TestCheckFunc {
	return func(expect interface{}) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			remoteResource, err := frame.GetCustomDomainClaimedMessageText(frame, &management.GetCustomDomainClaimedMessageTextRequest{Language: lang})
			if err != nil {
				return err
			}
			actual := remoteResource.GetCustomText().GetTitle()
			if actual != expect {
				return fmt.Errorf("expected %s, but got %s", expect, actual)
			}
			return nil
		}
	}
}