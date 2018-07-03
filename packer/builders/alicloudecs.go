package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func AlicloudECSResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			"access_key": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "This is the Alicloud access key. It must be provided, but it can also be sourced from the ALICLOUD_ACCESS_KEY environment variable.",
			},

			"secret_key": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "This is the Alicloud secret key. It must be provided, but it can also be sourced from the ALICLOUD_SECRET_KEY environment variable.",
			},

			"region": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "Type of the instance. For values, see Instance Type Table. You can also obtain the latest instance type table by invoking the Querying Instance Type Table interface.",
			},

			"skip_region_validation": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "The region validation can be skipped if this value is true, the default value is false.",
			},

			"security_token": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "STS access token, can be set through template or by exporting as environment variable such \"export SecurityToken=value\"",
			},

			"image_name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name of the user-defined image, [2, 128] English or Chinese characters. It must begin with an uppercase/lowercase letter or a Chinese character, and may contain numbers, _ or -. It cannot begin with http:// or https://.",
			},

			"image_version": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The version number of the image, with a length limit of 1 to 40 English characters.",
			},

			"image_description": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The description of the image, with a length limit of 0 to 256 characters. Leaving it blank means null, which is the default value. It cannot begin with http:// or https://.",
			},

			"image_share_account": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The IDs of to-be-added Aliyun accounts to which the image is shared. The number of accounts is 1 to 10. If number of accounts is greater than 10, this parameter is ignored.",
			},

			"image_unshare_account": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			"image_copy_regions": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Copy to the destination regionIds.",
			},

			"image_copy_names": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The name of the destination image, [2, 128] English or Chinese characters. It must begin with an uppercase/lowercase letter or a Chinese character, and may contain numbers, _ or -. It cannot begin with http:// or https://.",
			},

			"image_force_delete": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If this value is true, when the target image name is duplicated with an existing image, it will delete the existing image and then create the target image, otherwise, the creation will fail. The default value is false",
			},

			"image_force_delete_snapshots": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If this value is true, when delete the duplicated existing image, the source snapshot of this image will be delete either.",
			},

			"image_force_delete_instances": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"image_disk_mapping": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"disk_category": &schema.Schema{
							Optional:     true,
							Type:         schema.TypeString,
							ValidateFunc: validation.StringInSlice([]string{"cloud", "cloud_efficiency", "cloud_ssd"}, true),
							Description:  "Category of the data disk.",
						},
						"disk_delete_with_instance": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeBool,
							Description: "Whether or not the disk is released along with the instance: True indicates that when the instance is released, this disk will be released with it, False indicates that when the instance is released, this disk will be retained.",
						},
						"disk_description": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeString,
							Description: "The value of disk description is blank by default. [2, 256] characters. The disk description will appear on the console. It cannot begin with http:// or https://.",
						},
						"disk_device": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeString,
							Description: "Device information of the related instance: such as /dev/xvdb It is null unless the Status is In_use.",
						},
						"disk_name": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeString,
							Description: "The value of disk name is blank by default. [2, 128] English or Chinese characters, must begin with an uppercase/lowercase letter or Chinese character. Can contain numbers, ., _ and -. The disk name will appear on the console. It cannot begin with http:// or https://.",
						},
						"disk_size": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeInt,
							Description: "Size of the system disk, in GB",
						},
						"disk_snapshot_id": &schema.Schema{
							Optional:    true,
							Type:        schema.TypeInt,
							Description: "Snapshots are used to create the data disk After this parameter is specified, Size is ignored. The actual size of the created disk is the size of the specified snapshot",
						},
					},
				},
			},
			"zone_id": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "ID of the zone to which the disk belongs.",
			},

			"io_optimized": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether an ECS instance is I/O optimized or not. The default value is false.",
			},
			"instance_type": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "Type of the instance. For values, see Instance Type Table. You can also obtain the latest instance type table by invoking the Querying Instance Type Table interface.",
			},

			"description": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"source_image": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "This is the base image id which you want to create your customized images.",
			},

			"force_stop_instance": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether to force shutdown upon device restart. The default value is false",
			},

			"security_group_id": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "ID of the security group to which a newly created instance belongs. Mutual access is allowed between instances in one security group. If not specified, the newly created instance will be added to the default security group. If the default group doesn’t exist, or the number of instances in it has reached the maximum limit, a new security group will be created automatically.",
			},

			"security_group_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The security group name. The default value is blank. [2, 128] English or Chinese characters, must begin with an uppercase/lowercase letter or Chinese character. Can contain numbers, ., _ or -. It cannot begin with http:// or https://.",
			},

			"user_data": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The UserData of an instance must be encoded in Base64 format, and the maximum size of the raw data is 16 KB.",
			},

			"user_data_file": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The file name of the userdata.",
			},

			"vpc_id": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "VPC ID allocated by the system.",
			},

			"vpc_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The VPC name. The default value is blank. [2, 128] English or Chinese characters, must begin with an uppercase/lowercase letter or Chinese character. Can contain numbers, _ and -. The disk description will appear on the console. Cannot begin with http:// or https://.",
			},

			"vpc_cidr_block": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"192.168.0.0/16", "172.16.0.0/16"}, true),
				Description:  "Value options: 192.168.0.0/16 and 172.16.0.0/16. When not specified, the default value is 172.16.0.0/16.",
			},

			"vswitch_id": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The ID of the VSwitch to be used.",
			},
			"instance_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Display name of the instance, which is a string of 2 to 128 Chinese or English characters. It must begin with an uppercase/lowercase letter or a Chinese character and can contain numerals, ., _, or -. The instance name is displayed on the Alibaba Cloud console. If this parameter is not specified, the default value is InstanceId of the instance. It cannot begin with http:// or https://.",
			},

			"internet_charge_type": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"PayByBandwidth", "PayByTraffic"}, true),
				Description:  "Internet charge type, which can be PayByTraffic or PayByBandwidth",
			},

			"internet_max_bandwidth_out": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second).",
			},

			"temporary_key_pair_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the temporary key pair to generate. By default, Packer generates a name that looks like packer_<UUID>, where <UUID> is a 36 character unique identifier",
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
