package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func DockerTagResource() *schema.Resource {
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
			"repository": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The repository of the image",
			},

			"tag": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The tag for the image. By default this is not set",
			},
			"force": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, this post-processor forcibly tag the image even if tag name is collided. Default to false. But it will be ignored if Docker >= 1.12.0 was detected, since the force option was removed after 1.12.0. ",
			},
		},
	}
}
