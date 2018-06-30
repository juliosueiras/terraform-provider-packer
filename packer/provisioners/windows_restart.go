package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func WindowsRestartResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"restart_command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restart_check_command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restart_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
