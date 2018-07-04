package postprocessors

import "github.com/hashicorp/terraform/helper/schema"

func ChecksumResource() *schema.Resource {
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
			"checksum_types": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of strings of checksum types to compute. Allowed values are md5, sha1, sha224, sha256, sha384, sha512.",
			},

			"output": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify filename to store checksums. This defaults to packer_{{.BuildName}}_{{.BuilderType}}_{{.ChecksumType}}.checksum. For example, if you had a builder named database, you might see the file written as packer_database_docker_md5.checksum. The following variables are available to use in the output template:",
			},
		},
	}
}
