package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func VSphereCloneResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},

			"shutdown_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to use to gracefully shut down the machine once all the provisioning is done. By default this is an empty string, which tells Packer to just forcefully shut down the machine.",
			},

			"shutdown_timeout": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The amount of time to wait after executing the shutdown_command for the virtual machine to actually shut down. If it doesn't shut down in this time, it is an error. By default, the timeout is 5m or five minutes.",
			},

			"communicator": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"ssh": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     communicators.SSHCommunicatorResource(),
			},
			"winrm": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     communicators.WinRMCommunicatorResource(),
			},

			"disk_size": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "The size of the hard disk for the VM in megabytes. The builder uses expandable, not fixed-size virtual hard disks, so the actual file representing the disk will not use the full size unless it is full. By default this is set to 40000 (about 40 GB).",
			},

			"vm_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is the name of the VMX file for the new virtual machine, without the file extension. By default this is packer-BUILDNAME, where \"BUILDNAME\" is the name of the build.",
			},

			"notes": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"folder": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"host": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"create_snapshot": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"convert_to_template": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"template": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"linked_clone": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"cluster": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"resource_pool": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"datastore": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"cpus": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},
			"cpu_limit": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},
			"cpu_reservation": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},
			"cpu_hot_plug": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"ram": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},
			"ram_reservation": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},
			"ram_reserve_all": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"ram_hot_plug": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"nestedhv": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"configuration_parameters": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},
			"boot_order": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"vcenter_server": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "vCenter server hostname.",
			},
			"username": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "vSphere username",
			},
			"password": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "vSphere password",
			},
			"insecure_connection": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Do not validate vCenter server's TLS certificate. Defaults to false.",
			},
			"datacenter": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "VMware datacenter name. Required if there is more than one datacenter in vCenter.",
			},
		},
	}
}
