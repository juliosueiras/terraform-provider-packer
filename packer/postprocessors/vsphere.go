
package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VSphereResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
"packer_build_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"packer_builder_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"packer_debug": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"packer_force": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"packer_on_error": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"packer_user_variables": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"cluster": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"datacenter": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"datastore": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"disk_mode": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"insecure": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"options": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"overwrite": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"resource_pool": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vm_folder": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vm_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vm_network": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

