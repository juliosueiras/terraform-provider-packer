package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func isoBuilderResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"iso_checksum": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The checksum for the OS ISO file. Because ISO files are so large, this is required and Packer will verify it prior to booting a virtual machine with the ISO attached. The type of the checksum is specified with iso_checksum_type, documented below. At least one of iso_checksum and iso_checksum_url must be defined. This has precedence over iso_checksum_url type.",
			},

			"iso_checksum_url": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "A URL to a GNU or BSD style checksum file containing a checksum for the OS ISO file. At least one of iso_checksum and iso_checksum_url must be defined. This will be ignored if iso_checksum is non empty.",
			},

			"iso_checksum_type": &schema.Schema{
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"none", "md5", "sha1", "sha256", "sha512"}, true),
				Description:  "The type of the checksum specified in iso_checksum. Valid values are none, md5, sha1, sha256, or sha512 currently. While none will skip checksumming, this is not recommended since ISO files are generally large and corruption does happen from time to time.",
			},

			"iso_urls": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Multiple URLs for the ISO to download. Packer will try these in order. If anything goes wrong attempting to download or while downloading a single URL, it will move on to the next. All URLs must point to the same file (same checksum). By default this is empty and iso_url is used. Only one of iso_url or iso_urls can be specified.",
			},

			"iso_target_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path where the iso should be saved after download. By default will go in the packer cache, with a hash of the original filename as its name.",
			},

			"iso_target_extension": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The extension of the iso file after download. This defaults to iso",
			},

			"iso_url": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "A URL to the ISO containing the installation image. This URL can be either an HTTP URL or a file URL (or path to a file). If this is an HTTP URL, Packer will download it and cache it between runs.",
			},
		},
	}
}
