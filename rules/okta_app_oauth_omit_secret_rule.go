package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// OktaPolicyNameRule checks whether ...
type OktaAppOauthOmitSecretRule struct {
	tflint.DefaultRule
	resourceType  string
	attributeName string
	expected      bool
}

// NewOktaPolicyNameRule returns a new rule
func NewOktaAppOauthOmitSecretRule() *OktaAppOauthOmitSecretRule {
	return &OktaAppOauthOmitSecretRule{
		resourceType:  "okta_app_oauth",
		attributeName: "omit_secret",
		expected:      true,
	}
}

// Name returns the rule name
func (r *OktaAppOauthOmitSecretRule) Name() string {
	return "okta_app_oauth_omit_secret_rule"
}

// Enabled returns whether the rule is enabled by default
func (r *OktaAppOauthOmitSecretRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *OktaAppOauthOmitSecretRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *OktaAppOauthOmitSecretRule) Check(runner tflint.Runner) error {
	logger.Trace(fmt.Sprintf("Check %s rule", r.Name()))

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			err = runner.EmitIssue(r, "OAuth application secret should be omitted", resource.DefRange)
			if err != nil {
				return err
			}
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(omitSecret bool) error {
			if !omitSecret {
				err = runner.EmitIssue(r, "OAuth application secret should be omitted", attribute.Range)
				if err != nil {
					return err
				}
			}
			return nil
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
