# TFLint Ruleset for the Terraform Provider for Okta

[![Build Status](https://github.com/fionn/tflint-ruleset-okta/actions/workflows/main.yml/badge.svg)](https://github.com/fionn/tflint-ruleset-okta/actions)
[![Latest Release](https://img.shields.io/github/v/release/fionn/tflint-ruleset-okta.svg)](https://github.com/fionn/tflint-ruleset-okta/releases/tag/v0.1.5)

## Requirements

- TFLint v0.55
- Go v1.24

## Installation

Declare a configuration in `.tflint.hcl` with

```hcl
plugin "okta" {
  enabled = true
  version = "0.1.5"
  source  = "github.com/fionn/tflint-ruleset-okta"
}
```
and install with `tflint --init`.

## Rules

|Name|Description|Severity|Enabled|
| --- | --- | --- | --- |
|`okta_policy_name`|Check the length of `okta_auth_server_policy`'s `name` attribute|ERROR|✔|
|`okta_app_oauth_omit_secret`|Check that OAuth application secrets are omitted|WARNING|✔|

## Build

Build with `make build`.

Install with `make install`.
