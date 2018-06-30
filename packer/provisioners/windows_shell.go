package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func WindowsShellResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"inline": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scripts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"binary": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"environment_vars": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"execute_command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_retry_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
