package provisioners

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func ConvergeResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order for this provisioner to run in",
			},
			"override": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.ValidateJsonString,
				Optional:     true,
			},
			"bootstrap": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Set to allow the provisioner to download the latest Converge bootstrap script and the specified version of Converge from the internet.",
			},

			"version": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Set to a released Converge version for bootstrap.",
			},

			"bootstrap_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "the command used to bootstrap Converge. This has various configuration template variables available.",
			},

			"prevent_bootstrap_sudo": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "stop Converge from bootstrapping with administrator privileges via sudo",
			},

			"module_dir": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Module directories to transfer to the remote host for execution. See below for the specification.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": &schema.Schema{
							Required:    true,
							Type:        schema.TypeString,
							Description: "the path to the folder on the local machine.",
						},
						"destination": &schema.Schema{
							Required:    true,
							Type:        schema.TypeString,
							Description: "the path to the folder on the remote machine. Parent directories will not be created; use the shell module to do this",
						},
						"exclude": &schema.Schema{
							Required:    true,
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "files and directories to exclude from transfer.",
						},
					},
				},
			},

			"module": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "Path (or URL) to the root module that Converge will apply",
			},

			"working_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The directory that Converge will change to before execution.",
			},

			"params": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "parameters to pass into the root module.",
			},

			"execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "the command used to execute Converge. This has various configuration template variables available.",
			},

			"prevent_sudo": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "stop Converge from running with administrator privileges via sudo",
			},
		},
	}
}
