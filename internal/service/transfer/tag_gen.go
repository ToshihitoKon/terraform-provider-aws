// Code generated by internal/generate/tagresource/main.go; DO NOT EDIT.

package transfer

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/service/transfer"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func ResourceTag() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceTagCreate,
		ReadWithoutTimeout:   resourceTagRead,
		UpdateWithoutTimeout: resourceTagUpdate,
		DeleteWithoutTimeout: resourceTagDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"resource_arn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).TransferConn

	identifier := d.Get("resource_arn").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)

	if err := UpdateTagsWithContext(ctx, conn, identifier, nil, map[string]string{key: value}); err != nil {
		return diag.Errorf("creating %s resource (%s) tag (%s): %s", transfer.ServiceID, identifier, key, err)
	}

	d.SetId(tftags.SetResourceID(identifier, key))

	return resourceTagRead(ctx, d, meta)
}

func resourceTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).TransferConn
	identifier, key, err := tftags.GetResourceID(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	value, err := GetTagWithContext(ctx, conn, identifier, key)

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] %s resource (%s) tag (%s) not found, removing from state", transfer.ServiceID, identifier, key)
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.Errorf("reading %s resource (%s) tag (%s): %s", transfer.ServiceID, identifier, key, err)
	}

	d.Set("resource_arn", identifier)
	d.Set("key", key)
	d.Set("value", value)

	return nil
}

func resourceTagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).TransferConn
	identifier, key, err := tftags.GetResourceID(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	if err := UpdateTagsWithContext(ctx, conn, identifier, nil, map[string]string{key: d.Get("value").(string)}); err != nil {
		return diag.Errorf("updating %s resource (%s) tag (%s): %s", transfer.ServiceID, identifier, key, err)
	}

	return resourceTagRead(ctx, d, meta)
}

func resourceTagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).TransferConn
	identifier, key, err := tftags.GetResourceID(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	if err := UpdateTagsWithContext(ctx, conn, identifier, map[string]string{key: d.Get("value").(string)}, nil); err != nil {
		return diag.Errorf("deleting %s resource (%s) tag (%s): %s", transfer.ServiceID, identifier, key, err)
	}

	return nil
}
