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
	}
}

type Config struct {
	DefaultMediaType string
}
