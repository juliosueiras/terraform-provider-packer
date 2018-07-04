package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func CompressResource() *schema.Resource {
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

			"output": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path to save the compressed archive. The archive format is inferred from the filename. E.g. .tar.gz will be a gzipped tarball. .zip will be a zip file. If the extension can't be detected packer defaults to .tar.gz behavior but will not change the filename.  You can use {{.BuildName}} and {{.BuilderType}} in your output path. If you are executing multiple builders in parallel you should make sure output is unique for each one. For example packer_{{.BuildName}}.zip",
			},

			"format": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Disable archive format autodetection and use provided string.",
			},

			"compression_level": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Specify the compression level, for algorithms that support it, from 1 through 9 inclusive. Typically higher compression levels take longer but produce smaller files. Defaults to 6",
			},
		},
	}
}
