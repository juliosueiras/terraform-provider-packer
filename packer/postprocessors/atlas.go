package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func AtlasResource() *schema.Resource {
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
			"artifact": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The shorthand tag for your artifact that maps to Atlas, i.e hashicorp/foobar for atlas.hashicorp.com/hashicorp/foobar. You must have access to the organization—hashicorp in this example—in order to add an artifact to the organization in Atlas.",
			},
			"token": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Your access token for the Atlas API.",
			},

			"artifact_type": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "For uploading artifacts to Atlas. artifact_type can be set to any unique identifier, however, the following are recommended for consistency - amazon.image, azure.image, cloudstack.image, digitalocean.image, docker.image, googlecompute.image, hyperv.image, oneandone.image, openstack.image, parallels.image, profitbricks.image, qemu.image, triton.image, virtualbox.image, vmware.image, and custom.image.",
			},

			"atlas_url": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Override the base URL for Atlas. This is useful if you're using Atlas Enterprise in your own network. Defaults to https://atlas.hashicorp.com/api/v1.",
			},
			"metadata": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "Send metadata about the artifact.",
			},
		},
	}
}
