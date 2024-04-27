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

type QuoteData struct {
	Title  string   `json:"title"`
	Quotes []string `json:"quotes"`
}

func resourceRandomQuote() *schema.Resource {
	return &schema.Resource{
		Description: "Generates a random quote based on the media type and title specified.",
		Create:      resourceRandomQuoteCreate,
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
				Description: "The title of the media to base the quote generation on.",
			},
			"quote": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The randomly generated quote.",
			},
		},
		CustomizeDiff: customValidateMediaTypeAndTitle,
	}
}

func loadQuotes(mediaType, title string) ([]string, error) {
	sanitizedTitle := strings.ReplaceAll(title, " ", "_")
	fileName := fmt.Sprintf("%s.json", sanitizedTitle)
	dataDirPath := filepath.FromSlash(getDataDirPath())
	filePath := filepath.Join(dataDirPath, mediaType, fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %v", filePath, err)
	}

	var quoteData QuoteData
	if err := json.Unmarshal(data, &quoteData); err != nil {
		return nil, fmt.Errorf("failed to parse quotes from %s: %v", filePath, err)
	}

	return quoteData.Quotes, nil
}

func resourceRandomQuoteCreate(d *schema.ResourceData, m interface{}) error {
	mediaType, ok := d.Get("media_type").(string)
	if !ok {
		fmt.Println("Expceted a media_type. Found none")
	}

	title, ok := d.Get("title").(string)
	if !ok {
		fmt.Println("Expceted a title. None found.")
	}

	quotes, err := loadQuotes(mediaType, title)
	if err != nil {
		return fmt.Errorf("error loading quotes for %s '%s': %s", mediaType, title, err)
	}

	if len(quotes) == 0 {
		return fmt.Errorf("no quotes found for %s '%s'", mediaType, title)
	}

	// Setup a local random source.
	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)
	selectedQuote := quotes[localRand.Intn(len(quotes))]

	// Set the resource ID and the computed quote.
	d.SetId(strconv.FormatInt(time.Now().UnixNano(), 10))

	if err := d.Set("quote", selectedQuote); err != nil {
		log.Fatalf("Error setting quote: %v", err)
	}

	return nil
}
