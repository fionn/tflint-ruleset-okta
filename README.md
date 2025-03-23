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

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----

  mQINBGco6XgBEADlN00F0WdZURlnyD5MuiqyvB2i0qGiFrLZADbRKbVujyISSCkP
  0zvD/Qbj0oUH2WYKcXBqHvMK7FEsZlguO72oPKYdp9PX5qkMAv+oF+r9RIcxnx+f
  xQxdxjIGzzJRm98Xv32noTV+YjKwvmk1W0OAk02IVfSlRwPEziVp63WH1gosXjwW
  luK5B8NCvXu6XAd+evEqtghkkd2tK1s/h3Iru0dcygibHfbxh8cDmF8TWUPV7Tjn
  e8ANHHljoESH1/7F3cpq6OamO9q4yWKh4DRzqAQ93KTgHkkOZSJsnFyRNn1PTXb6
  y7Vk+WI4LNsvo/vqqqgh4N+6Ii5EUdrbFE/6L0ya+iBjE2aPP8FhIjjBb9fVOpdY
  ogo7eG4WrWuq9w+8ep+ygmSR4v9quMAV1I2NU90Od1cnrppzP72iwgs69k9M8dPH
  jBeEO5c6lcAkaxlxSetZTu9sPV0xuFFk9DTPYz0I8CLLWo2GSkooGjjduJWejbOW
  yca6RTgQX6M8yikkqZ7cBcpTV/ps3DTdDN4VgDtmnFB6sS+8QIa6/Q5xOJZeZnJ8
  CWq3km7JPG1XmiqEF54t0pV0axvbpmsD5HTW8ZmS8chx7pK1d3yiAvAOxRRq9Jro
  h6kETmZyhM2UPr5lWVNnIGm5TOJiS/f9J6k5x0RbvaMxxu4CaAMagUHJLwARAQAB
  tCNnaXRodWIuY29tL2Zpb25uIDxmaW9ubkBnaXRodWIuY29tPokCUQQTAQgAOxYh
  BF3RuD1BjR3luO5n/MHAfDRDGEtyBQJnKOl4AhsDBQsJCAcCAiICBhUKCQgLAgQW
  AgMBAh4HAheAAAoJEMHAfDRDGEtydMIP/Al4fpvl8HYfracFb2bR82I816pYOyHI
  DMVU7hv2Gm6+Y5pF0v+GO9Xq+k3+RWqME3QSV/FD1Xcny1FrbTL7RqRbYhJkfrUW
  hiq5QSJleoj4+CqRTKlU6mA6NBmnBuwEhNMo6Aj06VKIyifcUqp9oPuJq3L3+0sD
  eZBjKK44YMk0fEhA7I4XLHX58PfXxDY/S32wAJ6m9lbcCCeL2N4tB0E4QDREu/o/
  aGF93uVG0y3W8zAfWZRBwbcEVG0iMMfcRIty2xoMNrBpwrMO9hEXqJJJNyQQvXUQ
  IrH45FFrNTvcbVksYrtcRN+2JqHS8oRFInbjDcg7bWs2g8AI24Ve9xtiwhvvkHCR
  5LOhbxM+OmuouYpHKZ660QiXpmhiw3xXUjzUSQMEcaFvQHHH62M73FflmqqPprfP
  U32zyjpsj1oZd51sM6GkByEEsRCvr7596BSntMY1Ujmnui43R8SgYTkep+HAg6HL
  E25y9+9tUUMm4SPvJUjpJGYDxJgLW7xme/c/Y8oRwBkNpIyImdyklfwsony/JLNv
  hd7/dz2Dzi89qgLz3ivDrYWcVcEAJ3fwCJ5f4xPIrVv7kAZ5h5g2SBgTvNLqYBI7
  +wopdKt1ZOQ5K17Awk+ka+1AqpOMMDKkERHdRXQPyhnw6XLDhQePXET63Ra1nVwz
  q2ZhLadqJl0JiHUEEBYKAB0WIQRshvCwZ5KaqUFUxerw5hH2wB6mAAUCZyjp0AAK
  CRDw5hH2wB6mAMRCAPoCS09+62vOhCVPbu03WwSLaBbobMSSHVjpFvyCGz+/HgD9
  HC8rgJIKcpKxKQ15DlsGjTVDmoMVLg6/mUeUzB2OUAI=
  =atYQ
  -----END PGP PUBLIC KEY BLOCK-----
  KEY
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
