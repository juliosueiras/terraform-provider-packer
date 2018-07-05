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
				Optional: true,
				Type:     schema.TypeString,
			},

			"http_port_min": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"http_port_max": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"floppy_files": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"floppy_dirs": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"boot_wait": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"boot_command": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"format": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"export_opts": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"output_directory": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"headless": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"vrdp_bind_address": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"vrdp_port_min": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"vrdp_port_max": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"shutdown_command": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"shutdown_timeout": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"post_shutdown_delay": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"communicator": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"vboxmanage": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
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
				Optional: true,
				Type:     schema.TypeList,
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
				Optional: true,
				Type:     schema.TypeString,
			},

			"disk_size": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"guest_additions_mode": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"guest_additions_path": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"guest_additions_sha256": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"guest_additions_url": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"guest_os_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"hard_drive_discard": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"hard_drive_interface": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"sata_port_count": &schema.Schema{
				Optional: true,
				Type:     schema.TypeInt,
			},

			"hard_drive_nonrotational": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"iso_interface": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"keep_registered": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"skip_export": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"vm_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}
