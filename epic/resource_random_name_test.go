package epic

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// Acceptance tests. Set TF_ACC=1 env variable to enable.
func TestAccEpicRandomName_basic(t *testing.T) {

	// $env:DATA_DIR = "../data" or export DATA_DIR="../data"
	dataDir := getDataDirPath()

	if err := LoadAndCacheMediaTypes(dataDir); err != nil {
		t.Fatalf("Failed to load media types: %v", err)
	}

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
			{
				Config:      testAccCheckEpicRandomNameConfig_invalidMediaType(),
				ExpectError: regexp.MustCompile(`'not_a_real_media' is not a recognized media type`),
			},
			{
				Config:      testAccCheckEpicRandomNameConfig_invalidTitle(),
				ExpectError: regexp.MustCompile(`'fake_title' is not a valid title for media type 'movie'`),
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

func testAccCheckEpicRandomNameConfig_invalidMediaType() string {
	return `
provider "epic" {}

resource "epic_random_name" "test_invalid_media_type" {
    media_type = "not_a_real_media"
    title      = "lord of the rings"
}
`
}

func testAccCheckEpicRandomNameConfig_invalidTitle() string {
	return `
provider "epic" {}

resource "epic_random_name" "test_invalid_title" {
    media_type = "movie"
    title      = "fake_title"
}
`
}
