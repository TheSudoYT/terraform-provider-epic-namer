package epic

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRandomName() *schema.Resource {
	return &schema.Resource{
		Description: "Generates a random character name based on the media type and title specified.",
		Create:      resourceRandomNameCreate,
		Read:        schema.Noop,
		Delete:      schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"media_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of media, e.g., 'movie' or 'tv_series'.",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The title of the media to base the name generation on.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The randomly generated name.",
			},
		},
	}
}

func resourceRandomNameCreate(d *schema.ResourceData, m interface{}) error {
	mediaType := d.Get("media_type").(string)
	title := d.Get("title").(string)

	rand.Seed(time.Now().UnixNano())
	var names []string

	switch mediaType {
	case "movie":
		names = getNamesForMovie(title)
	case "tv_series":
		names = getNamesForTVSeries(title)
	default:
		return fmt.Errorf("unsupported media type: %s", mediaType)
	}

	selectedName := names[rand.Intn(len(names))]

	d.SetId(strconv.FormatInt(time.Now().UnixNano(), 10))
	d.Set("name", selectedName)

	return nil
}

func getNamesForMovie(title string) []string {
	// Placeholder: return a slice of names based on the movie title
	// Example implementation
	switch title {
	case "lord of the rings":
		return []string{"Aragorn", "Gandalf", "Bilbo", "Frodo", "Legolas"}
	default:
		return []string{"John Doe"}
	}
}

func getNamesForTVSeries(title string) []string {
	// Placeholder: return a slice of names based on the TV series title
	// Example implementation
	switch title {
	case "game of thrones":
		return []string{"Jon Snow", "Tyrion Lannister", "Daenerys Targaryen"}
	default:
		return []string{"Jane Doe"}
	}
}
