package packer

import (
	"encoding/json"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/juliosueiras/terraform-provider-packer/packer/builders"
	"github.com/juliosueiras/terraform-provider-packer/packer/postprocessors"
	"github.com/juliosueiras/terraform-provider-packer/packer/provisioners"
	"log"
	"strconv"
	"strings"
)

func dataSourceJSON() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJSONRead,

		Schema: map[string]*schema.Schema{
			"builders": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     builderResource(),
			},
			"provisioners": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     provisionerResource(),
			},
			"post_processors": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
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
			"vmware_iso": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     builders.VMwareISOBuilderResource(),
			},
		},
	}
}

func provisionerResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"windows_restart": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.WindowsRestartResource(),
			},
			"file": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.FileResource(),
			},
			"windows_shell": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.WindowsShellResource(),
			},
			"powershell": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     provisioners.PowerShellResource(),
			},
		},
	}
}

func postProcessorResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"vsphere": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem:     postprocessors.VSphereResource(),
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
			temp := e.([]interface{})[0].(map[string]interface{})

			for k, v := range temp {
				t[k] = v
			}

			for f := range t {
				if _, errChange := d.GetOkExists("builders." + strconv.FormatInt(int64(x), 10) + "." + i + ".0." + f); !errChange {
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
			} else {
				totalOrder++
			}
		}

		provisioners = make([]map[string]interface{}, totalOrder)

		for i, e := range v.(map[string]interface{}) {
			log.Println(e.([]interface{}))

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

	if provisioners[len(provisioners)-1] == nil {
		provisioners = provisioners[:len(provisioners)-1]
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
