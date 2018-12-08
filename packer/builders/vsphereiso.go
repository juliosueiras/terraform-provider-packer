package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func VSphereISOResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},

			"iso_paths": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of data store paths to ISO files that will be mounted to the VM. Example \"[datastore1] ISO/ubuntu-16.04.3-server-amd64.iso\"",
			},
			"floppy_files": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "A list of files to place onto a floppy disk that is attached when the VM is booted. This is most useful for unattended Windows installs, which look for an Autounattend.xml file on removable media. By default, no floppy will be attached. All files listed in this setting get placed into the root directory of the floppy and the floppy is attached as the first floppy device. Currently, no support exists for creating sub-directories on the floppy. Wildcard characters (*, ?, and []) are allowed. Directory names are also allowed, which will add all the files found in the directory to the floppy.",
			},
			"floppy_img_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Data store path to a floppy image that will be mounted to the VM. Cannot be used with floppy_files or floppy_dir options. Example [datastore1] ISO/VMware Tools/10.2.0/pvscsi-Windows8.flp.",
			},

			"boot_wait": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The time to wait after booting the initial virtual machine before typing the boot_command. The value of this should be a duration. Examples are 5s and 1m30s which will cause Packer to wait five seconds and one minute 30 seconds, respectively. If this isn't specified, the default is 10s or 10 seconds.",
			},

			"boot_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "This is an array of commands to type when the virtual machine is first booted. The goal of these commands should be to type just enough to initialize the operating system installer. Special keys can be typed as well, and are covered in the section below on the boot command. If this is not specified, it is assumed the installer will start itself.",
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

			"guest_os_type": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The guest OS type being installed. This will be set in the VMware VMX. By default this is other. By specifying a more specific OS type, VMware may perform some optimizations or virtual hardware changes to better support the operating system running in the virtual machine.",
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
			"vm_version": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},
			"disk_controller_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"disk_thin_provisioned": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"network_card": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"usb_controller": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"cdrom_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"firmware": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"network": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Set network VM will be connected to.",
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
