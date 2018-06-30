
package builders

import "github.com/hashicorp/terraform/helper/schema"

func AmazonEBSResource() *schema.Resource {
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

"ami_block_device_mappings": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"launch_block_device_mappings": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"associate_public_ip_address": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"availability_zone": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"disable_stop_instance": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ebs_optimized": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"enable_t2_unlimited": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"iam_instance_profile": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"shutdown_behavior": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"instance_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"run_tags": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"security_group_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"security_group_ids": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

		},

"source_ami": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"most_recent": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"spot_price": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"spot_price_auto_product": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"subnet_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"temporary_key_pair_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"temporary_security_group_source_cidr": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"user_data": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"user_data_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vpc_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"windows_password_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"communicator": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_port": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_private_key_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_pty": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_agent_auth": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_disable_agent_forwarding": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_handshake_attempts": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_bastion_host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_bastion_port": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_bastion_agent_auth": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_bastion_username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_bastion_password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_bastion_private_key_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_file_transfer_method": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_proxy_host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_proxy_port": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_proxy_username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_proxy_password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_keep_alive_interval": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"ssh_read_write_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"winrm_username": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"winrm_password": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"winrm_host": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"winrm_port": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"winrm_timeout": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"winrm_use_ssl": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"winrm_insecure": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"winrm_use_ntlm": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"ssh_keypair_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"ssh_interface": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"run_volume_tags": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

	},
	}
}

