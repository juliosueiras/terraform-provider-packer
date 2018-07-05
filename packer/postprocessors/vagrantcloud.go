package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VagrantCloudResource() *schema.Resource {
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

			"box_tag": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The shorthand tag for your box that maps to Vagrant Cloud, i.e hashicorp/precise64 for vagrantcloud.com/hashicorp/precise64",
			},

			"version": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The version number, typically incrementing a previous version. The version string is validated based on Semantic Versioning. The string must match a pattern that could be semver, and doesn't validate that the version comes after your previous versions.",
			},

			"version_description": &schema.Schema{
				Optional:    true,
				Description: "Optionally markdown text used as a full-length and in-depth description of the version, typically for denoting changes introduced",
				Type:        schema.TypeString,
			},

			"no_release": &schema.Schema{
				Optional:    true,
				Description: "If set to true, does not release the version on Vagrant Cloud, making it active. You can manually release the version via the API or Web UI. Defaults to false.",
				Type:        schema.TypeString,
			},

			"access_token": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Your access token for the Vagrant Cloud API. This can be generated on your tokens page. If not specified, the environment will be searched. First, VAGRANT_CLOUD_TOKEN is checked, and if nothing is found, finally ATLAS_TOKEN will be used.",
			},

			"vagrant_cloud_url": &schema.Schema{
				Optional:    true,
				Description: "Override the base URL for Vagrant Cloud. This is useful if you're using Vagrant Private Cloud in your own network. Defaults to https://vagrantcloud.com/api/v1",
				Type:        schema.TypeString,
			},

			"box_download_url": &schema.Schema{
				Optional:    true,
				Description: "Optional URL for a self-hosted box. If this is set the box will not be uploaded to the Vagrant Cloud.",
				Type:        schema.TypeString,
			},
		},
	}
}
