package provisioners

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func ChefSoloResource() *schema.Resource {
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
			"chef_environment": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the chef_environment sent to the Chef server. By default this is empty and will not use an environment",
			},

			"config_template": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Path to a template that will be used for the Chef configuration file. By default Packer only sets configuration it needs to match the settings set in the provisioner configuration. If you need to set configurations that the Packer provisioner doesn't support, then you should use a custom configuration template. See the dedicated \"Chef Configuration\" section below for more details.",
			},

			"cookbook_paths": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "This is an array of paths to \"cookbooks\" directories on your local filesystem. These will be uploaded to the remote machine in the directory specified by the staging_directory. By default, this is empty.",
			},

			"roles_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to the \"roles\" directory on your local filesystem. These will be uploaded to the remote machine in the directory specified by the staging_directory. By default, this is empty.",
			},

			"data_bags_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: " The path to the \"data_bags\" directory on your local filesystem. These will be uploaded to the remote machine in the directory specified by the staging_directory. By default, this is empty.",
			},

			"encrypted_data_bag_secret_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to the file containing the secret for encrypted data bags. By default, this is empty, so no secret will be available.",
			},

			"environments_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to the \"environments\" directory on your local filesystem. These will be uploaded to the remote machine in the directory specified by the staging_directory. By default, this is empty",
			},

			"execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command used to execute Chef. This has various configuration template variables available. See below for more information",
			},

			"install_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command used to install Chef. This has various configuration template variables available. See below for more information.",
			},
			"json": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "An arbitrary mapping of JSON that will be available as node attributes while running Chef.",
			},

			"remote_cookbook_paths": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "A list of paths on the remote machine where cookbooks will already exist. These may exist from a previous provisioner or step. If specified, Chef will be configured to look for cookbooks here. By default, this is empty.",
			},

			"prevent_sudo": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "By default, the configured commands that are executed to install and run Chef are executed with sudo. If this is true, then the sudo will be omitted. This has no effect when guest_os_type is window",
			},

			"run_list": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The run list for Chef. By default this is empty",
			},

			"skip_install": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, Chef will not automatically be installed on the machine using the Chef omnibus installers.",
			},

			"staging_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is the directory where all the configuration of Chef by Packer will be placed. By default this is \"/tmp/packer-chef-solo\" when guest_os_type unix and \"$env:TEMP/packer-chef-solo\" when windows. This directory doesn't need to exist but must have proper permissions so that the user that Packer uses is able to create directories and write into this folder. If the permissions are not correct, use a shell provisioner prior to this to configure it properly.",
			},

			"guest_os_type": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"unix", "windows"}, true),
				Description:  "The target guest OS type, either \"unix\" or \"windows\". Setting this to \"windows\" will cause the provisioner to use Windows friendly paths and commands. By default, this is \"unix\".",
			},

			"version": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The version of Chef to be installed. By default this is empty which will install the latest version of Chef.",
			},
		},
	}
}
