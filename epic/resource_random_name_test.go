package epic

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccEpicRandomName_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]*schema.Provider{
			"epic": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCheckEpicRandomNameConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEpicRandomNameExists("epic_random_name.test"),
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

func testAccCheckEpicRandomNameExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource ID is set")
		}

		return nil
	}
}
