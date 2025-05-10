package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_OktaAppImplicitAuthenticationPolicy_Present(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "OAuth application authentication policy is specified",
			Content: `
resource "okta_app_oauth" "example" {
  authentication_policy = "yolo"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "SAML application authentication policy is specified",
			Content: `
resource "okta_app_saml" "example" {
  authentication_policy = "yolo"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewOktaAppImplicitAuthenticationPolicyRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func Test_OktaAppImplicitAuthenticationPolicy_Missing(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "Authentication policy is omitted",
			Content: `
resource "okta_app_oauth" "example" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewOktaAppImplicitAuthenticationPolicyRule(),
					Message: "Application implicitly uses the default authentication policy",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 36},
					},
				},
			},
		},
		{
			Name: "SAML application authentication policy is omitted",
			Content: `
resource "okta_app_saml" "example" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewOktaAppImplicitAuthenticationPolicyRule(),
					Message: "Application implicitly uses the default authentication policy",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 35},
					},
				},
			},
		},
	}

	rule := NewOktaAppImplicitAuthenticationPolicyRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
