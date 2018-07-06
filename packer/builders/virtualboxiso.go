package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func VirtualboxISOResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},
			"iso": &schema.Schema{
				Required: true,
				Type:     schema.TypeList,
				Elem:     isoBuilderResource(),
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
			"http_directory": &schema.Schema{
				Optional:    true,
				Description: "Path to a directory to serve using an HTTP server. The files in this directory will be available over HTTP that will be requestable from the virtual machine. This is useful for hosting kickstart files and so on. By default this is an empty string, which means no HTTP server will be started. The address and port of the HTTP server will be available as variables in boot_command.",
				Type:        schema.TypeString,
			},

			"http_port_min": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "These are the minimum and maximum port to use for the HTTP server started to serve the http_directory. Because Packer often runs in parallel, Packer will choose a randomly available port in this range to run the HTTP server. If you want to force the HTTP server to be on one port, make this minimum and maximum port the same. By default the values are 8000 and 9000, respectively.",
			},

			"http_port_max": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "These are the minimum and maximum port to use for the HTTP server started to serve the http_directory. Because Packer often runs in parallel, Packer will choose a randomly available port in this range to run the HTTP server. If you want to force the HTTP server to be on one port, make this minimum and maximum port the same. By default the values are 8000 and 9000, respectively.",
			},

			"floppy_files": &schema.Schema{
				Optional:    true,
				Description: "A list of files to place onto a floppy disk that is attached when the VM is booted. This is most useful for unattended Windows installs, which look for an Autounattend.xml file on removable media. By default, no floppy will be attached. All files listed in this setting get placed into the root directory of the floppy and the floppy is attached as the first floppy device. Currently, no support exists for creating sub-directories on the floppy. Wildcard characters (*, ?, and []) are allowed. Directory names are also allowed, which will add all the files found in the directory to the floppy.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"floppy_dirs": &schema.Schema{
				Optional:    true,
				Description: "A list of directories to place onto the floppy disk recursively. This is similar to the floppy_files option except that the directory structure is preserved. This is useful for when your floppy disk includes drivers or if you just want to organize it's contents as a hierarchy. Wildcard characters (*, ?, and []) are allowed.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"boot_wait": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The time to wait after booting the initial virtual machine before typing the boot_command. The value of this should be a duration. Examples are 5s and 1m30s which will cause Packer to wait five seconds and one minute 30 seconds, respectively. If this isn't specified, the default is 10s or 10 seconds.",
			},

			"boot_command": &schema.Schema{
				Optional:    true,
				Description: "This is an array of commands to type when the virtual machine is first booted. The goal of these commands should be to type just enough to initialize the operating system installer. Special keys can be typed as well, and are covered in the section below on the boot command. If this is not specified, it is assumed the installer will start itself.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"format": &schema.Schema{
				Optional:    true,
				Description: "Either ovf or ova, this specifies the output format of the exported virtual machine. This defaults to ovf.",
				Type:        schema.TypeString,
			},

			"export_opts": &schema.Schema{
				Optional:    true,
				Description: "Additional options to pass to the VBoxManage export. This can be useful for passing product information to include in the resulting appliance file.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"output_directory": &schema.Schema{
				Optional:    true,
				Description: `This is the path to the directory where the resulting virtual machine will be created. This may be relative or absolute. If relative, the path is relative to the working directory when packer is executed. This directory must not exist or be empty prior to running the builder. By default this is output-BUILDNAME where "BUILDNAME" is the name of the build.`,
				Type:        schema.TypeString,
			},

			"headless": &schema.Schema{
				Optional:    true,
				Description: "Packer defaults to building VirtualBox virtual machines by launching a GUI that shows the console of the machine being built. When this value is set to true, the machine will start without a console.",
				Type:        schema.TypeBool,
			},

			"vrdp_bind_address": &schema.Schema{
				Optional:    true,
				Description: "The IP address that should be binded to for VRDP. By default packer will use 127.0.0.1 for this. If you wish to bind to all interfaces use 0.0.0.0.",
				Type:        schema.TypeString,
			},

			"vrdp_port_min": &schema.Schema{
				Optional:    true,
				Description: "The minimum and maximum port to use for VRDP access to the virtual machine. Packer uses a randomly chosen port in this range that appears available. By default this is 5900 to 6000. The minimum and maximum ports are inclusive.",
				Type:        schema.TypeInt,
			},

			"vrdp_port_max": &schema.Schema{
				Optional:    true,
				Description: "The minimum and maximum port to use for VRDP access to the virtual machine. Packer uses a randomly chosen port in this range that appears available. By default this is 5900 to 6000. The minimum and maximum ports are inclusive.",
				Type:        schema.TypeInt,
			},

			"shutdown_command": &schema.Schema{
				Optional:    true,
				Description: "The command to use to gracefully shut down the machine once all the provisioning is done. By default this is an empty string, which tells Packer to just forcefully shut down the machine unless a shutdown command takes place inside script so this may safely be omitted. If one or more scripts require a reboot it is suggested to leave this blank since reboots may fail and specify the final shutdown command in your last script.",
				Type:        schema.TypeString,
			},

			"shutdown_timeout": &schema.Schema{
				Optional:    true,
				Description: "The amount of time to wait after executing the shutdown_command for the virtual machine to actually shut down. If it doesn't shut down in this time, it is an error. By default, the timeout is 5m or five minutes.",
				Type:        schema.TypeString,
			},

			"post_shutdown_delay": &schema.Schema{
				Optional:    true,
				Description: "The amount of time to wait after shutting down the virtual machine. If you get the error Error removing floppy controller, you might need to set this to 5m or so. By default, the delay is 0s or disabled.",
				Type:        schema.TypeString,
			},

			"communicator": &schema.Schema{
				Optional:    true,
				Description: "",
				Type:        schema.TypeString,
			},

			"vboxmanage": &schema.Schema{
				Optional:    true,
				Description: "Custom VBoxManage commands to execute in order to further customize the virtual machine being created. The value of this is an array of commands to execute. The commands are executed in the order defined in the template. For each command, the command is defined itself as an array of strings, where each string represents a single argument on the command-line to VBoxManage (but excluding VBoxManage itself). Each arg is treated as a configuration template, where the Name variable is replaced with the VM name.",
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"values": &schema.Schema{
							Required:    true,
							Type:        schema.TypeList,
							Description: "Values for vboxmanage",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			"vboxmanage_post": &schema.Schema{
				Optional:    true,
				Description: "Identical to vboxmanage, except that it is run after the virtual machine is shutdown, and before the virtual machine is exported.",
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"values": &schema.Schema{
							Required:    true,
							Type:        schema.TypeList,
							Description: "Values for vboxmanage",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			"virtualbox_version_file": &schema.Schema{
				Optional:    true,
				Description: "The path within the virtual machine to upload a file that contains the VirtualBox version that was used to create the machine. This information can be useful for provisioning. By default this is .vbox_version, which will generally be upload it into the home directory. Set to an empty string to skip uploading this file, which can be useful when using the none communicator.",
				Type:        schema.TypeString,
			},

			"disk_size": &schema.Schema{
				Optional:    true,
				Description: "The size, in megabytes, of the hard disk to create for the VM. By default, this is 40000 (about 40 GB).",
				Type:        schema.TypeInt,
			},

			"guest_additions_mode": &schema.Schema{
				Optional:    true,
				Description: "The method by which guest additions are made available to the guest for installation. Valid options are upload, attach, or disable. If the mode is attach the guest additions ISO will be attached as a CD device to the virtual machine. If the mode is upload the guest additions ISO will be uploaded to the path specified by guest_additions_path. The default value is upload. If disable is used, guest additions won't be downloaded, either.",
				Type:        schema.TypeString,
			},

			"guest_additions_path": &schema.Schema{
				Optional:    true,
				Description: "The path on the guest virtual machine where the VirtualBox guest additions ISO will be uploaded. By default this is VBoxGuestAdditions.iso which should upload into the login directory of the user. This is a configuration template where the Version variable is replaced with the VirtualBox version.",
				Type:        schema.TypeString,
			},

			"guest_additions_sha256": &schema.Schema{
				Optional:    true,
				Description: "The SHA256 checksum of the guest additions ISO that will be uploaded to the guest VM. By default the checksums will be downloaded from the VirtualBox website, so this only needs to be set if you want to be explicit about the checksum.",
				Type:        schema.TypeString,
			},

			"guest_additions_url": &schema.Schema{
				Optional:    true,
				Description: "The URL to the guest additions ISO to upload. This can also be a file URL if the ISO is at a local path. By default, the VirtualBox builder will attempt to find the guest additions ISO on the local file system. If it is not available locally, the builder will download the proper guest additions ISO from the internet.",
				Type:        schema.TypeString,
			},

			"guest_os_type": &schema.Schema{
				Optional:    true,
				Description: "The guest OS type being installed. By default this is other, but you can get dramatic performance improvements by setting this to the proper value. To view all available values for this run VBoxManage list ostypes. Setting the correct value hints to VirtualBox how to optimize the virtual hardware to work best with that operating system.",
				Type:        schema.TypeString,
			},

			"hard_drive_discard": &schema.Schema{
				Optional:    true,
				Description: "When this value is set to true, a VDI image will be shrunk in response to the trim command from the guest OS. The size of the cleared area must be at least 1MB. Also set hard_drive_nonrotational to true to enable TRIM support.",
				Type:        schema.TypeBool,
			},

			"hard_drive_interface": &schema.Schema{
				Optional:    true,
				Description: "The type of controller that the primary hard drive is attached to, defaults to ide. When set to sata, the drive is attached to an AHCI SATA controller. When set to scsi, the drive is attached to an LsiLogic SCSI controller.",
				Type:        schema.TypeString,
			},

			"sata_port_count": &schema.Schema{
				Optional:    true,
				Description: "The number of ports available on any SATA controller created, defaults to 1. VirtualBox supports up to 30 ports on a maximum of 1 SATA controller. Increasing this value can be useful if you want to attach additional drives.",
				Type:        schema.TypeInt,
			},

			"hard_drive_nonrotational": &schema.Schema{
				Optional:    true,
				Description: "Forces some guests (i.e. Windows 7+) to treat disks as SSDs and stops them from performing disk fragmentation. Also set hard_drive_discard to true to enable TRIM support.",
				Type:        schema.TypeBool,
			},

			"iso_interface": &schema.Schema{
				Optional:    true,
				Description: "The type of controller that the ISO is attached to, defaults to ide. When set to sata, the drive is attached to an AHCI SATA controller.",
				Type:        schema.TypeString,
			},

			"keep_registered": &schema.Schema{
				Optional:    true,
				Description: "Set this to true if you would like to keep the VM registered with virtualbox. Defaults to false.",
				Type:        schema.TypeBool,
			},

			"skip_export": &schema.Schema{
				Optional:    true,
				Description: "Defaults to false. When enabled, Packer will not export the VM. Useful if the build output is not the resultant image, but created inside the VM.",
				Type:        schema.TypeBool,
			},
			"ssh_host_port_min": &schema.Schema{
				Optional:    true,
				Description: "The minimum and maximum port to use for the SSH port on the host machine which is forwarded to the SSH port on the guest machine. Because Packer often runs in parallel, Packer will choose a randomly available port in this range to use as the host port. By default this is 2222 to 4444.",
				Type:        schema.TypeInt,
			},
			"ssh_host_port_max": &schema.Schema{
				Optional:    true,
				Description: "The minimum and maximum port to use for the SSH port on the host machine which is forwarded to the SSH port on the guest machine. Because Packer often runs in parallel, Packer will choose a randomly available port in this range to use as the host port. By default this is 2222 to 4444.",
				Type:        schema.TypeInt,
			},
			"ssh_host_port_max": &schema.Schema{
				Optional:    true,
				Description: "The minimum and maximum port to use for the SSH port on the host machine which is forwarded to the SSH port on the guest machine. Because Packer often runs in parallel, Packer will choose a randomly available port in this range to use as the host port. By default this is 2222 to 4444.",
				Type:        schema.TypeInt,
			},
			"ssh_skip_nat_mapping": &schema.Schema{
				Optional:    true,
				Description: "Defaults to false. When enabled, Packer does not setup forwarded port mapping for SSH requests and uses ssh_port on the host to communicate to the virtual machine.",
				Type:        schema.TypeBool,
			},

			"vm_name": &schema.Schema{
				Optional:    true,
				Description: "This is the name of the OVF file for the new virtual machine, without the file extension. By default this is packer-BUILDNAME, where \"BUILDNAME\" is the name of the build.",
				Type:        schema.TypeString,
			},
		},
	}
}
