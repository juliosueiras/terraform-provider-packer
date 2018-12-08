package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func VMWareVMXResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
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
			"keep_registered": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Set this to true if you would like to keep the VM registered with the remote ESXi server. This is convenient if you use packer to provision VMs on ESXi and don't want to use ovftool to deploy the resulting artifact (VMX or OVA or whatever you used as format). Defaults to false.",
			},
			"skip_export": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Defaults to false. When enabled, Packer will not export the VM. Useful if the build output is not the resultant image, but created inside the VM",
			},

			"winrm": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     communicators.WinRMCommunicatorResource(),
			},
			"http_directory": &schema.Schema{
				Optional:    true,
				Description: "Path to a directory to serve using an HTTP server. The files in this directory will be available over HTTP that will be requestable from the virtual machine. This is useful for hosting kickstart files and so on. By default this is an empty string, which means no HTTP server will be started. The address and port of the HTTP server will be available as variables in boot_command. This is covered in more detail below.",
				Type:        schema.TypeString,
			},

			"http_port_min": &schema.Schema{
				Optional:    true,
				Description: "These are the minimum and maximum port to use for the HTTP server started to serve the http_directory. Because Packer often runs in parallel, Packer will choose a randomly available port in this range to run the HTTP server. If you want to force the HTTP server to be on one port, make this minimum and maximum port the same. By default the values are 8000 and 9000, respectively.",
				Type:        schema.TypeInt,
			},

			"http_port_max": &schema.Schema{
				Optional:    true,
				Description: "These are the minimum and maximum port to use for the HTTP server started to serve the http_directory. Because Packer often runs in parallel, Packer will choose a randomly available port in this range to run the HTTP server. If you want to force the HTTP server to be on one port, make this minimum and maximum port the same. By default the values are 8000 and 9000, respectively.",
				Type:        schema.TypeInt,
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
				Description: "The time to wait after booting the initial virtual machine before typing the boot_command. The value of this should be a duration. Examples are 5s and 1m30s which will cause Packer to wait five seconds and one minute 30 seconds, respectively. If this isn't specified, the default is 10s or 10 seconds.",
				Type:        schema.TypeString,
			},

			"boot_command": &schema.Schema{
				Optional:    true,
				Description: "This is an array of commands to type when the virtual machine is first booted. The goal of these commands should be to type just enough to initialize the operating system installer. Special keys can be typed as well, and are covered in the section below on the boot command. If this is not specified, it is assumed the installer will start itself.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"disable_vnc": &schema.Schema{
				Optional:    true,
				Description: "Whether to create a VNC connection or not. A boot_command cannot be used when this is false. Defaults to false.",
				Type:        schema.TypeBool,
			},

			"fusion_app_path": &schema.Schema{
				Optional:    true,
				Description: "Path to \"VMware Fusion.app\". By default this is /Applications/VMware Fusion.app but this setting allows you to customize this.",
				Type:        schema.TypeString,
			},

			"output_directory": &schema.Schema{
				Optional:    true,
				Description: "This is the path to the directory where the resulting virtual machine will be created. This may be relative or absolute. If relative, the path is relative to the working directory when packer is executed. This directory must not exist or be empty prior to running the builder. By default this is output-BUILDNAME where \"BUILDNAME\" is the name of the build.",
				Type:        schema.TypeString,
			},

			"headless": &schema.Schema{
				Optional:    true,
				Description: "Packer defaults to building VMware virtual machines by launching a GUI that shows the console of the machine being built. When this value is set to true, the machine will start without a console. For VMware machines, Packer will output VNC connection information in case you need to connect to the console to debug the build process.",
				Type:        schema.TypeBool,
			},

			"vnc_bind_address": &schema.Schema{
				Optional:    true,
				Description: "The IP address that should be binded to for VNC. By default packer will use 127.0.0.1 for this. If you wish to bind to all interfaces use 0.0.0.0.",
				Type:        schema.TypeString,
			},

			"vnc_port_min": &schema.Schema{
				Optional:    true,
				Description: "The minimum and maximum port to use for VNC access to the virtual machine. The builder uses VNC to type the initial boot_command. Because Packer generally runs in parallel, Packer uses a randomly chosen port in this range that appears available. By default this is 5900 to 6000. The minimum and maximum ports are inclusive.",
				Type:        schema.TypeInt,
			},

			"vnc_port_max": &schema.Schema{
				Optional:    true,
				Description: "The minimum and maximum port to use for VNC access to the virtual machine. The builder uses VNC to type the initial boot_command. Because Packer generally runs in parallel, Packer uses a randomly chosen port in this range that appears available. By default this is 5900 to 6000. The minimum and maximum ports are inclusive.",
				Type:        schema.TypeInt,
			},

			"vnc_disable_password": &schema.Schema{
				Optional:    true,
				Description: "Don't auto-generate a VNC password that is used to secure the VNC communication with the VM.",
				Type:        schema.TypeBool,
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

			"tools_upload_flavor": &schema.Schema{
				Optional:    true,
				Description: "The flavor of the VMware Tools ISO to upload into the VM. Valid values are darwin, linux, and windows. By default, this is empty, which means VMware tools won't be uploaded.",
				Type:        schema.TypeString,
			},

			"tools_upload_path": &schema.Schema{
				Optional:    true,
				Description: "The path in the VM to upload the VMware tools. This only takes effect if tools_upload_flavor is non-empty. This is a configuration template that has a single valid variable: Flavor, which will be the value of tools_upload_flavor. By default the upload path is set to {{.Flavor}}.iso.",
				Type:        schema.TypeString,
			},

			"vmx_data": &schema.Schema{
				Optional:    true,
				Description: "Arbitrary key/values to enter into the virtual machine VMX file. This is for advanced users who want to set properties such as memory, CPU, etc.",
				Type:        schema.TypeMap,
			},

			"vmx_data_post": &schema.Schema{
				Optional:    true,
				Description: "Identical to vmx_data, except that it is run after the virtual machine is shutdown, and before the virtual machine is exported.",
				Type:        schema.TypeMap,
			},

			"vmx_remove_ethernet_interfaces": &schema.Schema{
				Optional:    true,
				Description: "Remove all ethernet interfaces from the VMX file after building. This is for advanced users who understand the ramifications, but is useful for building Vagrant boxes since Vagrant will create ethernet interfaces when provisioning a box. Defaults to false.",
				Type:        schema.TypeBool,
			},

			"skip_compaction": &schema.Schema{
				Optional:    true,
				Description: "VMware-created disks are defragmented and compacted at the end of the build process using vmware-vdiskmanager. In certain rare cases, this might actually end up making the resulting disks slightly larger. If you find this to be the case, you can disable compaction using this configuration value. Defaults to false.",
				Type:        schema.TypeBool,
			},

			"source_path": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "Path to the source VMX file to clone.",
			},

			"vm_name": &schema.Schema{
				Optional:    true,
				Description: "This is the name of the VMX file for the new virtual machine, without the file extension. By default this is packer-BUILDNAME, where \"BUILDNAME\" is the name of the build.",
				Type:        schema.TypeString,
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
