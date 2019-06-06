package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VSphereTemplateResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"only": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "For specifying a builder input",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"except": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "For excluding a builder input",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"keep_input_artifact": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"pipeline": &schema.Schema{
				Optional:    true,
				MaxItems:    1,
				Type:        schema.TypeList,
				Description: "Create a sequence definition",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set": &schema.Schema{
							Required:    true,
							Type:        schema.TypeInt,
							Description: "The set(group) that is under",
						},
						"order": &schema.Schema{
							Required:    true,
							Type:        schema.TypeInt,
							Description: "The order to run in",
						},
					},
				},
			},

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
			"snapshot_enable": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"snapshot_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"snapshot_description": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"snapshot_memory": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"snapshot_quiesce": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
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
