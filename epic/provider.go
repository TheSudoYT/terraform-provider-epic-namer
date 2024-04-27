package epic

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"epic_random_name":  resourceRandomName(),
			"epic_random_quote": resourceRandomQuote(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		DefaultMediaType: d.Get("default_media_type").(string),
	}
	return &config, nil
}

type Config struct {
	DefaultMediaType string
}
