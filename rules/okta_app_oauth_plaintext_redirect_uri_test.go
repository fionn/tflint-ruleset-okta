package rules_test

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"

	"github.com/fionn/tflint-ruleset-okta/rules"
)

func TestOktaAppOauthPlaintextRedirectURI(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "Redirect URI one of one elements using HTTPS",
			Content: `
resource "okta_app_oauth" "example" {
  redirect_uris = ["https://example.com/"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "Redirect URI two of two elements using HTTPS",
			Content: `
resource "okta_app_oauth" "example" {
  redirect_uris = ["https://one.example.com/", "https://two.example.com/"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "Redirect URI one of one elements using HTTP",
			Content: `
resource "okta_app_oauth" "example" {
  redirect_uris = ["http://example.com/"]
}`,
			Expected: helper.Issues{
				{
					Rule:    rules.NewOktaAppOauthPlaintextRedirectURIRule(),
					Message: "Non-local redirect URI http://example.com/ should use TLS",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 42},
					},
				},
			},
		},
		{
			Name: "Redirect URI one of two elements using HTTP",
			Content: `
resource "okta_app_oauth" "example" {
  redirect_uris = ["https://example.com/", "http://example.com/"]
}`,
			Expected: helper.Issues{
				{
					Rule:    rules.NewOktaAppOauthPlaintextRedirectURIRule(),
					Message: "Non-local redirect URI http://example.com/ should use TLS",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 66},
					},
				},
			},
		},
		{
			Name: "Redirect URI one of two elements using HTTP",
			Content: `
resource "okta_app_oauth" "example" {
  redirect_uris = ["http://one.example.com/", "http://two.example.com/"]
}`,
			Expected: helper.Issues{
				{
					Rule:    rules.NewOktaAppOauthPlaintextRedirectURIRule(),
					Message: "Non-local redirect URI http://one.example.com/ should use TLS",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 73},
					},
				},
				{
					Rule:    rules.NewOktaAppOauthPlaintextRedirectURIRule(),
					Message: "Non-local redirect URI http://two.example.com/ should use TLS",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 73},
					},
				},
			},
		},
		{
			Name: "Redirect URI one of one element using HTTP locally",
			Content: `
resource "okta_app_oauth" "example" {
  redirect_uris = ["http://127.0.0.1/"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "Redirect URI one of one element using HTTP locally",
			Content: `
resource "okta_app_oauth" "example" {
			redirect_uris = ["http://127.0.0.1:5000/"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "Redirect URI one of one element using HTTPS locally",
			Content: `
resource "okta_app_oauth" "example" {
  redirect_uris = ["https://127.0.0.1/"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "Redirect URI one of one element using HTTP at localhost",
			Content: `
resource "okta_app_oauth" "example" {
  redirect_uris = ["http://localhost/"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "Redirect URI one of one element using HTTPS at localhost",
			Content: `
resource "okta_app_oauth" "example" {
  redirect_uris = ["https://localhost/"]
}`,
			Expected: helper.Issues{},
		},
	}

	rule := rules.NewOktaAppOauthPlaintextRedirectURIRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
