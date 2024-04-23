package epic

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var mediaTypeCache map[string]map[string][]string
var cacheLoaded bool

// MediaTypeData structure holds the title and names read from each JSON file.
type MediaTypeData struct {
	Title string   `json:"title"`
	Names []string `json:"names"`
}

// LoadMediaTypes scans the specified dataDir directory, reads each subdirectory as a media type,
// and loads each JSON file in those subdirectories as titles of that media type.
func LoadMediaTypes(dataDir string) (map[string]map[string][]string, error) {
	// This map will hold the media type as key and a map of titles with their names as value.
	mediaTypes := make(map[string]map[string][]string)

	// Read the main data directory.
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read data directory: %v", err)
	}

	// Iterate over each entry in the data directory.
	for _, entry := range entries {
		if entry.IsDir() {
			mediaType := entry.Name()
			mediaPath := filepath.Join(dataDir, mediaType)
			mediaTypeMap := make(map[string][]string)

			// Read files within the media type directory.
			files, err := os.ReadDir(mediaPath)
			if err != nil {
				return nil, fmt.Errorf("failed to read directory for media type '%s': %v", mediaType, err)
			}

			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
					filePath := filepath.Join(mediaPath, file.Name())
					fileData, err := os.ReadFile(filePath)
					if err != nil {
						return nil, fmt.Errorf("failed to read file '%s': %v", filePath, err)
					}

					var mediaData MediaTypeData
					if err := json.Unmarshal(fileData, &mediaData); err != nil {
						return nil, fmt.Errorf("failed to unmarshal data for file '%s': %v", filePath, err)
					}

					// Use the title from the JSON or the filename without extension as a fallback.
					title := mediaData.Title
					if title == "" {
						title = strings.TrimSuffix(file.Name(), ".json")
					}

					// Store the names under the title within the media type map.
					mediaTypeMap[title] = mediaData.Names
				}
			}

			// Store the media type map in the main map if it contains any titles.
			if len(mediaTypeMap) > 0 {
				mediaTypes[mediaType] = mediaTypeMap
			}
		}
	}

	return mediaTypes, nil
}

func LoadAndCacheMediaTypes(dataDir string) error {
	if !cacheLoaded {
		var err error
		mediaTypeCache, err = LoadMediaTypes(dataDir)
		if err != nil {
			return err
		}
		cacheLoaded = true
	}
	return nil
}

func customValidateMediaTypeAndTitle(ctx context.Context, diff *schema.ResourceDiff, v interface{}) error {
	mediaType := diff.Get("media_type").(string)
	title := diff.Get("title").(string)

	if !isValidMediaTypeAndTitle(mediaType, title) {
		return fmt.Errorf("'%s' is not a valid title for media type '%s'", title, mediaType)
	}

	return nil
}

func isValidMediaTypeAndTitle(mediaType, title string) bool {
	if !cacheLoaded {
		if err := LoadAndCacheMediaTypes("data"); err != nil {
			fmt.Printf("Error loading media types: %v\n", err)
			return false
		}
	}

	titlesMap, ok := mediaTypeCache[mediaType]
	if !ok {
		return false
	}

	for t := range titlesMap {
		if strings.EqualFold(t, title) {
			return true
		}
	}

	return false
}
