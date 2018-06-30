package main

import (
	"fmt"
	ansibleprovisioner "github.com/hashicorp/packer/provisioner/ansible"
	ansiblelocalprovisioner "github.com/hashicorp/packer/provisioner/ansible-local"
	chefclientprovisioner "github.com/hashicorp/packer/provisioner/chef-client"
	chefsoloprovisioner "github.com/hashicorp/packer/provisioner/chef-solo"
	convergeprovisioner "github.com/hashicorp/packer/provisioner/converge"
	powershellprovisioner "github.com/hashicorp/packer/provisioner/powershell"
	puppetmasterlessprovisioner "github.com/hashicorp/packer/provisioner/puppet-masterless"
	puppetserverprovisioner "github.com/hashicorp/packer/provisioner/puppet-server"
	saltmasterlessprovisioner "github.com/hashicorp/packer/provisioner/salt-masterless"
	shellprovisioner "github.com/hashicorp/packer/provisioner/shell"
	shelllocalprovisioner "github.com/hashicorp/packer/provisioner/shell-local"
	windowsrestartprovisioner "github.com/hashicorp/packer/provisioner/windows-restart"
	windowsshellprovisioner "github.com/hashicorp/packer/provisioner/windows-shell"

	alicloudimportpostprocessor "github.com/hashicorp/packer/post-processor/alicloud-import"
	amazonimportpostprocessor "github.com/hashicorp/packer/post-processor/amazon-import"
	artificepostprocessor "github.com/hashicorp/packer/post-processor/artifice"
	atlaspostprocessor "github.com/hashicorp/packer/post-processor/atlas"
	checksumpostprocessor "github.com/hashicorp/packer/post-processor/checksum"
	compresspostprocessor "github.com/hashicorp/packer/post-processor/compress"
	dockerimportpostprocessor "github.com/hashicorp/packer/post-processor/docker-import"
	dockerpushpostprocessor "github.com/hashicorp/packer/post-processor/docker-push"
	dockersavepostprocessor "github.com/hashicorp/packer/post-processor/docker-save"
	dockertagpostprocessor "github.com/hashicorp/packer/post-processor/docker-tag"
	googlecomputeexportpostprocessor "github.com/hashicorp/packer/post-processor/googlecompute-export"
	googlecomputeimportpostprocessor "github.com/hashicorp/packer/post-processor/googlecompute-import"
	manifestpostprocessor "github.com/hashicorp/packer/post-processor/manifest"
	shelllocalpostprocessor "github.com/hashicorp/packer/post-processor/shell-local"
	vagrantpostprocessor "github.com/hashicorp/packer/post-processor/vagrant"
	vagrantcloudpostprocessor "github.com/hashicorp/packer/post-processor/vagrant-cloud"
	vspherepostprocessor "github.com/hashicorp/packer/post-processor/vsphere"
	vspheretemplatepostprocessor "github.com/hashicorp/packer/post-processor/vsphere-template"

	alicloudecsbuilder "github.com/hashicorp/packer/builder/alicloud/ecs"
	amazonchrootbuilder "github.com/hashicorp/packer/builder/amazon/chroot"
	amazonebsbuilder "github.com/hashicorp/packer/builder/amazon/ebs"
	amazonebssurrogatebuilder "github.com/hashicorp/packer/builder/amazon/ebssurrogate"
	amazonebsvolumebuilder "github.com/hashicorp/packer/builder/amazon/ebsvolume"
	amazoninstancebuilder "github.com/hashicorp/packer/builder/amazon/instance"
	azurearmbuilder "github.com/hashicorp/packer/builder/azure/arm"
	cloudstackbuilder "github.com/hashicorp/packer/builder/cloudstack"
	digitaloceanbuilder "github.com/hashicorp/packer/builder/digitalocean"
	dockerbuilder "github.com/hashicorp/packer/builder/docker"
	filebuilder "github.com/hashicorp/packer/builder/file"
	googlecomputebuilder "github.com/hashicorp/packer/builder/googlecompute"
	hypervisobuilder "github.com/hashicorp/packer/builder/hyperv/iso"
	hypervvmcxbuilder "github.com/hashicorp/packer/builder/hyperv/vmcx"
	lxcbuilder "github.com/hashicorp/packer/builder/lxc"
	lxdbuilder "github.com/hashicorp/packer/builder/lxd"
	ncloudbuilder "github.com/hashicorp/packer/builder/ncloud"
	nullbuilder "github.com/hashicorp/packer/builder/null"
	oneandonebuilder "github.com/hashicorp/packer/builder/oneandone"
	openstackbuilder "github.com/hashicorp/packer/builder/openstack"
	oracleclassicbuilder "github.com/hashicorp/packer/builder/oracle/classic"
	oracleocibuilder "github.com/hashicorp/packer/builder/oracle/oci"
	parallelsisobuilder "github.com/hashicorp/packer/builder/parallels/iso"
	parallelspvmbuilder "github.com/hashicorp/packer/builder/parallels/pvm"
	profitbricksbuilder "github.com/hashicorp/packer/builder/profitbricks"
	qemubuilder "github.com/hashicorp/packer/builder/qemu"
	scalewaybuilder "github.com/hashicorp/packer/builder/scaleway"
	tritonbuilder "github.com/hashicorp/packer/builder/triton"
	virtualboxisobuilder "github.com/hashicorp/packer/builder/virtualbox/iso"
	virtualboxovfbuilder "github.com/hashicorp/packer/builder/virtualbox/ovf"
	vmwareisobuilder "github.com/hashicorp/packer/builder/vmware/iso"
	vmwarevmxbuilder "github.com/hashicorp/packer/builder/vmware/vmx"
	"io/ioutil"
	"reflect"
)

func main() {

	// Provisioners
	generateParameters("provisioners", "AnsibleResource", &ansibleprovisioner.Config{}, "ansible")
	generateParameters("provisioners", "AnsibleLocalResource", &ansiblelocalprovisioner.Config{}, "ansiblelocal")
	generateParameters("provisioners", "ChefClientResource", &chefclientprovisioner.Config{}, "chefclient")
	generateParameters("provisioners", "ChefSoloResource", &chefsoloprovisioner.Config{}, "chefsolo")
	generateParameters("provisioners", "ConvergeResource", &convergeprovisioner.Config{}, "converge")
	//generateParameters("provisioners", "FileResource", &fileprovisioner.Config{}, "file")
	generateParameters("provisioners", "PowerShellResource", &powershellprovisioner.Config{}, "powershell")
	generateParameters("provisioners", "PuppetMasterlessResource", &puppetmasterlessprovisioner.Config{}, "puppetmasterless")
	generateParameters("provisioners", "PuppetServerResource", &puppetserverprovisioner.Config{}, "puppetserver")
	generateParameters("provisioners", "SaltMasterlessResource", &saltmasterlessprovisioner.Config{}, "saltmasterless")
	generateParameters("provisioners", "ShellResource", &shellprovisioner.Config{}, "shell")
	generateParameters("provisioners", "ShellLocalResource", &shelllocalprovisioner.Provisioner{}, "shelllocal")
	generateParameters("provisioners", "WindowsRestartResource", &windowsrestartprovisioner.Config{}, "windowsrestart")
	generateParameters("provisioners", "WindowsShellResource", &windowsshellprovisioner.Config{}, "windowsshell")

	// Post-Processor
	generateParameters("postprocessors", "AlicloudImportResource", &alicloudimportpostprocessor.Config{}, "alicloudimport")
	generateParameters("postprocessors", "AmazonImportResource", &amazonimportpostprocessor.Config{}, "amazonimport")
	generateParameters("postprocessors", "ArtificeResource", &artificepostprocessor.Config{}, "artifice")
	generateParameters("postprocessors", "AtlasResource", &atlaspostprocessor.Config{}, "atlas")
	generateParameters("postprocessors", "ChecksumResource", &checksumpostprocessor.Config{}, "checksum")
	generateParameters("postprocessors", "CompressResource", &compresspostprocessor.Config{}, "compress")
	generateParameters("postprocessors", "DockerImportResource", &dockerimportpostprocessor.Config{}, "dockerimport")
	generateParameters("postprocessors", "DockerPushResource", &dockerpushpostprocessor.Config{}, "dockerpush")
	generateParameters("postprocessors", "DockerSaveResource", &dockersavepostprocessor.Config{}, "dockersave")
	generateParameters("postprocessors", "DockerTagResource", &dockertagpostprocessor.Config{}, "dockertag")
	generateParameters("postprocessors", "GoogleComputeExportResource", &googlecomputeexportpostprocessor.Config{}, "googlecomputeexport")
	generateParameters("postprocessors", "GoogleComputeImportResource", &googlecomputeimportpostprocessor.Config{}, "googlecomputeimport")
	generateParameters("postprocessors", "ManifestResource", &manifestpostprocessor.Config{}, "manifest")
	generateParameters("postprocessors", "ShellLocalResource", &shelllocalpostprocessor.PostProcessor{}, "shelllocal")
	generateParameters("postprocessors", "VagrantResource", &vagrantpostprocessor.Config{}, "vagrant")
	generateParameters("postprocessors", "VagrantCloudResource", &vagrantcloudpostprocessor.Config{}, "vagrantcloud")
	generateParameters("postprocessors", "VSphereResource", &vspherepostprocessor.Config{}, "vsphere")
	generateParameters("postprocessors", "VSphereTemplateResource", &vspheretemplatepostprocessor.Config{}, "vspheretemplate")

	// Builders
	generateParameters("builders", "AlicloudECSResource", &alicloudecsbuilder.Config{}, "alicloudecs")
	generateParameters("builders", "AmazonChrootResource", &amazonchrootbuilder.Config{}, "amazonchroot")
	generateParameters("builders", "AmazonEBSResource", &amazonebsbuilder.Config{}, "amazonebs")
	generateParameters("builders", "AmazonEBSSurrogateResource", &amazonebssurrogatebuilder.Config{}, "amazonebssurrogate")
	generateParameters("builders", "AmazonEBSVolumeResource", &amazonebsvolumebuilder.Config{}, "amazonebsvolume")
	generateParameters("builders", "AmazonInstanceResource", &amazoninstancebuilder.Config{}, "amazoninstance")
	generateParameters("builders", "AzureARMResource", &azurearmbuilder.Config{}, "azurearm")
	generateParameters("builders", "CloudstackResource", &cloudstackbuilder.Config{}, "cloudstack")
	generateParameters("builders", "DigitalOceanResource", &digitaloceanbuilder.Config{}, "digitalocean")
	generateParameters("builders", "DockerResource", &dockerbuilder.Config{}, "docker")
	generateParameters("builders", "FileResource", &filebuilder.Config{}, "file")
	generateParameters("builders", "GoogleComputeResource", &googlecomputebuilder.Config{}, "googlecompute")
	generateParameters("builders", "HyperVISOResource", &hypervisobuilder.Config{}, "hyperviso")
	generateParameters("builders", "HyperVVMCXResource", &hypervvmcxbuilder.Config{}, "hypervvmcx")
	generateParameters("builders", "LXCResource", &lxcbuilder.Config{}, "lxc")
	generateParameters("builders", "LXDResource", &lxdbuilder.Config{}, "lxd")
	generateParameters("builders", "NCloudResource", &ncloudbuilder.Config{}, "ncloud")
	generateParameters("builders", "NullResource", &nullbuilder.Config{}, "null")
	generateParameters("builders", "OneandoneResource", &oneandonebuilder.Config{}, "oneandone")
	generateParameters("builders", "OpenstackResource", &openstackbuilder.Config{}, "openstack")
	generateParameters("builders", "OracleClassicResource", &oracleclassicbuilder.Config{}, "oracleclassic")
	generateParameters("builders", "OracleOciResource", &oracleocibuilder.Config{}, "oracleoci")
	generateParameters("builders", "ParallelsISOResource", &parallelsisobuilder.Config{}, "parallelsiso")
	generateParameters("builders", "ParallelsPVMResource", &parallelspvmbuilder.Config{}, "parallelspvm")
	generateParameters("builders", "ProfitbricksResource", &profitbricksbuilder.Config{}, "profitbricks")
	generateParameters("builders", "QEMUResource", &qemubuilder.Config{}, "qemu")
	generateParameters("builders", "ScalewayResource", &scalewaybuilder.Config{}, "scaleway")
	generateParameters("builders", "TritonResource", &tritonbuilder.Config{}, "triton")
	generateParameters("builders", "VirtualboxISOResource", &virtualboxisobuilder.Config{}, "virtualboxiso")
	generateParameters("builders", "VirtualboxOVFResource", &virtualboxovfbuilder.Config{}, "virtualboxovf")
	generateParameters("builders", "VMWareISOResource", &vmwareisobuilder.Config{}, "vmwareiso")
	generateParameters("builders", "VMWareVMXResource", &vmwarevmxbuilder.Config{}, "vmwarevmx")
}

func generateParameters(packerType string, name string, config interface{}, filename string) {
	res := ExtractParams(config)
	goCodeString := `
package ` + packerType + `

import "github.com/hashicorp/terraform/helper/schema"

func ` + name + `() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{`

	if packerType == "provisioners" {
		goCodeString += `
			"execute_order": &schema.Schema{
				Type: schema.TypeInt,
				Required: true,
			},`

	}

	result := ""
	for _, v := range res {
		name := v.Tag.Get("mapstructure")
		nameType := v.Type.String()
		//if nameType == "ptr" {
		//	nameType = v.Type.Elem().String()
		//}

		result += "\n\"" + name + "\": &schema.Schema{\n" +
			`Optional: true,`

		switch nameType {
		case "string", "time.Duration", "*string", "images.ImageVisibility":
			result += `
			Type: schema.TypeString,
`

		case "[]common.BlockDevice", "[]ebsvolume.BlockDevice", "[]converge.ModuleDir", "[]ecs.AlicloudDiskDevice":
			result += `
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
			},
`
		case "[]int", "[]int32", "[]int64", "[]uint":
			result += `
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeInt},
`
		case "[][]string":
			result += `
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
`

		case "slice", "[]string":
			result += `
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},
`
		case "map[string]string", "map[string]*string", "map", "common.TagMap":
			result += `
			Type: schema.TypeMap,
`
		case "bool":
			result += `
			Type: schema.TypeBool,
`
		case "int32", "int64", "int", "uint", "uint64", "number":
			result += `
			Type: schema.TypeInt,
`
		default:
			nameType = v.Type.Kind().String()
			fmt.Println(nameType)
			fmt.Println(v.Type.String())
			switch nameType {
			case "string", "time.Duration":
				result += `
			Type: schema.TypeString,
`

			case "[]common.BlockDevice":
				result += `
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
			},
`
			case "[]int":
				result += `
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeInt},
`
			case "[][]string":
				result += `
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
`

			case "slice", "[]string":
				result += `
			Type: schema.TypeList,
			Elem: &schema.Schema{Type: schema.TypeString},
`
			case "map[string]string", "map", "common.TagMap":
				result += `
			Type: schema.TypeMap,
`
			case "bool":
				result += `
			Type: schema.TypeBool,
`
			case "int32", "int64", "int", "uint", "uint64", "number":
				result += `
			Type: schema.TypeInt,
`
			}
		}

		result += `
		},
`
	}
	result += `
	},
	}
}

`
	goCodeString += result
	x := []byte(goCodeString)

	if err := ioutil.WriteFile("../packer/"+packerType+"/"+filename+".go", x, 0644); err != nil {
		panic(err)
	}

}

func extractParams(i reflect.Type, e []reflect.StructField) []reflect.StructField {
	for l := 0; l < i.NumField(); l++ {
		if i.Field(l).Tag != "" {
			if i.Field(l).Type.Kind().String() == "struct" {
				e = extractParams(i.Field(l).Type, e)
			} else {
				e = append(e, i.Field(l))
			}
		}
	}

	return e
}

func ExtractParams(p interface{}) []reflect.StructField {
	k := reflect.TypeOf(p).Elem()
	resultParams := []reflect.StructField{}
	return unique(extractParams(k, resultParams))
}

func unique(structSlice []reflect.StructField) []reflect.StructField {
	keys := make(map[string]bool)
	list := []reflect.StructField{}
	for _, entry := range structSlice {
		if _, value := keys[entry.Tag.Get("mapstructure")]; !value {
			keys[entry.Tag.Get("mapstructure")] = true
			list = append(list, entry)
		}
	}
	return list
}
