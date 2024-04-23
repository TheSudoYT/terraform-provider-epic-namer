package epic

import (
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRandomName() *schema.Resource {
	return &schema.Resource{
		Create: resourceRandomNameCreate,
		Read:   schema.Noop,
		Update: schema.Noop,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceRandomNameCreate(d *schema.ResourceData, m interface{}) error {
	rand.NewSource(time.Now().UnixNano())
	names := []string{"Aragorn", "Gandalf", "Bilbo", "Thorin", "Legolas"}
	selectedName := names[rand.Intn(len(names))]

	d.SetId(time.Now().UTC().String())
	d.Set("name", selectedName)

	return nil
}
