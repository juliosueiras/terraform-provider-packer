
package builders

import "github.com/hashicorp/terraform/helper/schema"

func AzureARMResource() *schema.Resource {
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

"client_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"client_secret": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"object_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"tenant_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"subscription_id": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"capture_name_prefix": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"capture_container_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_publisher": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_offer": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_sku": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_version": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"image_url": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"custom_managed_image_resource_group_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"custom_managed_image_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"location": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"vm_size": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"managed_image_resource_group_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"managed_image_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"managed_image_storage_account_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"azure_tags": &schema.Schema{
Optional: true,
			Type: schema.TypeMap,

		},

"resource_group_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"storage_account": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"temp_compute_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"temp_resource_group_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"build_resource_group_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"cloud_environment_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"private_virtual_network_with_public_ip": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

"virtual_network_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"virtual_network_subnet_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"virtual_network_resource_group_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"custom_data_file": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"plan_name": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"plan_product": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"plan_publisher": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"plan_promotion_code": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"os_type": &schema.Schema{
Optional: true,
			Type: schema.TypeString,

		},

"os_disk_size_gb": &schema.Schema{
Optional: true,
			Type: schema.TypeInt,

		},

"disk_additional_size": &schema.Schema{
Optional: true,
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},

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

"async_resourcegroup_delete": &schema.Schema{
Optional: true,
			Type: schema.TypeBool,

		},

	},
	}
}

