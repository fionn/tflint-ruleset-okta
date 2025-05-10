package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type OktaAppImplicitAuthenticationPolicyRule struct {
	tflint.DefaultRule
	resourceTypes [2]string
	attributeName string
	expected      bool
}

// Create a rule.
func NewOktaAppImplicitAuthenticationPolicyRule() *OktaAppImplicitAuthenticationPolicyRule {
	return &OktaAppImplicitAuthenticationPolicyRule{
		resourceTypes: [2]string{"okta_app_oauth", "okta_app_saml"},
		attributeName: "authentication_policy",
		expected:      true,
	}
}

// Name returns the rule name
func (r *OktaAppImplicitAuthenticationPolicyRule) Name() string {
	return "okta_app_oauth_omit_secret"
}

// Enabled returns whether the rule is enabled by default
func (r *OktaAppImplicitAuthenticationPolicyRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (r *OktaAppImplicitAuthenticationPolicyRule) Severity() tflint.Severity {
	return tflint.NOTICE
}

func (r *OktaAppImplicitAuthenticationPolicyRule) Check(runner tflint.Runner) error {
	logger.Debug(fmt.Sprintf("checking %s rule", r.Name()))

	issueMessage := "Application implicitly uses the default authentication policy"

	for _, resourceType := range r.resourceTypes {

		resources, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
			Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
		}, nil)
		if err != nil {
			return err
		}

		for _, resource := range resources.Blocks {
			_, exists := resource.Body.Attributes[r.attributeName]
			if !exists {
				err = runner.EmitIssue(r, issueMessage, resource.DefRange)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
