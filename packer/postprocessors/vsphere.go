package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VSphereResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The cluster to upload the VM to.",
			},

			"datacenter": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the datacenter within vSphere to add the VM to.",
			},

			"datastore": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"disk_mode": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Target disk format. See ovftool manual for available options. By default, \"thick\" will be used.",
			},

			"host": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The vSphere host that will be contacted to perform the VM upload.",
			},

			"insecure": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether or not the connection to vSphere can be done over an insecure connection. By default this is false.",
			},

			"options": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Custom options to add in ovftool. See ovftool --help to list all the options",
			},

			"overwrite": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If it's true force the system to overwrite the existing files instead create new ones. Default is false",
			},

			"password": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "Password to use to authenticate to the vSphere endpoint.",
			},

			"resource_pool": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The resource pool to upload the VM to.",
			},

			"username": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The username to use to authenticate to the vSphere endpoint.",
			},

			"vm_folder": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The folder within the datastore to store the VM.",
			},

			"vm_name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name of the VM once it is uploaded.",
			},

			"vm_network": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the VM network this VM will be added to.",
			},
		},
	}
}
