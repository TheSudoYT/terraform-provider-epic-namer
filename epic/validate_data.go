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

// validate_data.go ensures that the map of media_types to titles is cached during the plan.
// I'm not actually sure if this carries over to the apply?
// It also validates that the media_types in the /data directory contain a title provided as an input.

var mediaTypeCache map[string]map[string][]string
var cacheLoaded bool

// MediaTypeData structure holds the title and names read from each JSON file.
type MediaTypeData struct {
	Title string   `json:"title"`
	Names []string `json:"names"`
}

// FOR TESTING ONLY!!
func getDataDirPath() string {
	dataDir := os.Getenv("DATA_DIR") // Set to "../data" for go tests.
	if dataDir == "" {
		dataDir = "data" // Default is "data" because that is what is required for users consuming the provider.
	}
	return dataDir
}

// LoadMediaTypes scans the specified dataDir directory, reads each subdirectory as a media type,
// and loads each JSON file in those subdirectories as titles of that media type.
func LoadMediaTypes(dataDir string) (map[string]map[string][]string, error) {

	// This map will hold the media type as key and a map of titles with their names as value.
	mediaTypes := make(map[string]map[string][]string)

	// Read the main data directory.
	dataDirPath := filepath.FromSlash(dataDir)
	entries, err := os.ReadDir(dataDirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read data directory: %v", err)
	}

	// Iterate over each entry in the data directory.
	for _, entry := range entries {
		if entry.IsDir() {
			mediaType := entry.Name()
			mediaPath := filepath.Join(dataDirPath, mediaType)
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
					// I actually don't know if this will cause unexpected outcomes for users? We shall see.
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

// Function to load media_type to title mapping from cache if the cache is not considered loaded already.
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

// Function called by customizeDiff: during resource creation.
func customValidateMediaTypeAndTitle(ctx context.Context, diff *schema.ResourceDiff, v interface{}) error {
	mediaType, ok := diff.Get("media_type").(string)
	if !ok {
		fmt.Println("Expceted a media type. Found none")
	}

	title, ok := diff.Get("title").(string)
	if !ok {
		fmt.Println("Expceted a title. Found none")
	}

	isValid, errMsg := isValidMediaTypeAndTitle(mediaType, title)
	if !isValid {
		return fmt.Errorf(errMsg)
	}

	return nil
}

// Function to validate that the user provided media_type and title are valid or not.
func isValidMediaTypeAndTitle(mediaType, title string) (bool, string) {
	if !cacheLoaded {
		if err := LoadAndCacheMediaTypes("data"); err != nil {
			return false, fmt.Sprintf("error loading media types: %v", err)
		}
	}

	// Check if the media_type is valid.
	titlesMap, ok := mediaTypeCache[mediaType]
	if !ok {
		return false, fmt.Sprintf("'%s' is not a recognized media type", mediaType)
	}

	// Check if the title is valid for the given media_type.
	for t := range titlesMap {
		if strings.EqualFold(t, title) {
			return true, ""
		}
	}

	return false, fmt.Sprintf("'%s' is not a valid title for media type '%s'", title, mediaType)
}
