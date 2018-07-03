package provisioners

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func PuppetMasterlessResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order for this provisioner to run in",
			},

			"guest_os_type": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"unix", "windows"}, true),
				Description:  "The remote host's OS type ('windows' or 'unix') to tailor command-line and path separators. (default: unix).",
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
				Description: "Additional options to pass to the Puppet command. This allows for customization of execute_command without having to completely replace or subsume its contents, making forward-compatible customizations much easier to maintain.  \nThis string is lazy-evaluated so one can incorporate logic driven by template variables as well as private elements of ExecuteTemplate (see source: provisioner/puppet-masterless/provisioner.go). [ {{if ne \"{{user environment}}\" \"\"}}--environment={{user environment}}{{end}}, {{if ne \".ModulePath\" \"\"}}--modulepath=\"{{.ModulePath}}{{.ModulePathJoiner}}$(puppet config print {{if ne \"{{user `environment`}}\" \"\"}}--environment={{user `environment`}}{{end}} modulepath)\"{{end}} ]",
			},
			"facter": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "Additional facts to make available to the Puppet run.",
			},

			"hiera_config_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Local path to self-contained Hiera data to be uploaded. NOTE: If you need data directories they must be previously transferred with a File provisioner.",
			},

			"ignore_exit_codes": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, Packer will ignore failures.",
			},

			"module_paths": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Array of local module directories to be uploaded",
			},

			"manifest_file": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "This is either a path to a puppet manifest (.pp file) or a directory containing multiple manifests that puppet will apply (the \"main manifest\"). These file(s) must exist on your local system and will be uploaded to the remote machine.",
			},

			"manifest_dir": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Local directory with manifests to be uploaded. This is useful if your main manifest uses imports, but the directory might not contain the manifest_file itself.",
			},

			"prevent_sudo": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "On Unix platforms Puppet is typically invoked with sudo. If true, it will be omitted. (default: false)",
			},

			"puppet_bin_dir": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to the Puppet binary. Ideally the program should be on the system (unix: $PATH, windows: %PATH%), but some builders (eg. Docker) do not run profile-setup scripts and therefore PATH might be empty or minimal.",
			},

			"staging_directory": &schema.Schema{
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
