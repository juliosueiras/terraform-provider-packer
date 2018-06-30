
package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func ShellLocalResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type: schema.TypeInt,
				Required: true,
			},
	},
	}
}

