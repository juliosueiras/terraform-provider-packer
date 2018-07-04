package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func ShellLocalResource() *schema.Resource {
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
		},
	}
}
