package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Content object describes the content in which an impression can appear, which may be syndicated or non-syndicated content (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_content
type Content struct {
	// ID uniquely identifying the content.
	ID string `json:"id,omitempty"`
	// Episode number.
	Episode int64 `json:"episode,omitempty"`
	// Content title.
	Title string `json:"title,omitempty"`
	// Content series.
	Series string `json:"series,omitempty"`
	// Content season (e.g., “Season 3”).
	Season string `json:"season,omitempty"`
	// Artist credited with the content.
	Artist string `json:"artist,omitempty"`
	// Genre that best describes the content (e.g., rock, pop, etc).
	Genre string `json:"genre,omitempty"`
	// Genre IDs from the taxonomy in gtax.
	Genres []string `json:"genres,omitempty"`
	// Taxonomy used for genres. Omitted means Content Category Taxonomy 3.1.
	GTax CategoryTaxonomy `json:"gtax,omitempty"`
	// Album to which the content belongs; typically for audio.
	Album string `json:"album,omitempty"`
	// International Standard Recording Code conforming to ISO-3901.
	ISRC string `json:"isrc,omitempty"`
	// URL of the content, for buy-side contextualization or review.
	URL string `json:"url,omitempty"`
	// Array of content categories describing the content using IDs from the taxonomy indicated in cattax.
	Cat []string `json:"cat,omitempty"`
	// The taxonomy in use for the cat attribute.
	CatTax CategoryTaxonomy `json:"cattax,omitempty"`
	// Production quality.
	ProdQ ProductionQuality `json:"prodq,omitempty"`
	// Type of content (game, video, text, etc.).
	Context ContentContext `json:"context,omitempty"`
	// Content rating (e.g., MPAA).
	Rating string `json:"rating,omitempty"`
	// User rating of the content (e.g., number of stars, likes, etc.).
	URating string `json:"urating,omitempty"`
	// Media rating per IQG guidelines.
	MRating MediaRating `json:"mrating,omitempty"`
	// Comma-separated list of keywords describing the content.
	Keywords string `json:"keywords,omitempty"`
	// Array of keywords about the site.
	KwArray []string `json:"kwarray,omitempty"`
	// Indication of live content, where 0 = not live, 1 = live (e.g., stream, live blog).
	Live int8 `json:"live,omitempty"`
	// Source relationship, where 0 = indirect, 1 = direct.
	SrcRel int8 `json:"srcrel,omitempty"`
	// Length of content in seconds; typically for video or audio.
	Len int64 `json:"len,omitempty"`
	// Content language using ISO-639-1-alpha-2.
	Lang string `json:"lang,omitempty"`
	// Content language using IETF BCP 47.
	LangB string `json:"langb,omitempty"`
	// Indicator of whether or not the content is embedded off-site from the the site or app described in those objects (e.g., an embedded video player), where 0 = no, 1 = yes.
	Embed int8 `json:"embed,omitempty"`
	// Details about the content producer.
	Producer *Producer `json:"producer,omitempty"`
	// Details about the network.
	Network *Network `json:"network,omitempty"`
	// Details about the channel.
	Channel *Channel `json:"channel,omitempty"`
	// Additional user data.
	Data []Data `json:"data,omitempty"`
	// 0 = replay or otherwise not real time, 1 = happening in real time.
	RealTime *int8 `json:"realtime,omitempty"`
	// 0 = not the first broadcast, 1 = first time the content is available.
	FirstBroadcast *int8 `json:"firstbroadcast,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
