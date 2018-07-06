package provisioners

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func SaltMasterlessResource() *schema.Resource {
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

			"skip_bootstrap": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "By default the salt provisioner runs salt bootstrap to install salt. Set this to true to skip this step.",
			},

			"bootstrap_args": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Arguments to send to the bootstrap script. Usage is somewhat documented on github, but the script itself has more detailed usage instructions. By default, no arguments are sent to the script.",
			},

			"disable_sudo": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "By default, the bootstrap install command is prefixed with sudo. When using a Docker builder, you will likely want to pass true since sudo is often not pre-installed.",
			},

			"custom_state": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "A state to be run instead of state.highstate. Defaults to state.highstate if unspecified.",
			},

			"minion_config": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to your local minion config file. This will be uploaded to the /etc/salt on the remote. This option overrides the remote_state_tree or remote_pillar_roots options.",
			},

			"grains_file": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to your local grains file. This will be uploaded to /etc/salt/grains on the remote.",
			},

			"local_state_tree": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The path to your local state tree. This will be uploaded to the remote_state_tree on the remote.",
			},

			"local_pillar_roots": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to your local pillar roots. This will be uploaded to the remote_pillar_roots on the remote.",
			},

			"remote_state_tree": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to your remote state tree. default: /srv/salt. This option cannot be used with minion_config.",
			},

			"remote_pillar_roots": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to your remote pillar roots. default: /srv/pillar. This option cannot be used with minion_config.",
			},

			"temp_config_dir": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Where your local state tree will be copied before moving to the /srv/salt directory. Default is /tmp/salt.",
			},

			"no_exit_on_failure": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Packer will exit if the salt-call command fails. Set this option to true to ignore Salt failures.",
			},

			"log_level": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Set the logging level for the salt-call run",
			},

			"salt_call_args": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Additional arguments to pass directly to salt-call. See salt-call documentation for more information. By default no additional arguments (besides the ones Packer generates) are passed to salt-call.",
			},

			"salt_bin_dir": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to the salt-call executable. Useful if it is not on the PATH.",
			},

			"guest_os_type": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"unix", "windows"}, true),
				Description:  "The target guest OS type, either \"unix\" or \"windows\"",
			},
		},
	}
}
