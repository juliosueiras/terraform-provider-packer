package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VSphereTemplateResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The vSphere host that contains the VM built by the vmware-iso.",
			},

			"insecure": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If it's true skip verification of server certificate. Default is false",
			},

			"username": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The username to use to authenticate to the vSphere endpoint.",
			},

			"password": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "Password to use to authenticate to the vSphere endpoint.",
			},

			"datacenter": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "If you have more than one, you will need to specify which one the ESXi used.",
			},

			"folder": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Target path where the template will be created.",
			},
		},
	}
}
