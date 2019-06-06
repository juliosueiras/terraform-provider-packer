package builders

import "github.com/hashicorp/terraform/helper/schema"

func LXDResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},

			"output_image": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the output artifact. Defaults to name.",
			},

			"command_wrapper": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Lets you prefix all builder commands, such as with ssh for a remote build host. Defaults to \"\".",
			},

			"image": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "The source image to use when creating the build container. This can be a (local or remote) image (name or fingerprint). E.G. my-base-image, ubuntu-daily:x, 08fababf6f27, ...",
			},

			"init_sleep": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The number of seconds to sleep between launching the LXD instance and provisioning it; defaults to 3 seconds.",
			},

			"publish_properties": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeMap,
				Description: "Pass key values to the publish step to be set as properties on the output image. This is most helpful to set the description, but can be used to set anything needed. See https://stgraber.org/2016/03/30/lxd-2-0-image-management-512/ for more properties.",
			},
		},
	}
}
