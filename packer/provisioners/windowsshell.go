package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func WindowsShellResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order for this provisioner to run in",
			},
			"environment_vars": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of key/value pairs to inject prior to the execute_command. The format should be key=value. Packer injects some environmental variables by default into the environment, as well, which are covered in the section below.",
			},
			"script": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to a script to upload and execute in the machine. This path can be absolute or relative. If it is relative, it is relative to the working directory when Packer is executed.",
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

			"remote_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path where the script will be uploaded to in the machine. This defaults to \"c:/Windows/Temp/script.bat\". This value must be a writable location and any parent directories must already exist.",
			},
			"binary": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, specifies that the script(s) are binary files, and Packer should therefore not convert Windows line endings to Unix line endings (if there are any). By default this is false.",
			},

			"execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to use to execute the script. By default this is {{ .Vars }}\"{{ .Path }}\". The value of this is treated as template engine. There are two available variables: Path, which is the path to the script to run, and Vars, which is the list of environment_vars, if configured.",
			},

			"start_retry_timeout": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The amount of time to attempt to start the remote process. By default this is \"5m\" or 5 minutes. This setting exists in order to deal with times when SSH may restart, such as a system reboot. Set this to a higher value if reboots take a longer amount of time.",
			},
		},
	}
}
