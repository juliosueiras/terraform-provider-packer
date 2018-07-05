package builders

import "github.com/hashicorp/terraform/helper/schema"

func DockerResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "for named builds",
			},
			"communicator": &schema.Schema{
				Optional: true,
				Type:     schema.TypeString,
			},
			"commit": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "If true, the container will be committed to an image rather than exported.",
			},
			"discard": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Throw away the container when the build is complete. This is useful for the artifice post-processor.",
			},
			"image": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The base image for the Docker container that will be started. This image will be pulled from the Docker registry if it doesn't already exist.",
			},
			"author": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Set the author (e-mail) of a commit.",
			},

			"export_path": &schema.Schema{
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The path where the final container will be exported as a tar file",
			},
			"container_dir": &schema.Schema{
				Optional:    true,
				Description: "The directory inside container to mount temp directory from host server for work file provisioner. By default this is set to /packer-files.",
				Type:        schema.TypeString,
			},

			"exec_user": &schema.Schema{
				Optional:    true,
				Description: "Username or UID (format: [:]) to run remote commands with. You may need this if you get permission errors trying to run the shell or other provisioners.",
				Type:        schema.TypeString,
			},

			"privileged": &schema.Schema{
				Optional:    true,
				Description: "If true, run the docker container with the --privileged flag. This defaults to false if not set.",
				Type:        schema.TypeBool,
			},
			"volumes": &schema.Schema{
				Optional:    true,
				Description: "A mapping of additional volumes to mount into this container. The key of the object is the host path, the value is the container path.",
				Type:        schema.TypeMap,
			},

			"run_command": &schema.Schema{
				Optional:    true,
				Description: "An array of arguments to pass to docker run in order to run the container. By default this is set to [\"-d\", \"-i\", \"-t\", \"{{.Image}}\", \"/bin/bash\"]. As you can see, you have a couple template variables to customize, as well.",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"changes": &schema.Schema{
				Optional:    true,
				Description: "Dockerfile instructions to add to the commit. Example of instructions are CMD, ENTRYPOINT, ENV, and EXPOSE. Example: [ \"USER ubuntu\", \"WORKDIR /app\", \"EXPOSE 8080\" ]",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"fix_upload_owner": &schema.Schema{
				Optional:    true,
				Description: "If true, files uploaded to the container will be owned by the user the container is running as. If false, the owner will depend on the version of docker installed in the system. Defaults to true.",
				Type:        schema.TypeBool,
			},
			"login": &schema.Schema{
				Optional:    true,
				Description: "Defaults to false. If true, the builder will login in order to pull the image. The builder only logs in for the duration of the pull. It always logs out afterwards. For log into ECR see ecr_login.",
				Type:        schema.TypeBool,
			},

			"login_password": &schema.Schema{
				Optional:    true,
				Description: "The password to use to authenticate to login.",
				Type:        schema.TypeString,
			},
			"message": &schema.Schema{
				Optional:    true,
				Description: "Set a message for the commit.",
				Type:        schema.TypeString,
			},

			"login_server": &schema.Schema{
				Optional:    true,
				Description: "The server address to login to.",
				Type:        schema.TypeString,
			},

			"login_username": &schema.Schema{
				Optional:    true,
				Description: "The username to use to authenticate to login.",
				Type:        schema.TypeString,
			},

			"ecr_login": &schema.Schema{
				Optional:    true,
				Description: "Defaults to false. If true, the builder will login in order to pull the image from Amazon EC2 Container Registry (ECR). The builder only logs in for the duration of the pull. If true login_server is required and login, login_username, and login_password will be ignored. For more information see the section on ECR.",
				Type:        schema.TypeBool,
			},
			"pull": &schema.Schema{
				Optional:    true,
				Description: "If true, the configured image will be pulled using docker pull prior to use. Otherwise, it is assumed the image already exists and can be used. This defaults to true if not set.",
				Type:        schema.TypeBool,
			},

			"aws_access_key": &schema.Schema{
				Optional:    true,
				Description: "The AWS access key used to communicate with AWS. Learn how to set this.",
				Type:        schema.TypeString,
			},

			"aws_secret_key": &schema.Schema{
				Optional:    true,
				Description: "The AWS secret key used to communicate with AWS. Learn how to set this.",
				Type:        schema.TypeString,
			},

			"aws_token": &schema.Schema{
				Optional:    true,
				Description: "The AWS access token to use. This is different from the access key and secret key. If you're not sure what this is, then you probably don't need it. This will also be read from the AWS_SESSION_TOKEN environmental variable.",
				Type:        schema.TypeString,
			},

			"aws_profile": &schema.Schema{
				Optional:    true,
				Description: "The AWS shared credentials profile used to communicate with AWS. Learn how to set this.",
				Type:        schema.TypeString,
			},
		},
	}
}
