package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// OktaPolicyNameRule checks whether ...
type OktaPolicyNameRule struct {
	tflint.DefaultRule
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewOktaPolicyNameRule returns a new rule
func NewOktaPolicyNameRule() *OktaPolicyNameRule {
	return &OktaPolicyNameRule{
		resourceType:  "okta_auth_server_policy",
		attributeName: "name",
		max:           50,
		min:           1,
	}
}

// Name returns the rule name
func (r *OktaPolicyNameRule) Name() string {
	return "okta_auth_server_policy_name_length"
}

// Enabled returns whether the rule is enabled by default
func (r *OktaPolicyNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *OktaPolicyNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

func (r *OktaPolicyNameRule) Check(runner tflint.Runner) error {
	logger.Debug(fmt.Sprintf("checking %s rule", r.Name()))

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(policyName string) error {
			if len(policyName) > r.max || len(policyName) < r.min {
				err = runner.EmitIssue(r, "Name must be from 1 to 50 characters", attribute.Range)
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
