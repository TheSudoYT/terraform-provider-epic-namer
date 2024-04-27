---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "epic_random_quote Resource - epic"
subcategory: ""
description: |-
  Generates a random quote based on the media type and title specified.
---

# epic_random_quote (Resource)

Generates a random quote based on the media type and title specified.

## Example Usage

```terraform
resource "epic_random_quote" "got_quote" {
  media_type = "tv_series"
  title      = "game of thrones"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `media_type` (String) The type of media, e.g., 'movie' or 'tv_series'.
- `title` (String) The title of the media to base the quote generation on.

### Read-Only

- `id` (String) The ID of this resource.
- `quote` (String) The randomly generated quote.