package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func DockerPushResource() *schema.Resource {
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

			"login_username": &schema.Schema{
				Optional:    true,
				Description: "The username to use to authenticate to login.",
				Type:        schema.TypeString,
			},

			"login_password": &schema.Schema{
				Optional:    true,
				Description: "The password to use to authenticate to login",
				Type:        schema.TypeString,
			},

			"login_server": &schema.Schema{
				Optional:    true,
				Description: "The server address to login to.",
				Type:        schema.TypeString,
			},
			"login": &schema.Schema{
				Optional:    true,
				Description: "Defaults to false. If true, the post-processor will login prior to pushing. For log into ECR see ecr_login.",
				Type:        schema.TypeBool,
			},

			"ecr_login": &schema.Schema{
				Optional:    true,
				Description: "Defaults to false. If true, the post-processor will login in order to push the image to Amazon EC2 Container Registry (ECR). The post-processor only logs in for the duration of the push. If true login_server is required and login, login_username, and login_password will be ignored.",
				Type:        schema.TypeBool,
			},

			"aws_access_key": &schema.Schema{
				Optional:    true,
				Description: "The AWS access key used to communicate with AWS",
				Type:        schema.TypeString,
			},

			"aws_secret_key": &schema.Schema{
				Optional:    true,
				Description: "The AWS secret key used to communicate with AWS",
				Type:        schema.TypeString,
			},

			"aws_token": &schema.Schema{
				Optional:    true,
				Description: "The AWS access token to use. This is different from the access key and secret key. If you're not sure what this is, then you probably don't need it. This will also be read from the AWS_SESSION_TOKEN environmental variable.",
				Type:        schema.TypeString,
			},

			"aws_profile": &schema.Schema{
				Optional:    true,
				Description: "The AWS shared credentials profile used to communicate with AWS",
				Type:        schema.TypeString,
			},
		},
	}
}
