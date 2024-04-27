package epic

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TitleData struct {
	Title string   `json:"title"`
	Names []string `json:"names"`
}

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
		CustomizeDiff: customValidateMediaTypeAndTitle,
	}
}

func loadNames(mediaType, title string) ([]string, error) {
	sanitizedTitle := strings.ReplaceAll(title, " ", "_")
	fileName := fmt.Sprintf("%s.json", sanitizedTitle)
	dataDirPath := filepath.FromSlash(getDataDirPath())
	filePath := filepath.Join(dataDirPath, mediaType, fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %v", filePath, err)
	}

	var titleData TitleData
	if err := json.Unmarshal(data, &titleData); err != nil {
		return nil, fmt.Errorf("failed to parse names from %s: %v", filePath, err)
	}

	return titleData.Names, nil
}

func resourceRandomNameCreate(d *schema.ResourceData, m interface{}) error {
	mediaType, ok := d.Get("media_type").(string)
	if !ok {
		fmt.Println("Expceted a media type. Found none")
	}

	title, ok := d.Get("title").(string)
	if !ok {
		fmt.Println("Expceted a title. Found none")
	}

	names, err := loadNames(mediaType, title)
	if err != nil {
		return fmt.Errorf("error loading names for %s '%s': %s", mediaType, title, err)
	}

	if len(names) == 0 {
		return fmt.Errorf("no names found for %s '%s'", mediaType, title)
	}

	// Setup a local random source.
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)
	selectedName := names[localRand.Intn(len(names))]

	// Set the resource ID and the computed name.
	d.SetId(strconv.FormatInt(time.Now().UnixNano(), 10))

	if err := d.Set("name", selectedName); err != nil {
		log.Fatalf("Error setting name: %v", err)
	}

	return nil
}
