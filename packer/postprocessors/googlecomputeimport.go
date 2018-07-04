package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func GoogleComputeImportResource() *schema.Resource {
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

			"packer_build_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_builder_type": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_debug": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"packer_force": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},

			"packer_on_error": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"packer_user_variables": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},

			"bucket": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"gcs_object_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_description": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_family": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"image_labels": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},

			"image_name": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"project_id": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},

			"account_file": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}
