package provisioners

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func ShellLocalResource() *schema.Resource {
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
			"command": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "This is a single command to execute. It will be written to a temporary file and run using the execute_command call below. If you are building a windows vm on AWS, Azure or Google Compute and would like to access the generated password that Packer uses to connect to the instance via WinRM, you can use the template variable {{.WinRMPassword}} to set this as an environment variable.",
			},
			"use_linux_pathing": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "This is only relevant to windows hosts. If you are running Packer in a Windows environment with the Windows Subsystem for Linux feature enabled, and would like to invoke a bash script rather than invoking a Cmd script, you'll need to set this flag to true; it tells Packer to use the linux subsystem path for your script rather than the Windows path. (e.g. /mnt/c/path/to/your/file instead of C:/path/to/your/file). Please see the example below for more guidance on how to use this feature. If you are not on a Windows host, or you do not intend to use the shell-local provisioner to run a bash script, please ignore this option.",
			},
			"inline_shebang": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The shebang value to use when running commands specified by inline. By default, this is /bin/sh -e. If you're not using inline, then this configuration has no effect. Important: If you customize this, be sure to include something like the -e flag, otherwise individual steps failing won't fail the provisioner.",
			},
			"script": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to a script to execute. This path can be absolute or relative. If it is relative, it is relative to the working directory when Packer is executed.",
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
				Description: "An array of scripts to execute. The scripts will be executed in the order specified. Each script is executed in isolation, so state such as variables from one script won't carry on to the next.",
			},
			"environment_vars": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of key/value pairs to inject prior to the execute_command. The format should be key=value. Packer injects some environmental variables by default into the environment, as well, which are covered in the section below. If you are building a windows vm on AWS, Azure or Google Compute and would like to access the generated password that Packer uses to connect to the instance via WinRM, you can use the template variable {{.WinRMPassword}} to set this as an environment variable. For example: \"environment_vars\": \"WINRMPASS={{.WinRMPassword}}\"",
			},
			"execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The command used to execute the script. By default t \"/bin/sh\", \"-c\", \"{{.Vars}}, \"{{.Script}}\"] on unix and [\"cmd\", \"/c\", \"{{.Vars}}\", \"{{.Script}}\"] on windows. This is treated as a template engine. There are two available variables: Script, which is the path to the script to run, and Vars, which is the list of environment_vars, if configured.\n If you choose to set this option, make sure that the first element in the array is the shell program you want to use (for example, \"sh\"), and a later element in the array must be {{.Script}}.\n This option provides you a great deal of flexibility. You may choose to provide your own shell program, for example \"/usr/local/bin/zsh\" or even \"powershell.exe\". However, with great power comes great responsibility - these commands are not officially supported and things like environment variables may not work if you use a different shell than the default. \n For backwards compatibility, you may also use {{.Command}}, but it is decoded the same way as {{.Script}}. We recommend using {{.Script}} for the sake of clarity, as even when you set only a single command to run, Packer writes it to a temporary file and then runs it as a script.  \nIf you are building a windows vm on AWS, Azure or Google Compute and would like to access the generated password that Packer uses to connect to the instance via WinRM, you can use the template variable {{.WinRMPassword}} to set this as an environment variable.",
			},
		},
	}
}
