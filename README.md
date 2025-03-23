# TFLint Okta Template
[![Build Status](https://github.com/vazuev/tflint-ruleset-okta/workflows/build/badge.svg?branch=main)](https://github.com/vazuev/tflint-ruleset-okta/actions)

This is a repository for building an Okta ruleset.

## Requirements

- TFLint v0.30+
- Go v1.24

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "okta" {
  enabled = true
  version = "0.1.02"
  source  = "github.com/vazuev/tflint-ruleset-okta"
  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----

  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
  yUl3m1viNNT6muQXVx655+y7EQX2NOBpcjI=
  =vEzD
  -----END PGP PUBLIC KEY BLOCK-----
  KEY
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|okta_policy_name_rule.go|Rule for checking length of okta_auth_server_policy_rule attribute name|ERROR|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```
