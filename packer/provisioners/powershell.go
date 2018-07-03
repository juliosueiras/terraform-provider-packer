package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func PowerShellResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
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
			"binary": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, specifies that the script(s) are binary files, and Packer should therefore not convert Windows line endings to Unix line endings (if there are any). By default this is false.",
			},
			"environment_vars": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of key/value pairs to inject prior to the execute_command. The format should be key=value. Packer injects some environmental variables by default into the environment, as well, which are covered in the section below. If you are running on AWS, Azure or Google Compute and would like to access the generated password that Packer uses to connect to the instance via WinRM, you can use the template variable {{.WinRMPassword}} to set this as an environment variable.",
			},

			"remote_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path where the PowerShell script will be uploaded to within the target build machine. This defaults to C:/Windows/Temp/script-UUID.ps1 where UUID is replaced with a dynamically generated string that uniquely identifies the script.  This setting allows users to override the default upload location. The value must be a writable location and any parent directories must already exist",
			},

			"remote_env_var_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Environment variables required within the remote environment are uploaded within a PowerShell script and then enabled by 'dot sourcing' the script immediately prior to execution of the main command or script.  The path the environment variables script will be uploaded to defaults to C:/Windows/Temp/packer-ps-env-vars-UUID.ps1 where UUID is replaced with a dynamically generated string that uniquely identifies the script.  This setting allows users to override the location the environment variable script is uploaded to. The value must be a writable location and any parent directories must already exist.",
			},

			"execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to use to execute the script",
			},

			"elevated_execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to use to execute the elevated script",
			},

			"start_retry_timeout": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The amount of time to attempt to start the remote process. By default this is \"5m\" or 5 minutes. This setting exists in order to deal with times when SSH may restart, such as a system reboot. Set this to a higher value if reboots take a longer amount of time.",
			},

			"elevated_env_var_format": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"elevated_user": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "If specified, the PowerShell script will be run with elevated privileges using the given Windows user. If you are running a build on AWS, Azure or Google Compute and would like to run using the generated password that Packer uses to connect to the instance via WinRM, you may do so by using the template variable {{.WinRMPassword}}.",
			},

			"elevated_password": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "If specified, the PowerShell script will be run with elevated privileges using the given Windows user. If you are running a build on AWS, Azure or Google Compute and would like to run using the generated password that Packer uses to connect to the instance via WinRM, you may do so by using the template variable {{.WinRMPassword}}.",
			},

			"valid_exit_codes": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Valid exit codes for the script. By default this is just 0.",
			},
		},
	}
}
