package builders

import "github.com/hashicorp/terraform/helper/schema"

func LXCResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},
			"config_file": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The path to the lxc configuration file.",
			},

			"output_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The directory in which to save the exported tar.gz. Defaults to output-<BuildName> in the current directory.",
			},

			"container_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the LXC container. Usually stored in /var/lib/lxc/containers/<container_name>. Defaults to packer-<BuildName>.",
			},

			"command_wrapper": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Allows you to specify a wrapper command, such as ssh so you can execute packer builds on a remote host. Defaults to Empty.",
			},

			"init_timeout": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The timeout in seconds to wait for the the container to start. Defaults to 20 seconds.",
			},

			"create_options": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Options to pass to lxc-create. For instance, you can specify a custom LXC container configuration file with [\"-f\", \"/path/to/lxc.conf\"]. Defaults to []. See man 1 lxc-create for available options.",
			},

			"start_options": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Options to pass to lxc-start. For instance, you can override parameters from the LXC container configuration file via [\"--define\", \"KEY=VALUE\"]. Defaults to []. See man 1 lxc-start for available options.",
			},

			"attach_options": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Options to pass to lxc-attach. For instance, you can prevent the container from inheriting the host machine's environment by specifying [\"--clear-env\"]. Defaults to []. See man 1 lxc-attach for available options.",
			},

			"template_name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The LXC template name to use.",
			},

			"template_parameters": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Options to pass to the given lxc-template command, usually located in /usr/share/lxc/templates/lxc-<template_name>. Note: This gets passed as ARGV to the template command. Ensure you have an array of strings, as a single string with spaces probably won't work. Defaults to []",
			},

			"template_environment_vars": &schema.Schema{
				Required:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Environmental variables to use to build the template with.",
			},

			"target_runlevel": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "The minimum run level to wait for the container to reach. Note some distributions (Ubuntu) simulate run levels and may report 5 rather than 3.",
			},
		},
	}
}
