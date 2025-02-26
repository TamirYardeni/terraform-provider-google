---
# ----------------------------------------------------------------------------
#
#     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
#
# ----------------------------------------------------------------------------
#
#     This file is automatically generated by Magic Modules and manual
#     changes will be clobbered when the file is regenerated.
#
#     Please read more about how to change this file in
#     .github/CONTRIBUTING.md.
#
# ----------------------------------------------------------------------------
subcategory: "Cloud Key Management Service"
page_title: "Google: google_kms_crypto_key_version"
description: |-
  A `CryptoKeyVersion` represents an individual cryptographic key, and the associated key material.
---

# google\_kms\_crypto\_key\_version

A `CryptoKeyVersion` represents an individual cryptographic key, and the associated key material.


Destroying a cryptoKeyVersion will not delete the resource from the project.


To get more information about CryptoKeyVersion, see:

* [API documentation](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions)
* How-to Guides
    * [Creating a key Version](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys.cryptoKeyVersions/create)

## Example Usage - Kms Crypto Key Version Basic


```hcl
resource "google_kms_key_ring" "keyring" {
  name     = "keyring-example"
  location = "global"
}

resource "google_kms_crypto_key" "cryptokey" {
  name            = "crypto-key-example"
  key_ring        = google_kms_key_ring.keyring.id
  rotation_period = "100000s"
}

resource "google_kms_crypto_key_version" "example-key" {
  crypto_key = google_kms_crypto_key.cryptokey.id
}
```

## Argument Reference

The following arguments are supported:


* `crypto_key` -
  (Required)
  The name of the cryptoKey associated with the CryptoKeyVersions.
  Format: `'projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}'`


- - -


* `state` -
  (Optional)
  The current state of the CryptoKeyVersion.
  Possible values are `PENDING_GENERATION`, `ENABLED`, `DISABLED`, `DESTROYED`, `DESTROY_SCHEDULED`, `PENDING_IMPORT`, and `IMPORT_FAILED`.


## Attributes Reference

In addition to the arguments listed above, the following computed attributes are exported:

* `id` - an identifier for the resource with format `{{name}}`

* `name` -
  The resource name for this CryptoKeyVersion.

* `protection_level` -
  The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.

* `generate_time` -
  The time this CryptoKeyVersion key material was generated

* `algorithm` -
  The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.

* `attestation` -
  Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.
  Only provided for key versions with protectionLevel HSM.
  Structure is [documented below](#nested_attestation).


<a name="nested_attestation"></a>The `attestation` block contains:

* `format` -
  The format of the attestation data.

* `content` -
  The attestation data provided by the HSM when the key operation was performed.

* `cert_chains` -
  (Optional)
  The certificate chains needed to validate the attestation
  Structure is [documented below](#nested_cert_chains).

* `external_protection_level_options` -
  (Optional)
  ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level and EXTERNAL_VPC protection levels.
  Structure is [documented below](#nested_external_protection_level_options).


<a name="nested_cert_chains"></a>The `cert_chains` block supports:

* `cavium_certs` -
  (Optional)
  Cavium certificate chain corresponding to the attestation.

* `google_card_certs` -
  (Optional)
  Google card certificate chain corresponding to the attestation.

* `google_partition_certs` -
  (Optional)
  Google partition certificate chain corresponding to the attestation.

<a name="nested_external_protection_level_options"></a>The `external_protection_level_options` block supports:

* `external_key_uri` -
  (Optional)
  The URI for an external resource that this CryptoKeyVersion represents.

* `ekm_connection_key_path` -
  (Optional)
  The path to the external key material on the EKM when using EkmConnection e.g., "v0/my/key". Set this field instead of externalKeyUri when using an EkmConnection.

## Timeouts

This resource provides the following
[Timeouts](/docs/configuration/resources.html#timeouts) configuration options:

- `create` - Default is 20 minutes.
- `update` - Default is 20 minutes.
- `delete` - Default is 20 minutes.

## Import


CryptoKeyVersion can be imported using any of these accepted formats:

```
$ terraform import google_kms_crypto_key_version.default {{name}}
```
