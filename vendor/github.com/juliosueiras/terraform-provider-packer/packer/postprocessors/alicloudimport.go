package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func AlicloudImportResource() *schema.Resource {
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

			"access_key": &schema.Schema{
				Required:    true,
				Description: "This is the Alicloud access key. It must be provided, but it can also be sourced from the ALICLOUD_ACCESS_KEY environment variable.",
				Type:        schema.TypeString,
			},

			"secret_key": &schema.Schema{
				Required:    true,
				Description: "This is the Alicloud secret key. It must be provided, but it can also be sourced from the ALICLOUD_SECRET_KEY environment variable.",
				Type:        schema.TypeString,
			},

			"region": &schema.Schema{
				Required:    true,
				Description: "This is the Alicloud region. It must be provided, but it can also be sourced from the ALICLOUD_REGION environment variables.",
				Type:        schema.TypeString,
			},

			"image_name": &schema.Schema{
				Required:    true,
				Description: "The name of the user-defined image, [2, 128] English or Chinese characters. It must begin with an uppercase/lowercase letter or a Chinese character, and may contain numbers, _ or -. It cannot begin with http:// or https://.",
				Type:        schema.TypeString,
			},

			"image_description": &schema.Schema{
				Optional:    true,
				Description: "The description of the image, with a length limit of 0 to 256 characters. Leaving it blank means null, which is the default value. It cannot begin with http:// or https://.",
				Type:        schema.TypeString,
			},

			"image_force_delete": &schema.Schema{
				Optional:    true,
				Description: "If this value is true, when the target image name is duplicated with an existing image, it will delete the existing image and then create the target image, otherwise, the creation will fail. The default value is false.",
				Type:        schema.TypeBool,
			},

			"oss_bucket_name": &schema.Schema{
				Required:    true,
				Description: "The name of the OSS bucket where the RAW or VHD file will be copied to for import. If the Bucket isn't exist, post-process will create it for you.",
				Type:        schema.TypeString,
			},

			"oss_key_name": &schema.Schema{
				Optional:    true,
				Description: "The name of the object key in oss_bucket_name where the RAW or VHD file will be copied to for import.",
				Type:        schema.TypeString,
			},

			"skip_clean": &schema.Schema{
				Optional:    true,
				Description: "Whether we should skip removing the RAW or VHD file uploaded to OSS after the import process has completed. true means that we should leave it in the OSS bucket, false means to clean it out. Defaults to false.",
				Type:        schema.TypeBool,
			},

			"image_os_type": &schema.Schema{
				Required:    true,
				Description: "Type of the OS linux/windows",
				Type:        schema.TypeString,
			},

			"image_platform": &schema.Schema{
				Required:    true,
				Description: "platform such CentOS",
				Type:        schema.TypeString,
			},

			"image_architecture": &schema.Schema{
				Required:    true,
				Description: "Platform type of the image system:i386 | x86_64",
				Type:        schema.TypeString,
			},

			"image_system_size": &schema.Schema{
				Optional:    true,
				Description: "Size of the system disk, in GB, values range:",
				Type:        schema.TypeInt,
			},

			"format": &schema.Schema{
				Required:    true,
				Description: "The format of the image for import, now alicloud only support RAW and VHD.",
				Type:        schema.TypeString,
			},
		},
	}
}
