package provisioners

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func AnsibleResource() *schema.Resource {
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

			"extra_arguments": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Extra arguments to pass to Ansible. These arguments will not be passed through a shell and arguments should not be quoted.",
			},

			"ansible_env_vars": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Environment variables to set before running Ansible.",
			},

			"playbook_file": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The playbook to be run by Ansible.",
			},

			"groups": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The groups into which the Ansible host should be placed. When unspecified, the host is not associated with any groups.",
			},

			"empty_groups": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The groups which should be present in inventory file but remain empty.",
			},

			"host_alias": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The alias by which the Ansible host should be known. Defaults to default. This setting is ignored when using a custom inventory file.",
			},

			"user": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The ansible_user to use. Defaults to the user running packer.",
			},

			"local_port": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The port on which to attempt to listen for SSH connections. This value is a starting point. The provisioner will attempt listen for SSH connections on the first available of ten ports, starting at local_port. A system-chosen port is used when local_port is missing or empty",
			},

			"ssh_host_key_file": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The SSH key that will be used to run the SSH server on the host machine to forward commands to the target machine. Ansible connects to this server and will validate the identity of the server using the system known_hosts. The default behavior is to generate and use a onetime key. Host key checking is disabled via the ANSIBLE_HOST_KEY_CHECKING environment variable if the key is generated.",
			},

			"ssh_authorized_key_file": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The SSH public key of the Ansible ssh_user. The default behavior is to generate and use a onetime key. If this key is generated, the corresponding private key is passed to ansible-playbook with the --private-key option.",
			},

			"sftp_command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to run on the machine being provisioned by Packer to handle the SFTP protocol that Ansible will use to transfer files. The command should read and write on stdin and stdout, respectively. Defaults to /usr/lib/sftp-server -e.",
			},
			"command": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The command to invoke ansible. Defaults to ansible-playbook.",
			},

			"skip_version_check": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Check if ansible is installed prior to running. Set this to true, for example, if you're going to install ansible during the packer run.",
			},
			"inventory_directory": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The directory in which to place the temporary generated Ansible inventory file. By default, this is the system-specific temporary file location. The fully-qualified name of this temporary file will be passed to the -i argument of the ansible command when this provisioner runs ansible. Specify this if you have an existing inventory directory with host_vars group_vars that you would like to use in the playbook that this provisioner will run.",
			},

			"inventory_file": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The inventory file to use during provisioning. When unspecified, Packer will create a temporary inventory file and will use the host_alias.",
			},
		},
	}
}
