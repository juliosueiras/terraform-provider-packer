package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func WindowsRestartResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order for this provisioner to run in",
			},
			"restart_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to execute to initiate the restart. By default this is shutdown /r /f /t 0 /c \"packer restart\".",
			},

			"restart_check_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "A command to execute to check if the restart succeeded. This will be done in a loop.",
			},

			"restart_timeout": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The timeout to wait for the restart. By default this is 5 minutes. Example value: 5m. If you are installing updates or have a lot of startup services, you will probably need to increase this duration.",
			},
		},
	}
}
