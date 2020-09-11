package provider

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(time.Minute),
			Read:   schema.DefaultTimeout(time.Minute),
			Update: schema.DefaultTimeout(time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sample_attribute": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	f, err := os.Create("/tmp/" + d.Get("name").(string) + ".txt")
	if err != nil {
		return diag.FromErr(err)
	}
	defer f.Close()
	b, err := json.Marshal(d.Get("sample_attribute").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = f.Write([]byte(b))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("name", d.Get("name"))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("sample_attribute", d.Get("sample_attribute"))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(d.Get("name").(string))
	return nil
}

func resourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	b, err := ioutil.ReadFile("/tmp/" + d.Id() + ".txt")
	if err != nil {
		return diag.FromErr(err)
	}
	m := map[string]string{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("sample_attribute", m)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("name", d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	f, err := os.Create("/tmp/" + d.Id() + ".txt")
	if err != nil {
		return diag.FromErr(err)
	}
	defer f.Close()
	b, err := json.Marshal(d.Get("sample_attribute").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	_, err = f.Write([]byte(b))
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	err := os.Remove("/tmp/" + d.Id() + ".txt")
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
