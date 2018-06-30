
package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VagrantCloudResource() *schema.Resource {
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

"box_tag": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"version": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"version_description": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"no_release": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"access_token": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vagrant_cloud_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"box_download_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

