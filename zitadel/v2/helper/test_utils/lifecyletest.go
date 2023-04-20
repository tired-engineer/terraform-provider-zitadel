package test_utils

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func RunLifecyleTest(
	t *testing.T,
	frame BaseTestFrame,
	resourceFunc func(string, string) string,
	initialProperty, updatedProperty,
	initialSecret, updatedSecret string,
	checkRemoteProperty func(expect string) resource.TestCheckFunc,
	checkDestroy, checkImportState resource.TestCheckFunc,
	importStateIdFunc resource.ImportStateIdFunc,
	wrongImportID,
	secretAttribute string,
) {
	var importStateVerifyIgnore []string
	initialConfig := fmt.Sprintf("%s\n%s", frame.ProviderSnippet, resourceFunc(initialProperty, initialSecret))
	updatedNameConfig := fmt.Sprintf("%s\n%s", frame.ProviderSnippet, resourceFunc(updatedProperty, initialSecret))
	updatedSecretConfig := fmt.Sprintf("%s\n%s", frame.ProviderSnippet, resourceFunc(updatedProperty, updatedSecret))
	steps := []resource.TestStep{
		{ // Check first plan has a diff
			Config:             initialConfig,
			ExpectNonEmptyPlan: true,
			// ExpectNonEmptyPlan just works with PlanOnly set to true
			PlanOnly: true,
		}, { // Check resource is created
			Config: initialConfig,
			Check: resource.ComposeAggregateTestCheckFunc(
				checkRemoteProperty(initialProperty),
				CheckStateHasIDSet(frame),
			),
		}, { // Check updating name has a diff
			Config:             updatedNameConfig,
			ExpectNonEmptyPlan: true,
			// ExpectNonEmptyPlan just works with PlanOnly set to true
			PlanOnly: true,
		}, { // Check remote state can be updated
			Config: updatedNameConfig,
			Check:  checkRemoteProperty(updatedProperty),
		},
	}
	if secretAttribute != "" {
		steps = append(steps, resource.TestStep{ // Check that secret has a diff
			Config:             updatedSecretConfig,
			ExpectNonEmptyPlan: true,
			// ExpectNonEmptyPlan just works with PlanOnly set to true
			PlanOnly: true,
		}, resource.TestStep{ // Check secret can be updated
			Config: updatedSecretConfig,
		})
		importStateVerifyIgnore = []string{secretAttribute}
	}
	if wrongImportID != "" {
		steps = append(steps, resource.TestStep{ // Expect import error if secret is not given
			ResourceName:  frame.TerraformName,
			ImportState:   true,
			ImportStateId: wrongImportID,
			ExpectError:   regexp.MustCompile(wrongImportID),
		})
	}
	if checkImportState != nil {
		steps = append(steps, resource.TestStep{ // Expect importing works
			ResourceName:            frame.TerraformName,
			ImportState:             true,
			ImportStateIdFunc:       importStateIdFunc,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: importStateVerifyIgnore,
			Check:                   checkImportState,
		})
	}
	resource.Test(t, resource.TestCase{
		ProviderFactories: ZitadelProviderFactories(frame.ConfiguredProvider),
		CheckDestroy:      checkDestroy,
		Steps:             steps,
	})
}
