package google

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGoogleBigqueryTable() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGoogleBigqueryTableRead,

		Schema: map[string]*schema.Schema{
			"project": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dataset_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"table_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"query": {
				Type:     schema.TypeString,
				Required: false,
			},
			"schema": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"natypeme": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			/*			"data": {
							Type:     schema.TypeList,
							Computed: true,
						},
						"table_info": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_info": {
							Type:     schema.TypeString,
							Computed: true,
						},*/
		},
	}
}

func dataSourceGoogleBigqueryTableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	var diags diag.Diagnostics

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return diag.FromErr(err)
	}

	project := d.Get("project").(string)
	dataset := d.Get("dataset_name").(string)
	table := d.Get("table_name").(string)

	url := fmt.Sprintf("https://bigquery.googleapis.com/bigquery/v2/projects/%s/datasets/%s/tables/%s", project, dataset, table)
	res, err := sendRequest(config, "GET", "", url, userAgent, nil)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error retrieving table data: %s", err))
	}

	var resSchema map[string]interface{}
	resSchema = res["schema"].(map[string]interface{})["fields"].(map[string]interface{})
	if err := d.Set("schema", resSchema); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
