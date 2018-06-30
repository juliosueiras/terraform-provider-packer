package builders

import "github.com/hashicorp/terraform/helper/schema"

// VMwareISOBuilderResource ...
func VMwareISOBuilderResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"winrm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     winrmCommunicatorResource(),
			},
			"ssh": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     sshCommunicatorResource(),
			},
			"iso": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem:     isoBuilderResource(),
			},
			"communicator": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"shutdown_command": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"boot_command": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"boot_wait": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cdrom_adapter_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"disable_vnc": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"disk_additional_size": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"disk_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"disk_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"floppy_dirs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"remote": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     remoteVMwareISOBuilderResource(),
			},
			"floppy_files": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fusion_app_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"guest_os_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"headless": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"http_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_port_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"http_port_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"iso_target_extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iso_target_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"iso_urls": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_adapter_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"output_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parallel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"shutdown_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_compaction": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"skip_export": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"keep_registered": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ovftool_options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sound": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tools_upload_flavor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tools_upload_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"usb": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vmdk_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vmx_data": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"vmx_data_post": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"vmx_remove_ethernet_interfaces": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vmx_template_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnc_bind_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vnc_disable_password": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vnc_port_min": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vnc_port_max": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func remoteVMwareISOBuilderResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"remote_cache_datastore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_cache_directory": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_datastore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"remote_private_key_file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
