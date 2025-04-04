package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-okta/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "okta",
			Version: "0.1.6",
			Rules: []tflint.Rule{
				rules.NewOktaPolicyNameRule(),
				rules.NewOktaAppOauthOmitSecretRule(),
				rules.NewOktaAppOauthPlaintextRedirectURIRule(),
			},
		},
	})
}
