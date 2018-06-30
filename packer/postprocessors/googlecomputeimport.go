
package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func GoogleComputeImportResource() *schema.Resource {
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

"bucket": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"gcs_object_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_description": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_family": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_labels": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"image_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"project_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"account_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"keep_input_artifact": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

	},
	}
}

