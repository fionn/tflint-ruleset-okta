package rules

import (
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
	"testing"
)

func TestOktaPolicyNameRuleOneChar(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "Length is 1",
			Content: `
resource "okta_auth_server_policy" "auditlogservice_read" {
  name = "f"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewOktaPolicyNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func TestOktaPolicyNameRule50Char(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "Length is 50",
			Content: `
resource "okta_auth_server_policy" "auditlogservice_read" {
  name = "Check ____________________________________ 50 char"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewOktaPolicyNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func TestOktaPolicyNameRuleZeroChar(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "Length is 0",
			Content: `
resource "okta_auth_server_policy" "auditlogservice_read" {
  name = ""
}`,
			Expected: helper.Issues{
				{
					Rule:    NewOktaPolicyNameRule(),
					Message: "Name must be from 1 to 50 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 12},
					},
				},
			},
		},
	}

	rule := NewOktaPolicyNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}

func TestOktaPolicyNameRule51Char(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "name must be from 1 to 50 characters",
			Content: `
resource "okta_auth_server_policy" "auditlogservice_read" {
  name = "Check _____________________________________ 51 char"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewOktaPolicyNameRule(),
					Message: "Name must be from 1 to 50 characters",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 63},
					},
				},
			},
		},
	}

	rule := NewOktaPolicyNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
