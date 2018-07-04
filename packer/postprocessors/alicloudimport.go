package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func AlicloudImportResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"only": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "For specifying a builder input",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"except": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "For excluding a builder input",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"keep_input_artifact": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"pipeline": &schema.Schema{
				Optional:    true,
				MaxItems:    1,
				Type:        schema.TypeList,
				Description: "Create a sequence definition",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set": &schema.Schema{
							Required:    true,
							Type:        schema.TypeInt,
							Description: "The set(group) that is under",
						},
						"order": &schema.Schema{
							Required:    true,
							Type:        schema.TypeInt,
							Description: "The order to run in",
						},
					},
				},
			},

			"packer_build_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_builder_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_debug": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"packer_force": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"packer_on_error": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_user_variables": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},

			"access_key": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"secret_key": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"region": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"skip_region_validation": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"security_token": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_version": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_description": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_share_account": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"image_unshare_account": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"image_copy_regions": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"image_copy_names": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"image_force_delete": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"image_force_delete_snapshots": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"image_force_delete_instances": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"associate_public_ip_address": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"zone_id": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"io_optimized": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"instance_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"description": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"source_image": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"force_stop_instance": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"security_group_id": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"security_group_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"user_data": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"user_data_file": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"vpc_id": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"vpc_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"vpc_cidr_block": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"vswitch_id": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"instance_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"internet_charge_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"internet_max_bandwidth_out": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"temporary_key_pair_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"communicator": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_host": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_port": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"ssh_username": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_private_key_file": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_pty": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_agent_auth": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_disable_agent_forwarding": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_handshake_attempts": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"ssh_bastion_host": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_bastion_port": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"ssh_bastion_agent_auth": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_bastion_username": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_bastion_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_bastion_private_key_file": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_file_transfer_method": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_proxy_host": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_proxy_port": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"ssh_proxy_username": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_proxy_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_keep_alive_interval": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_read_write_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_username": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_password": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_host": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_port": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"winrm_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"winrm_use_ssl": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"winrm_insecure": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"winrm_use_ntlm": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"ssh_keypair_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"ssh_private_ip": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"oss_bucket_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"oss_key_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"skip_clean": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"tags": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},

			"image_os_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_platform": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_architecture": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_system_size": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"format": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}
