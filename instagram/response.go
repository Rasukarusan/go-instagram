package api

/**
 * InstagramのHTML内JSONの構造体
 */
type InstagramResponse struct {
	EntryData EntryData `json:"entry_data"`
}

type EntryData struct {
	PostPage []struct {
		Graphql struct {
			ShortCodeMedia ShortCodeMedia `json:"shortcode_media"`
		}
	}
}

type ShortCodeMedia struct {
	DisplayURL string `json:"display_url"`
	Owner      struct {
		Username string
	}
	EdgeMediaToCaption struct {
		Edges []struct {
			Node struct {
				Text string
			}
		}
	} `json:"edge_media_to_caption"`
}
