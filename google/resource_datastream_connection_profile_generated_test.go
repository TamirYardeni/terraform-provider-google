// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDatastreamConnectionProfile_datastreamConnectionProfileBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatastreamConnectionProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastreamConnectionProfile_datastreamConnectionProfileBasicExample(context),
			},
			{
				ResourceName:            "google_datastream_connection_profile.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location"},
			},
		},
	})
}

func testAccDatastreamConnectionProfile_datastreamConnectionProfileBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "tf-test-my-profile%{random_suffix}"

	gcs_profile {
		bucket    = "my-bucket"
		root_path = "/path"
	}
}
`, context)
}

func TestAccDatastreamConnectionProfile_datastreamConnectionProfileBigqueryPrivateConnectionExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatastreamConnectionProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastreamConnectionProfile_datastreamConnectionProfileBigqueryPrivateConnectionExample(context),
			},
			{
				ResourceName:            "google_datastream_connection_profile.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location"},
			},
		},
	})
}

func testAccDatastreamConnectionProfile_datastreamConnectionProfileBigqueryPrivateConnectionExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_datastream_private_connection" "private_connection" {
	display_name          = "Connection profile"
	location              = "us-central1"
	private_connection_id = "tf-test-my-connection%{random_suffix}"

	labels = {
		key = "value"
	}

	vpc_peering_config {
		vpc = google_compute_network.default.id
		subnet = "10.0.0.0/29"
	}
}

resource "google_compute_network" "default" {
	name = "tf-test-my-network%{random_suffix}"
}

resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "tf-test-my-profile%{random_suffix}"

	bigquery_profile {}

	private_connectivity {
		private_connection = google_datastream_private_connection.private_connection.id
	}
}
`, context)
}

func TestAccDatastreamConnectionProfile_datastreamConnectionProfileFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDatastreamConnectionProfileDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDatastreamConnectionProfile_datastreamConnectionProfileFullExample(context),
			},
			{
				ResourceName:            "google_datastream_connection_profile.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"connection_profile_id", "location", "forward_ssh_connectivity.0.password"},
			},
		},
	})
}

func testAccDatastreamConnectionProfile_datastreamConnectionProfileFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_datastream_connection_profile" "default" {
	display_name          = "Connection profile"
	location              = "us-central1"
	connection_profile_id = "tf-test-my-profile%{random_suffix}"

	gcs_profile {
		bucket    = "my-bucket"
		root_path = "/path"
	}

	forward_ssh_connectivity {
		hostname = "google.com"
		username = "my-user"
		port     = 8022
		password = "swordfish"
	}
}
`, context)
}

func testAccCheckDatastreamConnectionProfileDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_datastream_connection_profile" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DatastreamBasePath}}projects/{{project}}/locations/{{location}}/connectionProfiles/{{connection_profile_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("DatastreamConnectionProfile still exists at %s", url)
			}
		}

		return nil
	}
}
