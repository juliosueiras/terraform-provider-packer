
package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func AmazonImportResource() *schema.Resource {
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

"access_key": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"custom_endpoint_ec2": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"mfa_code": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"profile": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"region": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"secret_key": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"skip_region_validation": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"skip_metadata_api_check": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"token": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"s3_bucket_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"s3_key_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"skip_clean": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"tags": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"ami_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ami_description": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ami_users": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"ami_groups": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"license_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"role_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

	},
	}
}

