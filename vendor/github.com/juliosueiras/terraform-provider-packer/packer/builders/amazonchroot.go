package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func AmazonChrootResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
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

			"ami_name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name of the resulting AMI that will appear when managing AMIs in the AWS console or via APIs. This must be unique. To help make this unique, use a function like timestamp (see template engine for more info)",
			},

			"ami_description": &schema.Schema{
				Optional:    true,
				Description: "The description to set for the resulting AMI(s). By default this description is empty. This is a template engine, see Build template data for more information.",
				Type:        schema.TypeString,
			},

			"ami_virtualization_type": &schema.Schema{
				Optional:    true,
				Description: "ami_virtualization_type (string) - The type of virtualization for the AMI you are building. This option is required to register HVM images. Can be \"paravirtual\" (default) or \"hvm\".  ",
				Type:        schema.TypeString,
			},

			"ami_users": &schema.Schema{
				Optional:    true,
				Description: "A list of account IDs that have access to launch the resulting AMI(s). By default no additional users other than the user creating the AMI has permissions to launch it.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"ami_groups": &schema.Schema{
				Optional:    true,
				Description: "A list of groups that have access to launch the resulting AMI(s). By default no groups have permission to launch the AMI. all will make the AMI publicly accessible.",
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

			"skip_region_validation": &schema.Schema{
				Optional:    true,
				Description: "Set to true if you want to skip validation of the ami_regions configuration option. Default false.",
				Type:        schema.TypeBool,
			},

			"tags": &schema.Schema{
				Optional:    true,
				Description: "Tags applied to the AMI. This is a template engine, see Build template data for more information.",
				Type:        schema.TypeMap,
			},

			"ena_support": &schema.Schema{
				Optional:    true,
				Description: "Enable enhanced networking (ENA but not SriovNetSupport) on HVM-compatible AMIs. If true, add ec2:ModifyInstanceAttribute to your AWS IAM policy. Note: you must make sure enhanced networking is enabled on your instance. See Amazon's documentation on enabling enhanced networking. Default false.",
				Type:        schema.TypeBool,
			},

			"sriov_support": &schema.Schema{
				Optional:    true,
				Description: "Enable enhanced networking (SriovNetSupport but not ENA) on HVM-compatible AMIs. If true, add ec2:ModifyInstanceAttribute to your AWS IAM policy. Note: you must make sure enhanced networking is enabled on your instance. See Amazon's documentation on enabling enhanced networking. Default false.",
				Type:        schema.TypeBool,
			},

			"force_deregister": &schema.Schema{
				Optional:    true,
				Description: "Force Packer to first deregister an existing AMI if one with the same name already exists. Default false.",
				Type:        schema.TypeBool,
			},

			"force_delete_snapshot": &schema.Schema{
				Optional:    true,
				Description: "Force Packer to delete snapshots associated with AMIs, which have been deregistered by force_deregister. Default false.",
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

			"region_kms_key_ids": &schema.Schema{
				Optional:    true,
				Description: "a map of regions to copy the ami to, along with the custom kms key id to use for encryption for that region. Keys must match the regions provided in ami_regions. If you just want to encrypt using a default ID, you can stick with kms_key_id and ami_regions. If you want a region to be encrypted with that region's default key ID, you can use an empty string \"\" instead of a key id in this map. (e.g. \"us-east-1\": \"\") However, you cannot use default key IDs if you are using this in conjunction with snapshot_users -- in that situation you must use custom keys.  ",
				Type:        schema.TypeMap,
			},

			"snapshot_tags": &schema.Schema{
				Optional:    true,
				Description: "Tags to apply to snapshot. They will override AMI tags if already applied to snapshot. This is a template engine, see Build template data for more information.",
				Type:        schema.TypeMap,
			},

			"snapshot_users": &schema.Schema{
				Optional:    true,
				Description: "A list of account IDs that have access to create volumes from the snapshot(s). By default no additional users other than the user creating the AMI has permissions to create volumes from the backing snapshot(s).",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"snapshot_groups": &schema.Schema{
				Optional:    true,
				Description: "A list of groups that have access to create volumes from the snapshot(s). By default no groups have permission to create volumes from the snapshot(s). all will make the snapshot publicly accessible.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"access_key": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The access key used to communicate with AWS. Learn how to set this.",
			},

			"custom_endpoint_ec2": &schema.Schema{
				Optional:    true,
				Description: "This option is useful if you use a cloud provider whose API is compatible with aws EC2. Specify another endpoint like this https://ec2.custom.endpoint.com.",
				Type:        schema.TypeString,
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

			"region": &schema.Schema{
				Optional:    true,
				Description: "Region of AMI",
				Type:        schema.TypeString,
			},

			"secret_key": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The secret key used to communicate with AWS. Learn how to set this.",
			},

			"chroot_mount": &schema.Schema{
				Optional:    true,
				Description: "This is a list of devices to mount into the chroot environment. This configuration parameter requires some additional documentation which is in the \"Chroot Mounts\" section below. Please read that section for more information on how to use this.",
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"values": &schema.Schema{
							Required:    true,
							Type:        schema.TypeList,
							Description: "Values of mount",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			"command_wrapper": &schema.Schema{
				Optional:    true,
				Description: "How to run shell commands. This defaults to {{.Command}}. This may be useful to set if you want to set environmental variables or perhaps run it with sudo or so on. This is a configuration template where the .Command variable is replaced with the command to be run. Defaults to \"{{.Command}}\".",
				Type:        schema.TypeString,
			},

			"copy_files": &schema.Schema{
				Optional:    true,
				Description: "Paths to files on the running EC2 instance that will be copied into the chroot environment prior to provisioning. Defaults to /etc/resolv.conf so that DNS lookups work. Pass an empty list to skip copying /etc/resolv.conf. You may need to do this if you're building an image that uses systemd",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"device_path": &schema.Schema{
				Optional:    true,
				Description: "The path to the device where the root volume of the source AMI will be attached. This defaults to \"\" (empty string), which forces Packer to find an open device automatically.",
				Type:        schema.TypeString,
			},

			"nvme_device_path": &schema.Schema{
				Optional:    true,
				Description: "When we call the mount command (by default mount -o device dir), the string provided in nvme_mount_path will replace device in that command. When this option is not set, device in that command will be something like /dev/sdf1, mirroring the attached device name. This assumption works for most instances but will fail with c5 and m5 instances. In order to use the chroot builder with c5 and m5 instances, you must manually set nvme_device_path and device_path.",
				Type:        schema.TypeString,
			},

			"from_scratch": &schema.Schema{
				Optional:    true,
				Description: "Build a new volume instead of starting from an existing AMI root volume snapshot. Default false. If true, source_ami is no longer used and the following options become required: ami_virtualization_type, pre_mount_commands and root_volume_size. The below options are also required in this mode only:",
				Type:        schema.TypeBool,
			},

			"mount_options": &schema.Schema{
				Optional:    true,
				Description: "Options to supply the mount command when mounting devices. Each option will be prefixed with -o and supplied to the mount command ran by Packer. Because this command is ran in a shell, user discretion is advised. See this manual page for the mount command for valid file system specific options",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"mount_partition": &schema.Schema{
				Optional:    true,
				Description: "The partition number containing the / partition. By default this is the first partition of the volume, (for example, xvda1) but you can designate the entire block device by setting \"mount_partition\": \"0\" in your config, which will mount xvda instead.",
				Type:        schema.TypeString,
			},

			"mount_path": &schema.Schema{
				Optional:    true,
				Description: "The path where the volume will be mounted. This is where the chroot environment will be. This defaults to /mnt/packer-amazon-chroot-volumes/{{.Device}}. This is a configuration template where the .Device variable is replaced with the name of the device where the volume is attached.",
				Type:        schema.TypeString,
			},

			"post_mount_commands": &schema.Schema{
				Optional:    true,
				Description: "As pre_mount_commands, but the commands are executed after mounting the root device and before the extra mount and copy steps. The device and mount path are provided by {{.Device}} and {{.MountPath}}.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"pre_mount_commands": &schema.Schema{
				Optional:    true,
				Description: "A series of commands to execute after attaching the root volume and before mounting the chroot. This is not required unless using from_scratch. If so, this should include any partitioning and filesystem creation commands. The path to the device is provided by {{.Device}}.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"root_device_name": &schema.Schema{
				Optional:    true,
				Description: "The root device name. For example, xvda.",
				Type:        schema.TypeString,
			},

			"root_volume_size": &schema.Schema{
				Optional:    true,
				Description: "The size of the root volume in GB for the chroot environment and the resulting AMI. Default size is the snapshot size of the source_ami unless from_scratch is true, in which case this field must be defined.",
				Type:        schema.TypeInt,
			},

			"source_ami": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The source AMI whose root volume will be copied and provisioned on the currently running instance. This must be an EBS-backed AMI with a root volume snapshot that you have access to. Note: this is not used when from_scratch is set to true.",
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
		},
	}
}
