package rules

import (
	"fmt"
	"net"
	"net/url"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func isNonLocalPlaintext(redirectURI string) (bool, error) {
	uri, err := url.Parse(redirectURI)
	if err != nil {
		return false, err
	}

	if uri.Scheme == "https" {
		return false, nil
	}

	if uri.Hostname() == "localhost" {
		return false, nil
	}

	ipAddress := net.ParseIP(uri.Hostname())
	if ipAddress != nil {
		return !ipAddress.IsLoopback(), nil
	}

	return true, nil
}

type OktaAppOauthPlaintextRedirectURIRule struct {
	tflint.DefaultRule
	resourceType  string
	attributeName string
}

func NewOktaAppOauthPlaintextRedirectURIRule() *OktaAppOauthPlaintextRedirectURIRule {
	return &OktaAppOauthPlaintextRedirectURIRule{
		resourceType:  "okta_app_oauth",
		attributeName: "redirect_uris",
	}
}

func (r *OktaAppOauthPlaintextRedirectURIRule) Name() string {
	return "okta_app_oauth_plaintext_redirect_uri"
}

func (r *OktaAppOauthPlaintextRedirectURIRule) Enabled() bool {
	return true
}

func (r *OktaAppOauthPlaintextRedirectURIRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *OktaAppOauthPlaintextRedirectURIRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(redirectURIs []string) error {
			for _, redirectURI := range redirectURIs {
				ruleViolation, err := isNonLocalPlaintext(redirectURI)
				if err != nil {
					return err
				}
				if ruleViolation {
					err = runner.EmitIssue(r, fmt.Sprintf("Non-local redirect URI %s should use TLS", redirectURI), attribute.Range)
					if err != nil {
						return err
					}
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
