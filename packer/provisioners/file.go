package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func FileResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"generated": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}
