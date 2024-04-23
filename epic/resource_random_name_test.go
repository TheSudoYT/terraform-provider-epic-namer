package epic

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Acceptance tests. Set TF_ACC=1 env variable to enable.
func TestAccEpicRandomName_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"epic": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckEpicRandomNameConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epic_random_name.test", "media_type", "movie"),
					resource.TestCheckResourceAttr("epic_random_name.test", "title", "lord of the rings"),
				),
			},
		},
	})
}

func testAccCheckEpicRandomNameConfig_basic() string {
	return `
provider "epic" {}

resource "epic_random_name" "test" {
    media_type = "movie"
    title      = "lord of the rings"
}
`
}
