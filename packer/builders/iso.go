package builders

import "github.com/hashicorp/terraform/helper/schema"

func isoBuilderResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"iso_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"iso_checksum": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"iso_checksum_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iso_checksum_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
