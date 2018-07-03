package provisioners

import "github.com/hashicorp/terraform/helper/schema"

func FileResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"execute_order": &schema.Schema{
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The order for this provisioner to run in",
			},
			"source": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The path to a local file or directory to upload to the machine. The path can be absolute or relative. If it is relative, it is relative to the working directory when Packer is executed. If this is a directory, the existence of a trailing slash is important. Read below on uploading directories.",
			},
			"destination": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The path where the file will be uploaded to in the machine. This value must be a writable location and any parent directories must already exist.",
			},
			"direction": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The direction of the file transfer. This defaults to \"upload\". If it is set to \"download\" then the file \"source\" in the machine will be downloaded locally to \"destination\"",
			},
			"generated": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "For advanced users only. If true, check the file existence only before uploading, rather than upon pre-build validation. This allows to upload files created on-the-fly. This defaults to false. We don't recommend using this feature, since it can cause Packer to become dependent on system state. We would prefer you generate your files before the Packer run, but realize that there are situations where this may be unavoidable.",
			},
		},
	}
}
