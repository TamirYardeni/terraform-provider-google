package google

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleBigqueryTable(t *testing.T) {
	t.Parallel()

	newVersion := map[string]func() (*schema.Provider, error){
		"mynewprovider": func() (*schema.Provider, error) { return testAccProviders["google"], nil },
	}

	project := "tf-project-" + randString(t, 10)
	dataset := "tf-dataset-" + randString(t, 10)
	table := "tf-table-" + randString(t, 10)

	vcrTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newVersion,
		CheckDestroy:      testAccStorageBucketDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceGoogleBigqueryTableConfig(project, dataset, table),
				Check: resource.ComposeTestCheckFunc(
					checkDataSourceStateMatchesResourceStateWithIgnores("data.google_bigquery_table.foo", "google_bigquery_table.bar", map[string]struct{}{"force_destroy": {}}),
				),
			},
		},
	})
}

func testAccDataSourceGoogleBigqueryTableConfig(project string, dataset string, table string) string {
	return fmt.Sprintf(`\

resource "google_bigquery_dataset" "test_dataset" {
  dataset_id = "%s"
  project    = "%s"
}

resource "google_bigquery_table" "foo" {
  dataset_id = google_bigquery_dataset.test_dataset.dataset_id
  table_id   = "%s"
  schema	 = <<EOF
				[
				  {
					"name": "testAa",
					"type": "TIMESTAMP"
				  },
				  {
					"name": "testB",
					"type": "STRING"
				  }
				]
				EOF
}

data "google_bigquery_table" "bar" {
	project = google_bigquery_dataset.test_dataset.project
	dataset_name = google_bigquery_dataset.test_dataset.dataset_id
	table_name = google_bigquery_table.foo.table_id
}
`, project, dataset, table)
}
