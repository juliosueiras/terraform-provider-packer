package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func AnsibleLocalResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order for this provisioner to run in",
			},
			"extra_arguments": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of extra arguments to pass to the ansible command. By default, this is empty. These arguments will be passed through a shell and arguments should be quoted accordingly.",
			},

			"group_vars": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "a path to the directory containing ansible group variables on your local system to be copied to the remote machine. By default, this is empty.",
			},

			"host_vars": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "a path to the directory containing ansible host variables on your local system to be copied to the remote machine. By default, this is empty.",
			},

			"playbook_dir": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "a path to the complete ansible directory structure on your local system to be copied to the remote machine as the staging_directory before all other files and directories.",
			},
			"command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to invoke ansible. Defaults to \"ANSIBLE_FORCE_COLOR=1 PYTHONUNBUFFERED=1 ansible-playbook\". Note, This disregards the value of -color when passed to packer build. To disable colors, set this to PYTHONUNBUFFERED=1 ansible-playbook",
			},

			"playbook_file": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The playbook file to be executed by ansible. This file must exist on your local system and will be uploaded to the remote machine.",
			},
			"playbook_paths": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of directories of playbook files on your local system. These will be uploaded to the remote machine under staging_directory/playbooks. By default, this is empty.",
			},

			"role_paths": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of paths to role directories on your local system. These will be uploaded to the remote machine under staging_directory/roles. By default, this is empty.",
			},

			"staging_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The directory where all the configuration of Ansible by Packer will be placed. By default this is /tmp/packer-provisioner-ansible-local/<uuid>, where <uuid> is replaced with a unique ID so that this provisioner can be run more than once. If you'd like to know the location of the staging directory in advance, you should set this to a known location. This directory doesn't need to exist but must have proper permissions so that the SSH user that Packer uses is able to create directories and write into this folder. If the permissions are not correct, use a shell provisioner prior to this to configure it properly.",
			},

			"clean_staging_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If set to true, the content of the staging_directory will be removed after executing ansible. By default, this is set to false.",
			},

			"inventory_file": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The inventory file to be used by ansible. This file must exist on your local system and will be uploaded to the remote machine",
			},

			"inventory_groups": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "A comma-separated list of groups to which packer will assign the host 127.0.0.1",
			},

			"galaxy_file": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "A requirements file which provides a way to install roles with the ansible-galaxy cli on the remote machine. By default, this is empty.",
			},
			"galaxycommand": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to invoke ansible-galaxy. By default, this is ansible-galaxy.",
			},
		},
	}
}
