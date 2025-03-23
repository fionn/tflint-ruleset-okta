package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_OktaAppOauthOmitSecret_True(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "Secret is omitted",
			Content: `
resource "okta_app_oauth" "example" {
  omit_secret = true
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewOktaAppOauthOmitSecretRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func Test_OktaAppOauthOmitSecret_False(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "Secret is not omitted",
			Content: `
resource "okta_app_oauth" "example" {
  omit_secret = false
}`,
			Expected: helper.Issues{
				{
					Rule:    NewOktaAppOauthOmitSecretRule(),
					Message: "OAuth application secret should be omitted",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 22},
					},
				},
			},
		},
	}

	rule := NewOktaAppOauthOmitSecretRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func Test_OktaAppOauthOmitSecret_Missing(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "Secret is implicitly not omitted",
			Content: `
resource "okta_app_oauth" "example" {
}`,
			Expected: helper.Issues{
				{
					Rule:    NewOktaAppOauthOmitSecretRule(),
					Message: "OAuth application secret should be omitted",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 36},
					},
				},
			},
		},
	}

	rule := NewOktaAppOauthOmitSecretRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
