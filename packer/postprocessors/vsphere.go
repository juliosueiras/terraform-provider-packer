package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VSphereResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"datastore": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vm_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disk_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"insecure": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"resource_pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_folder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"overwrite": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}
