package builders

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/communicators"
)

func NCloudResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},

			"ncloud_access_key": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "User's access key. Go to [Account Management > Authentication Key] to create and view your authentication key.",
			},

			"ncloud_secret_key": &schema.Schema{
				Required:    true,
				Type:        schema.TypeString,
				Description: "User's secret key paired with the access key. Go to [Account Management > Authentication Key] to create and view your authentication key.",
			},

			"server_image_product_code": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Product code of an image to create. (member_server_image_no is required if not specified)",
			},

			"server_product_code": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Product (spec) code to create",
			},

			"member_server_image_no": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Previous image code. If there is an image previously created, it can be used to create a new image. (server_image_product_code is required if not specified)",
			},

			"server_image_name": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Name of an image to create.",
			},

			"server_image_description": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Description of an image to create",
			},

			"user_data": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
				Description: `Init script to run when an instance is created.

				For Linux servers, Python, Perl, and Shell scripts can be used. The path of the script to run should be included at the beginning of the script, like #!/usr/bin/env python, #!/bin/perl, or #!/bin/bash.
				For Windows servers, only Visual Basic scripts can be used.
				All scripts must be written in English.`,
			},

			"user_data_file": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "A path to a file containing a user_data script. See above for more information.",
			},

			"block_storage_size": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "You can add block storage ranging from 10 GB to 2000 GB, in increments of 10 GB",
			},

			"region": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
				Description: `Name of the region where you want to create an image. (default: Korea)

				values: Korea / US-West / HongKong / Singapore / Japan / Germany`,
			},

			"access_control_group_configuration_no": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "This is used to allow winrm access when you create a Windows server. An ACG that specifies an access source (0.0.0.0/0) and allowed port (5985) must be created in advance.",
			},

			"communicator": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"ssh": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     communicators.SSHCommunicatorResource(),
			},
			"winrm": &schema.Schema{
				Optional: true,
				Type:     schema.TypeList,
				Elem:     communicators.WinRMCommunicatorResource(),
			},
		},
	}
}
