package packer

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/builders"
	"github.com/juliosueiras/terraform-provider-packer/packer/postprocessors"
	"github.com/juliosueiras/terraform-provider-packer/packer/provisioners"
	"log"
	"strconv"
	"strings"
)

var BlockNames = map[string]bool{
	"ebs_volume":         true,
	"image_disk_mapping": true,
}

var SpecialNames = map[string]bool{
	"vboxmanage":      true,
	"vboxmanage_post": true,
}

func dataSourceJSON() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJSONRead,

		Schema: map[string]*schema.Schema{
			"builders": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     builderResource(),
			},
			"provisioners": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisionerResource(),
			},
			"post_processors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     postProcessorResource(),
			},
			"variables": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"json": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func builderResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alicloud_ecs": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.AlicloudECSResource(),
			},
			"amazon_chroot": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.AmazonChrootResource(),
			},
			"amazon_ebs": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.AmazonEBSResource(),
			},
			"amazon_ebssurrogate": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.AmazonEBSSurrogateResource(),
			},
			"amazon_ebsvolume": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.AmazonEBSVolumeResource(),
			},
			"amazon_instance": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.AmazonInstanceResource(),
			},
			"azurearm": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.AzureARMResource(),
			},
			"cloudstack": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.CloudstackResource(),
			},
			"digitalocean": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.DigitalOceanResource(),
			},
			"docker": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.DockerResource(),
			},
			"file": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.FileResource(),
			},
			"googlecompute": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.GoogleComputeResource(),
			},
			"hyperv_iso": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.HyperVISOResource(),
			},
			"hyperv_vmcx": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.HyperVVMCXResource(),
			},
			"lxc": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.LXCResource(),
			},
			"lxd": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.LXDResource(),
			},
			"ncloud": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.NCloudResource(),
			},
			"null": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.NullResource(),
			},
			"oneandone": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.OneandoneResource(),
			},
			"oracle_classic": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.OracleClassicResource(),
			},
			"oracle_oci": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.OracleOciResource(),
			},
			"parallels_iso": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.ParallelsISOResource(),
			},
			"parallels_pvm": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.ParallelsPVMResource(),
			},
			"profitbricks": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.ProfitbricksResource(),
			},
			"qemu": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.QEMUResource(),
			},
			"scaleway": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.ScalewayResource(),
			},
			"triton": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.TritonResource(),
			},
			"virtualbox_iso": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.VirtualboxISOResource(),
			},
			"virtualbox_ovf": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.VirtualboxOVFResource(),
			},
			"vmware_iso": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.VMWareISOResource(),
			},
			"vmware_vmx": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.VMWareVMXResource(),
			},
		},
	}
}

func provisionerResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ansible": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.AnsibleResource(),
			},
			"ansible_local": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.AnsibleLocalResource(),
			},
			"chef_client": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.ChefClientResource(),
			},
			"chef_solo": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.ChefSoloResource(),
			},
			"converge": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.ConvergeResource(),
			},
			"file": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.FileResource(),
			},
			"powershell": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.PowerShellResource(),
			},
			"puppet_masterless": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.PuppetMasterlessResource(),
			},
			"puppet_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.PuppetServerResource(),
			},
			"salt_masterless": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.SaltMasterlessResource(),
			},
			"shell": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.ShellResource(),
			},
			"shell_local": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.ShellLocalResource(),
			},
			"windows_restart": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.WindowsRestartResource(),
			},
			"windows_shell": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.WindowsShellResource(),
			},
		},
	}
}

func postProcessorResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"alicloud_import": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.AlicloudImportResource(),
			},
			"amazon_import": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.AmazonImportResource(),
			},
			"artifice": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.ArtificeResource(),
			},
			"atlas": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.AtlasResource(),
			},
			"checksum": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.ChecksumResource(),
			},
			"compress": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.CompressResource(),
			},
			"docker_import": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.DockerImportResource(),
			},
			"docker_push": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.DockerPushResource(),
			},
			"docker_save": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.DockerSaveResource(),
			},
			"docker_tag": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.DockerTagResource(),
			},
			"googlecompute_export": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.GoogleComputeExportResource(),
			},
			"googlecompute_import": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.GoogleComputeImportResource(),
			},
			"manifest": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.ManifestResource(),
			},
			"shell_local": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.ShellLocalResource(),
			},
			"vagrant": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.VagrantResource(),
			},
			"vagrant_cloud": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.VagrantCloudResource(),
			},
			"vsphere": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.VSphereResource(),
			},
			"vsphere_template": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.VSphereTemplateResource(),
			},
		},
	}
}

func dataSourceJSONRead(d *schema.ResourceData, meta interface{}) error {

	builders := make([]map[string]interface{}, 0)
	provisioners := make([]map[string]interface{}, 0)
	postprocessors := make([]map[string]interface{}, 0)
	for x, v := range d.Get("builders").([]interface{}) {
		log.Print(x, v, builders)
		if v == nil {
			break
		}
		for i, e := range v.(map[string]interface{}) {
			t := make(map[string]interface{})

			//if _, errChange := d.GetOkExists("builders." + strconv.FormatInt(int64(x), 10) + "." + i); errChange && (len(e.([]interface{})) == 0) {
			if len(e.([]interface{})) == 0 {
				//	t["type"] = strings.Replace(i, "_", "-", -1)
				//	builders = append(builders, t)
				continue
			}

			temp := e.([]interface{})[0].(map[string]interface{})

			for k, v := range temp {
				t[k] = v
			}

			for f := range t {
				if BlockNames[f] {
					tempBlock := t[f]
					t[f+"s"] = removeUnusedParams(d, tempBlock.([]interface{}), x, i, f[:len(f)])
					delete(t, f)
				} else if SpecialNames[f] && len(t[f].([]interface{})) != 0 {
					resultSpecial := make([][]string, len(t[f].([]interface{})))
					for w, b := range t[f].([]interface{}) {
						resultSpecial[w] = append(resultSpecial[w], b.(map[string]interface{})["command"].(string))

						for _, value := range b.(map[string]interface{})["values"].([]interface{}) {
							resultSpecial[w] = append(resultSpecial[w], value.(string))
						}
					}

					t[f] = resultSpecial
				} else if _, errChange := d.GetOkExists("builders." + strconv.FormatInt(int64(x), 10) + "." + i + ".0." + f); !errChange && !BlockNames[f[:len(f)-2]] {
					delete(t, f)
				}
			}

			t["type"] = strings.Replace(i, "_", "-", -1)

			for _, e := range []string{"iso", "ssh", "winrm", "remote"} {
				if t[e] == nil {
					continue
				}

				mapType := t[e].([]interface{})
				if len(mapType) != 0 && mapType[0] != nil {
					for k, v := range mapType[0].(map[string]interface{}) {
						if _, change := d.GetOkExists("builders." + strconv.FormatInt(int64(x), 10) + "." + i + ".0." + e + ".0." + k); change {
							t[k] = v
						}
					}
				}
				delete(t, e)
			}

			builders = append(builders, t)
		}
	}

	for _, v := range d.Get("provisioners").([]interface{}) {
		if v == nil {
			break
		}

		totalOrder := 0
		for _, e := range v.(map[string]interface{}) {
			if len(e.([]interface{})) > 1 {
				totalOrder += len(e.([]interface{}))
			} else if len(e.([]interface{})) == 1 {
				totalOrder++
			}
		}

		provisioners = make([]map[string]interface{}, totalOrder)

		for i, e := range v.(map[string]interface{}) {
			log.Println(e.([]interface{}))
			if len(e.([]interface{})) == 0 {
				//	t["type"] = strings.Replace(i, "_", "-", -1)
				//	builders = append(builders, t)
				continue
			}

			for m, l := range e.([]interface{}) {
				t := make(map[string]interface{})

				for k, v := range l.(map[string]interface{}) {
					t[k] = v
				}

				for f := range t {
					log.Println("provisioners.0." + i + "." + strconv.FormatInt(int64(m), 10) + "." + f)
					if _, errChange := d.GetOkExists("provisioners.0." + i + "." + strconv.FormatInt(int64(m), 10) + "." + f); !errChange {
						delete(t, f)
					}
				}

				t["type"] = strings.Replace(i, "_", "-", -1)

				executeOrder := t["execute_order"].(int)

				if (executeOrder-1) < 0 || executeOrder > len(provisioners) {

					return errors.New(strconv.FormatInt(int64(t["execute_order"].(int)), 10) + " execute order is incorrect")
				}

				delete(t, "execute_order")
				provisioners[executeOrder-1] = t
			}
		}
	}

	for x, v := range d.Get("post_processors").([]interface{}) {
		log.Print(x, v, postprocessors)
		if v == nil {
			break
		}
		for i, e := range v.(map[string]interface{}) {
			t := make(map[string]interface{})

			//if _, errChange := d.GetOkExists("post_processors." + strconv.FormatInt(int64(x), 10) + "." + i); errChange && len(e.([]interface{})) == 0 {
			//	t["type"] = strings.Replace(i, "_", "-", -1)
			//	postprocessors = append(postprocessors, t)
			//	continue
			//}
			if len(e.([]interface{})) == 0 {
				continue
			}

			temp := e.([]interface{})[0].(map[string]interface{})

			for k, v := range temp {
				t[k] = v
			}

			for f := range t {
				if _, errChange := d.GetOkExists("post_processors." + strconv.FormatInt(int64(x), 10) + "." + i + ".0." + f); !errChange {
					delete(t, f)
				}
			}

			t["type"] = strings.Replace(i, "_", "-", -1)

			for _, e := range []string{"iso", "ssh", "winrm", "remote"} {
				if t[e] == nil {
					continue
				}

				mapType := t[e].([]interface{})
				if len(mapType) != 0 && mapType[0] != nil {
					for k, v := range mapType[0].(map[string]interface{}) {
						if _, change := d.GetOkExists("post_processors." + strconv.FormatInt(int64(x), 10) + "." + i + ".0." + e + ".0." + k); change {
							t[k] = v
						}
					}
				}
				delete(t, e)
			}

			postprocessors = append(postprocessors, t)
		}
	}

	res := map[string]interface{}{
		"builders":        builders,
		"provisioners":    provisioners,
		"post-processors": postprocessors,
	}

	if _, change := d.GetOkExists("variables"); change {
		res["variables"] = d.Get("variables")
	}

	result, err := json.MarshalIndent(res, "", "  ")

	if err != nil {
		return err
	}

	d.Set("json", string(result))
	d.SetId(strconv.Itoa(hashcode.String(string(result))))

	return nil
}

func removeUnusedParams(d *schema.ResourceData, block []interface{}, num int, builder string, blockName string) []interface{} {
	for i, m := range block {
		if m != nil {
			for k := range m.(map[string]interface{}) {

				if _, change := d.GetOkExists("builders." + strconv.FormatInt(int64(num), 10) + "." + builder + ".0." + blockName + "." + strconv.FormatInt(int64(i), 10) + "." + k); !change {
					delete(m.(map[string]interface{}), k)
				}
			}
		} else {
			block = append(block[:i], block[i+1:]...)
		}
	}

	return block
}
