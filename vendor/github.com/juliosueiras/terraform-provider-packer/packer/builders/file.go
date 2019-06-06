package builders

import "github.com/hashicorp/terraform/helper/schema"

func FileResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},
			"source": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path for a file which will be copied as the artifact.",
			},

			"target": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The path for a file which will be copied as the artifact.",
			},

			"content": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The content that will be put into the artifact.",
			},
		},
	}
}
