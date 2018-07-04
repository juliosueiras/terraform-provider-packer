package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func VagrantResource() *schema.Resource {
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

			"compression_level": &schema.Schema{
				Optional:    true,
				Description: "An integer representing the compression level to use when creating the Vagrant box. Valid values range from 0 to 9, with 0 being no compression and 9 being the best compression. By default, compression is enabled at level 6.",
				Type:        schema.TypeInt,
			},

			"include": &schema.Schema{
				Optional:    true,
				Description: "Paths to files to include in the Vagrant box. These files will each be copied into the top level directory of the Vagrant box (regardless of their paths). They can then be used from the Vagrantfile.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"output": &schema.Schema{
				Optional:    true,
				Description: "The full path to the box file that will be created by this post-processor. This is a configuration template. The variable Provider is replaced by the Vagrant provider the box is for. The variable ArtifactId is replaced by the ID of the input artifact. The variable BuildName is replaced with the name of the build. By default, the value of this config is packer_{{.BuildName}}_{{.Provider}}.box.",
				Type:        schema.TypeString,
			},

			"vagrantfile_template": &schema.Schema{
				Optional:    true,
				Description: "Path to a template to use for the Vagrantfile that is packaged with the box.",
				Type:        schema.TypeString,
			},

			"override": &schema.Schema{
				Optional: true,
				Type:     schema.TypeMap,
			},
		},
	}
}

func vagrantOverrideResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compression_level": &schema.Schema{
				Optional:    true,
				Description: "An integer representing the compression level to use when creating the Vagrant box. Valid values range from 0 to 9, with 0 being no compression and 9 being the best compression. By default, compression is enabled at level 6.",
				Type:        schema.TypeInt,
			},

			"include": &schema.Schema{
				Optional:    true,
				Description: "Paths to files to include in the Vagrant box. These files will each be copied into the top level directory of the Vagrant box (regardless of their paths). They can then be used from the Vagrantfile.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"output": &schema.Schema{
				Optional:    true,
				Description: "The full path to the box file that will be created by this post-processor. This is a configuration template. The variable Provider is replaced by the Vagrant provider the box is for. The variable ArtifactId is replaced by the ID of the input artifact. The variable BuildName is replaced with the name of the build. By default, the value of this config is packer_{{.BuildName}}_{{.Provider}}.box.",
				Type:        schema.TypeString,
			},

			"vagrantfile_template": &schema.Schema{
				Optional:    true,
				Description: "Path to a template to use for the Vagrantfile that is packaged with the box.",
				Type:        schema.TypeString,
			},
			"keep_input_artifact": &schema.Schema{
				Optional: true,
				Type:     schema.TypeBool,
			},
		},
	}
}
