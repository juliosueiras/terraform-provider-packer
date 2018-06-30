
package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func AnsibleLocalResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type: schema.TypeInt,
				Required: true,
			},
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

"extra_arguments": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"group_vars": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"host_vars": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"playbook_dir": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"playbook_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"playbook_files": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"playbook_paths": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"role_paths": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"staging_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"clean_staging_directory": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"inventory_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"inventory_groups": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"galaxy_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

