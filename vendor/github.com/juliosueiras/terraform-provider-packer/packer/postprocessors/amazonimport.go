package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func AmazonImportResource() *schema.Resource {
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
				Type:        schema.TypeString,
				Description: "The access key used to communicate with AWS",
			},

			"custom_endpoint_ec2": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This option is useful if you use a cloud provider whose API is compatible with aws EC2. Specify another endpoint like this https://ec2.custom.endpoint.com",
			},

			"mfa_code": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The MFA TOTP code. This should probably be a user variable since it changes all the time",
			},

			"profile": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The profile to use in the shared credentials file for AWS. See Amazon's documentation on specifying profiles for more details",
			},

			"region": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name of the region, such as us-east-1 in which to upload the OVA file to S3 and create the AMI. A list of valid regions can be obtained with AWS CLI tools or by consulting the AWS website",
			},

			"secret_key": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The secret key used to communicate with AWS",
			},

			"skip_region_validation": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Set to true if you want to skip validation of the region configuration option. Default false.",
			},

			"token": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The access token to use. This is different from the access key and secret key. If you're not sure what this is, then you probably don't need it. This will also be read from the AWS_SESSION_TOKEN environmental variable",
			},

			"s3_bucket_name": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name of the S3 bucket where the OVA file will be copied to for import. This bucket must exist when the post-processor is run",
			},

			"s3_key_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the key in s3_bucket_name where the OVA file will be copied to for import. If not specified, this will default to \"packer-import-{{timestamp}}.ova\". This key (ie, the uploaded OVA) will be removed after import, unless skip_clean is true.",
			},

			"skip_clean": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether we should skip removing the OVA file uploaded to S3 after the import process has completed. \"true\" means that we should leave it in the S3 bucket, \"false\" means to clean it out. Defaults to false",
			},

			"tags": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "Tags applied to the created AMI and relevant snapshots",
			},

			"ami_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the ami within the console. If not specified, this will default to something like ami-import-sfwerwf. Please note, specifying this option will result in a slightly longer execution time.",
			},

			"ami_description": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The description to set for the resulting imported AMI. By default this description is generated by the AMI import process.",
			},

			"ami_users": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "A list of account IDs that have access to launch the imported AMI. By default no additional users other than the user importing the AMI has permission to launch it",
			},

			"ami_groups": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "A list of groups that have access to launch the imported AMI. By default no groups have permission to launch the AMI. all will make the AMI publicly accessible. AWS currently doesn't accept any value other than \"all\".",
			},

			"license_type": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The license type to be used for the Amazon Machine Image (AMI) after importing. Valid values: AWS or BYOL (default). For more details regarding licensing, see Prerequisites in the VM Import/Export User Guide.",
			},

			"role_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the role to use when not using the default role, 'vmimport'",
			},
		},
	}
}