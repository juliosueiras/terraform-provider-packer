
package builders

import "github.com/hashicorp/terraform/helper/schema"

func AmazonChrootResource() *schema.Resource {
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

"ami_block_device_mappings": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"ami_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ami_description": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ami_virtualization_type": &schema.Schema{
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

"ami_product_codes": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"ami_regions": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"skip_region_validation": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"tags": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"ena_support": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"sriov_support": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"force_deregister": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"force_delete_snapshot": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"encrypt_boot": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"kms_key_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"region_kms_key_ids": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"snapshot_tags": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"snapshot_users": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"snapshot_groups": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

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

"skip_metadata_api_check": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"token": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"chroot_mounts": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"command_wrapper": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"copy_files": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"device_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"nvme_device_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"from_scratch": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"mount_options": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"mount_partition": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"mount_path": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"post_mount_commands": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"pre_mount_commands": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"root_device_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"root_volume_size": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"source_ami": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"most_recent": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

	},
	}
}

