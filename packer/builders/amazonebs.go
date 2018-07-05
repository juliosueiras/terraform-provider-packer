package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func AmazonEBSResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},
			"communicator": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"ssh": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     communicators.SSHCommunicatorResource(),
			},
			"winrm": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     communicators.WinRMCommunicatorResource(),
			},
			"access_key": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The access key used to communicate with AWS. Learn how to set this.",
			},
			"ami_name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name of the resulting AMI that will appear when managing AMIs in the AWS console or via APIs. This must be unique. To help make this unique, use a function like timestamp (see template engine for more info)",
			},
			"instance_type": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The EC2 instance type to use while building the AMI, such as t2.small.",
			},
			"region": &schema.Schema{
				Optional:    true,
				Description: "The name of the region, such as us-east-1, in which to launch the EC2 instance to create the AMI.",
				Type:        schema.TypeString,
			},
			"secret_key": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The secret key used to communicate with AWS. Learn how to set this.",
			},
			"source_ami": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The initial AMI used as a base for the newly created machine. source_ami_filter may be used instead to populate this automatically.",
			},
			"ami_block_device_mapping": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Add one or more block device mappings to the AMI. These will be attached when booting a new instance from your AMI. Your options here may vary depending on the type of VM you use.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_on_termination": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the EBS volume is deleted on instance termination. Default false. NOTE: If this value is not explicitly set to true and volumes are not cleaned up by an alternative method, additional volumes will accumulate after every build.",
						},
						"device_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The device name exposed to the instance (for example, /dev/sdh or xvdh). Required when specifying volume_size. ",
						},
						"encrypted": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to encrypt the volume or not",
						},
						"kms_key_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.",
						},
						"iops": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The number of I/O operations per second (IOPS) that the volume supports. See the documentation on IOPs for more information",
						},
						"no_device": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Suppresses the specified device included in the block device mapping of the AMI",
						},
						"snapshot_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the snapshot",
						},
						"virtual_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The virtual device name. See the documentation on Block Device Mapping for more information",
						},
						"volume_size": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The size of the volume, in GiB. Required if not specifying a snapshot_id",
						},
						"volume_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The volume type. gp2 for General Purpose (SSD) volumes, io1 for Provisioned IOPS (SSD) volumes, and standard for Magnetic volumes",
						},
					},
				},
			},
			"ami_description": &schema.Schema{
				Optional:    true,
				Description: "The description to set for the resulting AMI(s). By default this description is empty. This is a template engine, see Build template data for more information.",
				Type:        schema.TypeString,
			},
			"ami_groups": &schema.Schema{
				Optional:    true,
				Description: "A list of groups that have access to launch the resulting AMI(s). By default no groups have permission to launch the AMI. all will make the AMI publicly accessible. AWS currently doesn't accept any value other than all.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"ami_product_codes": &schema.Schema{
				Optional:    true,
				Description: "A list of product codes to associate with the AMI. By default no product codes are associated with the AMI.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"ami_regions": &schema.Schema{
				Optional:    true,
				Description: "A list of regions to copy the AMI to. Tags and attributes are copied along with the AMI. AMI copying takes time depending on the size of the AMI, but will generally take many minutes.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"ami_users": &schema.Schema{
				Optional:    true,
				Description: "A list of account IDs that have access to launch the resulting AMI(s). By default no additional users other than the user creating the AMI has permissions to launch it.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"ami_virtualization_type": &schema.Schema{
				Optional:    true,
				Description: "The type of virtualization for the AMI you are building. This option must match the supported virtualization type of source_ami. Can be paravirtual or hvm.",
				Type:        schema.TypeString,
			},
			"associate_public_ip_address": &schema.Schema{
				Optional:    true,
				Description: "If using a non-default VPC, public IP addresses are not provided by default. If this is toggled, your new instance will get a Public IP.",
				Type:        schema.TypeBool,
			},
			"availability_zone": &schema.Schema{
				Optional:    true,
				Description: "Destination availability zone to launch instance in. Leave this empty to allow Amazon to auto-assign.",
				Type:        schema.TypeString,
			},
			"custom_endpoint_ec2": &schema.Schema{
				Optional:    true,
				Description: "This option is useful if you use a cloud provider whose API is compatible with aws EC2. Specify another endpoint like this https://ec2.custom.endpoint.com.",
				Type:        schema.TypeString,
			},
			"disable_stop_instance": &schema.Schema{
				Optional:    true,
				Description: "Packer normally stops the build instance after all provisioners have run. For Windows instances, it is sometimes desirable to run Sysprep which will stop the instance for you. If this is set to true, Packer will not stop the instance and will wait for you to stop it manually. You can do this with a windows-shell provisioner.",
				Type:        schema.TypeBool,
			},
			"ebs_optimized": &schema.Schema{
				Optional:    true,
				Description: "Mark instance as EBS Optimized. Default false.",
				Type:        schema.TypeBool,
			},
			"ena_support": &schema.Schema{
				Optional:    true,
				Description: "Enable enhanced networking (ENA but not SriovNetSupport) on HVM-compatible AMIs. If true, add ec2:ModifyInstanceAttribute to your AWS IAM policy. Note: you must make sure enhanced networking is enabled on your instance. See Amazon's documentation on enabling enhanced networking. Default false.",
				Type:        schema.TypeBool,
			},
			"enable_t2_unlimited": &schema.Schema{
				Optional:    true,
				Description: "Enabling T2 Unlimited allows the source instance to burst additional CPU beyond its available CPU Credits for as long as the demand exists. This is in contrast to the standard configuration that only allows an instance to consume up to its available CPU Credits. See the AWS documentation for T2 Unlimited and the 'T2 Unlimited Pricing' section of the Amazon EC2 On-Demand Pricing document for more information. By default this option is disabled and Packer will set up a T2 Standard instance instead.",
				Type:        schema.TypeBool,
			},
			"force_deregister": &schema.Schema{
				Optional:    true,
				Description: "Force Packer to first deregister an existing AMI if one with the same name already exists. Default false.",
				Type:        schema.TypeBool,
			},
			"force_delete_snapshot": &schema.Schema{
				Optional:    true,
				Description: "Force Packer to delete snapshots associated with AMIs, which have been deregistered by force_deregister. Default false",
				Type:        schema.TypeBool,
			},
			"encrypt_boot": &schema.Schema{
				Optional:    true,
				Description: "Instruct packer to automatically create a copy of the AMI with an encrypted boot volume (discarding the initial unencrypted AMI in the process). Packer will always run this operation, even if the base AMI has an encrypted boot volume to start with. Default false.",
				Type:        schema.TypeBool,
			},
			"kms_key_id": &schema.Schema{
				Optional:    true,
				Description: "The ID of the KMS key to use for boot volume encryption. This only applies to the main region, other regions where the AMI will be copied will be encrypted by the default EBS KMS key.",
				Type:        schema.TypeString,
			},
			"iam_instance_profile": &schema.Schema{
				Optional:    true,
				Description: "The name of an IAM instance profile to launch the EC2 instance with",
				Type:        schema.TypeString,
			},
			"launch_block_device_mapping": &schema.Schema{
				Optional:    true,
				Description: "Add one or more block devices before the Packer build starts. If you add instance store volumes or EBS volumes in addition to the root device volume, the created AMI will contain block device mapping information for those volumes. Amazon creates snapshots of the source instance's root volume and any other EBS volumes described here. When you launch an instance from this new AMI, the instance automatically launches with these additional volumes, and will restore them from snapshots taken from the source instance.",
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_on_termination": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the EBS volume is deleted on instance termination. Default false. NOTE: If this value is not explicitly set to true and volumes are not cleaned up by an alternative method, additional volumes will accumulate after every build.",
						},
						"device_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The device name exposed to the instance (for example, /dev/sdh or xvdh). Required when specifying volume_size. ",
						},
						"encrypted": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to encrypt the volume or not",
						},
						"kms_key_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.",
						},
						"iops": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The number of I/O operations per second (IOPS) that the volume supports. See the documentation on IOPs for more information",
						},
						"no_device": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Suppresses the specified device included in the block device mapping of the AMI",
						},
						"snapshot_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the snapshot",
						},
						"virtual_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The virtual device name. See the documentation on Block Device Mapping for more information",
						},
						"volume_size": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The size of the volume, in GiB. Required if not specifying a snapshot_id",
						},
						"volume_type": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The volume type. gp2 for General Purpose (SSD) volumes, io1 for Provisioned IOPS (SSD) volumes, and standard for Magnetic volumes",
						},
					},
				},
			},
			"mfa_code": &schema.Schema{
				Optional:    true,
				Description: "The MFA TOTP code. This should probably be a user variable since it changes all the time.",
				Type:        schema.TypeString,
			},
			"profile": &schema.Schema{
				Optional:    true,
				Description: "The profile to use in the shared credentials file for AWS. See Amazon's documentation on specifying profiles for more details.",
				Type:        schema.TypeString,
			},
			"region_kms_key_ids": &schema.Schema{
				Optional:    true,
				Description: "a map of regions to copy the ami to, along with the custom kms key id to use for encryption for that region. Keys must match the regions provided in ami_regions. If you just want to encrypt using a default ID, you can stick with kms_key_id and ami_regions. If you want a region to be encrypted with that region's default key ID, you can use an empty string \"\" instead of a key id in this map. (e.g. \"us-east-1\": \"\") However, you cannot use default key IDs if you are using this in conjunction with snapshot_users -- in that situation you must use custom keys.  ",
				Type:        schema.TypeMap,
			},
			"run_tags": &schema.Schema{
				Optional:    true,
				Description: "Tags to apply to the instance that is launched to create the AMI. These tags are not applied to the resulting AMI unless they're duplicated in tags. This is a template engine, see Build template data for more information.",
				Type:        schema.TypeMap,
			},
			"run_volume_tags": &schema.Schema{
				Optional:    true,
				Description: "Tags to apply to the volumes that are launched to create the AMI. These tags are not applied to the resulting AMI unless they're duplicated in tags. This is a template engine, see Build template data for more information.",
				Type:        schema.TypeMap,
			},
			"security_group_id": &schema.Schema{
				Optional:    true,
				Description: "The ID (not the name) of the security group to assign to the instance. By default this is not set and Packer will automatically create a new temporary security group to allow SSH access. Note that if this is specified, you must be sure the security group allows access to the ssh_port given below",
				Type:        schema.TypeString,
			},
			"security_group_ids": &schema.Schema{
				Optional:    true,
				Description: "A list of security groups as described above. Note that if this is specified, you must omit the security_group_id.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"temporary_security_group_source_cidr": &schema.Schema{
				Optional:    true,
				Description: "An IPv4 CIDR block to be authorized access to the instance, when packer is creating a temporary security group. The default is 0.0.0.0/0 (ie, allow any IPv4 source). This is only used when security_group_id or security_group_ids is not specified.",
				Type:        schema.TypeString,
			},
			"shutdown_behavior": &schema.Schema{
				Optional:    true,
				Description: "Automatically terminate instances on shutdown in case Packer exits ungracefully. Possible values are \"stop\" and \"terminate\", default is stop.",
				Type:        schema.TypeString,
			},
			"skip_region_validation": &schema.Schema{
				Optional:    true,
				Description: "Set to true if you want to skip validation of the ami_regions configuration option. Default false.",
				Type:        schema.TypeBool,
			},
			"snapshot_groups": &schema.Schema{
				Optional:    true,
				Description: "A list of groups that have access to create volumes from the snapshot(s). By default no groups have permission to create volumes from the snapshot(s). all will make the snapshot publicly accessible.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"snapshot_users": &schema.Schema{
				Optional:    true,
				Description: "A list of account IDs that have access to create volumes from the snapshot(s). By default no additional users other than the user creating the AMI has permissions to create volumes from the backing snapshot(s).",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"snapshot_tags": &schema.Schema{
				Optional:    true,
				Description: "Tags to apply to snapshot. They will override AMI tags if already applied to snapshot. This is a template engine, see Build template data for more information",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"source_ami_filter": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Filters used to populate the source_ami field",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filters": &schema.Schema{
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "filters used to select a source_ami. NOTE: This will fail unless exactly one AMI is returned. Any filter described in the docs for DescribeImages is valid.",
						},
						"owners": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "This scopes the AMIs to certain Amazon account IDs. This is helpful to limit the AMIs to a trusted third party, or to your own account.",
						},
						"most_recent": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Selects the newest created image when true. This is most useful for selecting a daily distro build.",
						},
					},
				},
			},
			"spot_price": &schema.Schema{
				Optional:    true,
				Description: "The maximum hourly price to pay for a spot instance to create the AMI. Spot instances are a type of instance that EC2 starts when the current spot price is less than the maximum price you specify. Spot price will be updated based on available spot instance capacity and current spot instance requests. It may save you some costs. You can set this to auto for Packer to automatically discover the best spot price or to \"0\" to use an on demand instance (default).",
				Type:        schema.TypeString,
			},
			"spot_price_auto_product": &schema.Schema{
				Optional:    true,
				Description: "equired if spot_price is set to auto. This tells Packer what sort of AMI you're launching to find the best spot price. This must be one of: Linux/UNIX, SUSE Linux, Windows, Linux/UNIX (Amazon VPC), SUSE Linux (Amazon VPC), Windows (Amazon VPC)",
				Type:        schema.TypeString,
			},
			"sriov_support": &schema.Schema{
				Optional:    true,
				Description: "Enable enhanced networking (SriovNetSupport but not ENA) on HVM-compatible AMIs. If true, add ec2:ModifyInstanceAttribute to your AWS IAM policy. Note: you must make sure enhanced networking is enabled on your instance. See Amazon's documentation on enabling enhanced networking. Default false.",
				Type:        schema.TypeBool,
			},
			"ssh_keypair_name": &schema.Schema{
				Optional:    true,
				Description: "If specified, this is the key that will be used for SSH with the machine. The key must match a key pair name loaded up into Amazon EC2. By default, this is blank, and Packer will generate a temporary keypair unless ssh_password is used. ssh_private_key_file or ssh_agent_auth must be specified when ssh_keypair_name is utilized.",
				Type:        schema.TypeString,
			},
			"ssh_interface": &schema.Schema{
				Optional:    true,
				Description: "One of public_ip, private_ip, public_dns or private_dns. If set, either the public IP address, private IP address, public DNS name or private DNS name will used as the host for SSH. The default behaviour if inside a VPC is to use the public IP address if available, otherwise the private IP address will be used. If not in a VPC the public DNS name will be used. Also works for WinRM.",
				Type:        schema.TypeString,
			},
			"subnet_id": &schema.Schema{
				Optional:    true,
				Description: "If using VPC, the ID of the subnet, such as subnet-12345def, where Packer will launch the EC2 instance. This field is required if you are using an non-default VPC.",
				Type:        schema.TypeString,
			},
			"tags": &schema.Schema{
				Optional:    true,
				Description: "Tags applied to the AMI and relevant snapshots. This is a template engine, see Build template data for more information.",
				Type:        schema.TypeMap,
			},
			"temporary_key_pair_name": &schema.Schema{
				Optional:    true,
				Description: "The name of the temporary key pair to generate. By default, Packer generates a name that looks like packer_<UUID>, where <UUID> is a 36 character unique identifier",
				Type:        schema.TypeString,
			},
			"token": &schema.Schema{
				Optional:    true,
				Description: "The access token to use. This is different from the access key and secret key. If you're not sure what this is, then you probably don't need it. This will also be read from the AWS_SESSION_TOKEN environmental variable.",
				Type:        schema.TypeString,
			},
			"user_data": &schema.Schema{
				Optional:    true,
				Description: "User data to apply when launching the instance. Note that you need to be careful about escaping characters due to the templates being JSON. It is often more convenient to use user_data_file, instead",
				Type:        schema.TypeString,
			},
			"user_data_file": &schema.Schema{
				Optional:    true,
				Description: "Path to a file that will be used for the user data when launching the instance.",
				Type:        schema.TypeString,
			},
			"vpc_id": &schema.Schema{
				Optional:    true,
				Description: "If launching into a VPC subnet, Packer needs the VPC ID in order to create a temporary security group within the VPC. Requires subnet_id to be set. If this field is left blank, Packer will try to get the VPC ID from the subnet_id.",
				Type:        schema.TypeString,
			},
			"windows_password_timeout": &schema.Schema{
				Optional:    true,
				Description: "The timeout for waiting for a Windows password for Windows instances. Defaults to 20 minutes. Example value: 10m",
				Type:        schema.TypeString,
			},
		},
	}
}
