package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func VMWareISOResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},

			"http_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to a directory to serve using an HTTP server. The files in this directory will be available over HTTP that will be requestable from the virtual machine. This is useful for hosting kickstart files and so on. By default this is an empty string, which means no HTTP server will be started. The address and port of the HTTP server will be available as variables in boot_command. This is covered in more detail below.",
			},

			"http_port_min": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "These are the minimum and maximum port to use for the HTTP server started to serve the http_directory. Because Packer often runs in parallel, Packer will choose a randomly available port in this range to run the HTTP server. If you want to force the HTTP server to be on one port, make this minimum and maximum port the same. By default the values are 8000 and 9000, respectively",
			},

			"http_port_max": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "These are the minimum and maximum port to use for the HTTP server started to serve the http_directory. Because Packer often runs in parallel, Packer will choose a randomly available port in this range to run the HTTP server. If you want to force the HTTP server to be on one port, make this minimum and maximum port the same. By default the values are 8000 and 9000, respectively",
			},

			"iso": &schema.Schema{
				Required: true,
				Type:     schema.TypeList,
				Elem:     isoBuilderResource(),
			},
			"floppy_files": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "A list of files to place onto a floppy disk that is attached when the VM is booted. This is most useful for unattended Windows installs, which look for an Autounattend.xml file on removable media. By default, no floppy will be attached. All files listed in this setting get placed into the root directory of the floppy and the floppy is attached as the first floppy device. Currently, no support exists for creating sub-directories on the floppy. Wildcard characters (*, ?, and []) are allowed. Directory names are also allowed, which will add all the files found in the directory to the floppy.",
			},

			"floppy_dirs": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "A list of directories to place onto the floppy disk recursively. This is similar to the floppy_files option except that the directory structure is preserved. This is useful for when your floppy disk includes drivers or if you just want to organize it's contents as a hierarchy. Wildcard characters (*, ?, and []) are allowed.",
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

			"disable_vnc": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether to create a VNC connection or not. A boot_command cannot be used when this is false. Defaults to false.",
			},

			"fusion_app_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to \"VMware Fusion.app\". By default this is /Applications/VMware Fusion.app but this setting allows you to customize this.",
			},

			"output_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is the path to the directory where the resulting virtual machine will be created. This may be relative or absolute. If relative, the path is relative to the working directory when packer is executed. This directory must not exist or be empty prior to running the builder. By default this is output-BUILDNAME where \"BUILDNAME\" is the name of the build.",
			},

			"headless": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Packer defaults to building VMware virtual machines by launching a GUI that shows the console of the machine being built. When this value is set to true, the machine will start without a console. For VMware machines, Packer will output VNC connection information in case you need to connect to the console to debug the build process.",
			},

			"vnc_bind_address": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The IP address that should be binded to for VNC. By default packer will use 127.0.0.1 for this. If you wish to bind to all interfaces use 0.0.0.0.",
			},

			"vnc_port_min": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "The minimum and maximum port to use for VNC access to the virtual machine. The builder uses VNC to type the initial boot_command. Because Packer generally runs in parallel, Packer uses a randomly chosen port in this range that appears available. By default this is 5900 to 6000. The minimum and maximum ports are inclusive.",
			},

			"vnc_port_max": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "The minimum and maximum port to use for VNC access to the virtual machine. The builder uses VNC to type the initial boot_command. Because Packer generally runs in parallel, Packer uses a randomly chosen port in this range that appears available. By default this is 5900 to 6000. The minimum and maximum ports are inclusive.",
			},

			"vnc_disable_password": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Don't auto-generate a VNC password that is used to secure the VNC communication with the VM.",
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
			"tools_upload_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path in the VM to upload the VMware tools. This only takes effect if tools_upload_flavor is non-empty. This is a configuration template that has a single valid variable: Flavor, which will be the value of tools_upload_flavor. By default the upload path is set to {{.Flavor}}.iso. This setting is not used when remote_type is esx5.",
			},
			"tools_upload_flavor": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"darwin", "linux", "windows"}, true),
				Description:  "The flavor of the VMware Tools ISO to upload into the VM. Valid values are darwin, linux, and windows. By default, this is empty, which means VMware tools won't be uploaded.",
			},

			"vmx_data": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "Arbitrary key/values to enter into the virtual machine VMX file. This is for advanced users who want to set properties such as memory, CPU, etc.",
			},

			"vmx_data_post": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "Identical to vmx_data, except that it is run after the virtual machine is shutdown, and before the virtual machine is exported.",
			},

			"vmx_remove_ethernet_interfaces": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Remove all ethernet interfaces from the VMX file after building. This is for advanced users who understand the ramifications, but is useful for building Vagrant boxes since Vagrant will create ethernet interfaces when provisioning a box. Defaults to false.",
			},

			"disk_additional_size": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "The size(s) of any additional hard disks for the VM in megabytes. If this is not specified then the VM will only contain a primary hard disk. The builder uses expandable, not fixed-size virtual hard disks, so the actual file representing the disk will not use the full size unless it is full.",
			},

			"disk_adapter_type": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The adapter type of the VMware virtual disk to create. This option is for advanced usage, modify only if you know what you're doing. Some of the options you can specify are ide, sata, nvme or scsi (which uses the \"lsilogic\" scsi interface by default). If you specify another option, Packer will assume that you're specifying a scsi interface of that specified type. For more information, please consult the Virtual Disk Manager User's Guide for desktop VMware clients. For ESXi, refer to the proper ESXi documentation.",
			},

			"vmdk_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The filename of the virtual disk that'll be created, without the extension. This defaults to packer",
			},

			"disk_size": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "The size of the hard disk for the VM in megabytes. The builder uses expandable, not fixed-size virtual hard disks, so the actual file representing the disk will not use the full size unless it is full. By default this is set to 40000 (about 40 GB).",
			},

			"disk_type_id": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The type of VMware virtual disk to create. This option is for advanced usage.",
			},

			"format": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"cdrom_adapter_type": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"ide", "sata", "scsi"}, true),
				Description:  "The adapter type (or bus) that will be used by the cdrom device. This is chosen by default based on the disk adapter type. VMware tends to lean towards ide for the cdrom device unless sata is chosen for the disk adapter and so Packer attempts to mirror this logic. This field can be specified as either ide, sata, or scsi.",
			},

			"guest_os_type": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The guest OS type being installed. This will be set in the VMware VMX. By default this is other. By specifying a more specific OS type, VMware may perform some optimizations or virtual hardware changes to better support the operating system running in the virtual machine.",
			},

			"version": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The vmx hardware version for the new virtual machine. Only the default value has been tested, any other value is experimental. Default value is 9.",
			},

			"vm_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is the name of the VMX file for the new virtual machine, without the file extension. By default this is packer-BUILDNAME, where \"BUILDNAME\" is the name of the build.",
			},

			"network_adapter_type": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is the ethernet adapter type the the virtual machine will be created with. By default the e1000 network adapter type will be used by Packer. For more information, please consult the Choosing a network adapter for your virtual machine for desktop VMware clients. For ESXi, refer to the proper ESXi documentation.",
			},

			"network": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is the network type that the virtual machine will be created with. This can be one of the generic values that map to a device such as hostonly, nat, or bridged. If the network is not one of these values, then it is assumed to be a VMware network device. (VMnet0..x)",
			},

			"sound": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Enable VMware's virtual soundcard device for the VM",
			},

			"usb": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: " Enable VMware's USB bus for the guest VM. To enable usage of the XHCI bus for USB 3 (5 Gbit/s), one can use the vmx_data option to enable it by specifying true for the usb_xhci.present property",
			},

			"serial": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This specifies a serial port to add to the VM. It has a format of Type:option1,option2,.... The field Type can be one of the following values: FILE, DEVICE, PIPE, AUTO, or NONE.",
			},

			"parallel": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "his specifies a parallel port to add to the VM. It has the format of Type:option1,option2,.... Type can be one of the following values: FILE, DEVICE, AUTO, or NONE.",
			},

			"keep_registered": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Set this to true if you would like to keep the VM registered with the remote ESXi server. This is convenient if you use packer to provision VMs on ESXi and don't want to use ovftool to deploy the resulting artifact (VMX or OVA or whatever you used as format). Defaults to false.",
			},

			"ovftool_options": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Extra options to pass to ovftool during export. Each item in the array is a new argument. The options --noSSLVerify, --skipManifestCheck, and --targetType are reserved, and should not be passed to this argument.",
			},

			"skip_compaction": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "VMware-created disks are defragmented and compacted at the end of the build process using vmware-vdiskmanager. In certain rare cases, this might actually end up making the resulting disks slightly larger. If you find this to be the case, you can disable compaction using this configuration value. Defaults to false.",
			},

			"skip_export": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Defaults to false. When enabled, Packer will not export the VM. Useful if the build output is not the resultant image, but created inside the VM",
			},

			"vmx_template_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to a configuration template that defines the contents of the virtual machine VMX file for VMware. This is for advanced users only as this can render the virtual machine non-functional. See below for more information. For basic VMX modifications, try vmx_data first.",
			},

			"remote": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"remote_type": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeString,
							Description: "The type of remote machine that will be used to build this VM rather than a local desktop product. The only value accepted for this currently is esx5. If this is not set, a desktop product will be used. By default, this is not set.",
						},

						"remote_datastore": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeString,
							Description: "The path to the datastore where the resulting VM will be stored when it is built on the remote machine. By default this is datastore1. This only has an effect if remote_type is enabled.",
						},

						"remote_cache_datastore": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeString,
							Description: "The path to the datastore where supporting files will be stored during the build on the remote machine. By default this is the same as the remote_datastore option. This only has an effect if remote_type is enabled.",
						},

						"remote_cache_directory": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeString,
							Description: "The path where the ISO and/or floppy files will be stored during the build on the remote machine. The path is relative to the remote_cache_datastore on the remote machine. By default this is packer_cache. This only has an effect if remote_type is enabled.",
						},

						"remote_host": &schema.Schema{
							Required:    true,
							Type:        schema.TypeString,
							Description: "The host of the remote machine used for access. This is only required if remote_type is enabled.",
						},

						"remote_port": &schema.Schema{
							Optional: true,
							Type:     schema.TypeInt,
						},

						"remote_username": &schema.Schema{
							Required:    true,
							Type:        schema.TypeString,
							Description: "The username for the SSH user that will access the remote machine. This is required if remote_type is enabled.",
						},

						"remote_password": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeString,
							Description: "The SSH password for the user used to access the remote machine. By default this is empty. This only has an effect if remote_type is enabled.",
						},

						"remote_private_key_file": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeString,
							Description: "The path to the PEM encoded private key file for the user used to access the remote machine. By default this is empty. This only has an effect if remote_type is enabled.",
						},
					},
				},
			},
		},
	}
}
