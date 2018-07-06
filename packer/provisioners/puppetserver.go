package provisioners

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func PuppetServerResource() *schema.Resource {
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

			"client_cert_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to the directory on your disk that contains the client certificate for the node. This defaults to nothing, in which case a client cert won't be uploaded",
			},

			"client_private_key_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: " Path to the directory on your disk that contains the client private key for the node. This defaults to nothing, in which case a client private key won't be uploaded.",
			},

			"execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command-line to execute Puppet. This also has various configuration template variables available.",
			},

			"extra_arguments": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Additional options to pass to the Puppet command. This allows for customization of execute_command without having to completely replace or subsume its contents, making forward-compatible customizations much easier to maintain.\n  This string is lazy-evaluated so one can incorporate logic driven by template variables as well as private elements of ExecuteTemplate (see source: provisioner/puppet-server/provisioner.go). [ {{if ne \"{{user environment}}\" \"\"}}--environment={{user environment}}{{end}} ]",
			},
			"facter": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "Additional facts to make available to the Puppet run",
			},

			"guest_os_type": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"unix", "windows"}, true),
				Description:  "The remote host's OS type ('windows' or 'unix') to tailor command-line and path separators. (default: unix)",
			},

			"ignore_exit_codes": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, Packer will ignore failures",
			},

			"prevent_sudo": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: " On Unix platforms Puppet is typically invoked with sudo. If true, it will be omitted. (default: false)",
			},

			"puppet_bin_dir": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to the Puppet binary. Ideally the program should be on the system (unix: $PATH, windows: %PATH%), but some builders (eg. Docker) do not run profile-setup scripts and therefore PATH might be empty or minimal.",
			},

			"puppet_node": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the node. If this isn't set, the fully qualified domain name will be used.",
			},

			"puppet_server": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Hostname of the Puppet server. By default \"puppet\" will be used.",
			},

			"staging_dir": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Directory to where uploaded files will be placed (unix: \"/tmp/packer-puppet-masterless\", windows: \"%SYSTEMROOT%/Temp/packer-puppet-masterless\"). It doesn't need to pre-exist, but the parent must have permissions sufficient for the account Packer connects as to create directories and write files. Use a Shell provisioner to prepare the way if needed.",
			},

			"working_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Directory from which execute_command will be run. If using Hiera files with relative paths, this option can be helpful. (default: staging_directory)",
			},
		},
	}
}
