# TFLint Ruleset for the Terraform Provider for Okta

[![Build Status](https://github.com/fionn/tflint-ruleset-okta/actions/workflows/main.yml/badge.svg)](https://github.com/fionn/tflint-ruleset-okta/actions)
[![Latest Release](https://img.shields.io/github/v/release/fionn/tflint-ruleset-okta.svg)](https://github.com/fionn/tflint-ruleset-okta/releases/latest)

## Requirements

- TFLint v0.55
- Go v1.24

## Installation

Declare a configuration in `.tflint.hcl` with

```hcl
plugin "okta" {
  enabled = true
  version = "x.y.z"  # Replace with desired version.
  source  = "github.com/fionn/tflint-ruleset-okta"
}
```
and install with `tflint --init`.

## Rules

|Name|Description|Severity|Enabled|
| --- | --- | --- | --- |
|`okta_auth_server_policy_name_length`|Check the length of `okta_auth_server_policy`'s `name` attribute|ERROR|✓|
|`okta_app_oauth_omit_secret`|Check that OAuth application secrets are omitted|WARNING|✓|
|`okta_app_oauth_plaintext_redirect_uri`|Check that remote redirect URIs are using HTTPS|WARNING|✓|
|`okta_app_implicit_authentication_policy`|Check that applications specify an authentication policy|NOTICE||
