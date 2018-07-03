package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func ShellResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order for this provisioner to run in",
			},
			"inline_shebang": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The shebang value to use when running commands specified by inline. By default, this is /bin/sh -e. If you're not using inline, then this configuration has no effect. Important: If you customize this, be sure to include something like the -e flag, otherwise individual steps failing won't fail the provisioner.",
			},

			"script": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to a script to upload and execute in the machine. This path can be absolute or relative. If it is relative, it is relative to the working directory when Packer is executed.  ",
			},
			"inline": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "This is an array of commands to execute. The commands are concatenated by newlines and turned into a single file, so they are all executed within the same context. This allows you to change directories in one command and use something in the directory in the next and so on. Inline scripts are the easiest way to pull off simple tasks within the machine.",
			},
			"scripts": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of scripts to execute. The scripts will be uploaded and executed in the order specified. Each script is executed in isolation, so state such as variables from one script won't carry on to the next.",
			},
			"environment_vars": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of key/value pairs to inject prior to the execute_command. The format should be key=value. Packer injects some environmental variables by default into the environment, as well, which are covered in the section below.",
			},

			"remote_folder": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The folder where the uploaded script will reside on the machine. This defaults to '/tmp'.",
			},

			"remote_file": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The filename the uploaded script will have on the machine. This defaults to 'script_nnn.sh'",
			},

			"remote_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The full path to the uploaded script will have on the machine. By default this is remote_folder/remote_file, if set this option will override both remote_folder and remote_file.",
			},

			"execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to use to execute the script. By default this is chmod +x {{ .Path }}; {{ .Vars }} {{ .Path }}. The value of this is treated as configuration template. There are two available variables: Path, which is the path to the script to run, and Vars, which is the list of environment_vars, if configured.",
			},

			"start_retry_timeout": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The amount of time to attempt to start the remote process. By default this is 5m or 5 minutes. This setting exists in order to deal with times when SSH may restart, such as a system reboot. Set this to a higher value if reboots take a longer amount of time.",
			},

			"skip_clean": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, specifies that the helper scripts uploaded to the system will not be removed by Packer. This defaults to false (clean scripts from the system).",
			},
			"binary": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, specifies that the script(s) are binary files, and Packer should therefore not convert Windows line endings to Unix line endings (if there are any). By default this is false",
			},

			"expect_disconnect": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Defaults to false. Whether to error if the server disconnects us. A disconnect might happen if you restart the ssh server or reboot the host.",
			},
		},
	}
}
