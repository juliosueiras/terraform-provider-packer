package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func ShellLocalResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"only": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "For specifying a builder input",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"except": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "For excluding a builder input",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"keep_input_artifact": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
			"pipeline": &schema.Schema{
				Optional:    true,
				MaxItems:    1,
				Type:        schema.TypeList,
				Description: "Create a sequence definition",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"set": &schema.Schema{
							Required:    true,
							Type:        schema.TypeInt,
							Description: "The set(group) that is under",
						},
						"order": &schema.Schema{
							Required:    true,
							Type:        schema.TypeInt,
							Description: "The order to run in",
						},
					},
				},
			},
			"command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is a single command to execute. It will be written to a temporary file and run using the execute_command call below.",
			},
			"inline": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "This is an array of commands to execute. The commands are concatenated by newlines and turned into a single file, so they are all executed within the same context. This allows you to change directories in one command and use something in the directory in the next and so on. Inline scripts are the easiest way to pull off simple tasks within the machine.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"script": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to a script to execute. This path can be absolute or relative. If it is relative, it is relative to the working directory when Packer is executed.",
			},
			"scripts": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "An array of scripts to execute. The scripts will be executed in the order specified. Each script is executed in isolation, so state such as variables from one script won't carry on to the next.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"environment_vars": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "An array of key/value pairs to inject prior to the execute_command. The format should be key=value. Packer injects some environmental variables by default into the environment, as well",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"execute_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "The command used to execute the script. By default this is [\"/bin/sh\", \"-c\", \"{{.Vars}}, \"{{.Script}}\"] on unix and [\"cmd\", \"/c\", \"{{.Vars}}\", \"{{.Script}}\"] on windows. This is treated as a template engine. There are two available variables: Script, which is the path to the script to run, and Vars, which is the list of environment_vars, if configured. If you choose to set this option, make sure that the first element in the array is the shell program you want to use (for example, \"sh\" or \"/usr/local/bin/zsh\" or even \"powershell.exe\" although anything other than a flavor of the shell command language is not explicitly supported and may be broken by assumptions made within Packer). It's worth noting that if you choose to try to use shell-local for Powershell or other Windows commands, the environment variables will not be set properly for your environment.  For backwards compatibility, execute_command will accept a string instead of an array of strings. If a single string or an array of strings with only one element is provided, Packer will replicate past behavior by appending your execute_command to the array of strings [\"sh\", \"-c\"]. For example, if you set \"execute_command\": \"foo bar\", the final execute_command that Packer runs will be [\"sh\", \"-c\", \"foo bar\"]. If you set \"execute_command\": [\"foo\", \"bar\"], the final execute_command will remain [\"foo\", \"bar\"].  Again, the above is only provided as a backwards compatibility fix; we strongly recommend that you set execute_command as an array of strings.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"inline_shebang": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The shebang value to use when running commands specified by inline. By default, this is /bin/sh -e. If you're not using inline, then this configuration has no effect. Important: If you customize this, be sure to include something like the -e flag, otherwise individual steps failing won't fail the provisioner.",
			},
			"use_linux_pathing": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "This is only relevant to windows hosts. If you are running Packer in a Windows environment with the Windows Subsystem for Linux feature enabled, and would like to invoke a bash script rather than invoking a Cmd script, you'll need to set this flag to true; it tells Packer to use the linux subsystem path for your script rather than the Windows path. (e.g. /mnt/c/path/to/your/file instead of C:/path/to/your/file). Please see the example below for more guidance on how to use this feature. If you are not on a Windows host, or you do not intend to use the shell-local post-processor to run a bash script, please ignore this option. If you set this flag to true, you still need to provide the standard windows path to the script when providing a script. This is a beta feature.",
			},
		},
	}
}
