package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/fionn/tflint-ruleset-okta/rules"
)

var version = "dev"

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "okta",
			Version: version,
			Rules: []tflint.Rule{
				rules.NewOktaPolicyNameRule(),
				rules.NewOktaAppOauthOmitSecretRule(),
				rules.NewOktaAppOauthPlaintextRedirectURIRule(),
				rules.NewOktaAppImplicitAuthenticationPolicyRule(),
			},
		},
	})
}
