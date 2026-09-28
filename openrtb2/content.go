package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object describes the content in which the impression will appear, which may be syndicated or non-syndicated content (OpenRTB 2.6 §3.2.16).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectcontent
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
	// Taxonomy used for genres. Omitted means Content Category Taxonomy 3.1.
	GTax adcom1.CategoryTaxonomy `json:"gtax,omitempty"`
	// Genre IDs from the taxonomy in gtax.
	Genres []string `json:"genres,omitempty"`
	// Album to which the content belongs; typically for audio.
	Album string `json:"album,omitempty"`
	// International Standard Recording Code conforming to ISO- 3901.
	ISRC string `json:"isrc,omitempty"`
	// Details about the content Producer (Section 3.2.17).
	Producer *Producer `json:"producer,omitempty"`
	// URL of the content, for buy-side contextualization or review.
	URL string `json:"url,omitempty"`
	// The taxonomy in use.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`
	// Array of IAB content categories that describe the content.
	Cat []string `json:"cat,omitempty"`
	// Production quality.
	ProdQ *adcom1.ProductionQuality `json:"prodq,omitempty"`
	// Removed in OpenRTB 2.6. Kept so OpenRTB 2.5 content still decodes.
	VideoQuality *adcom1.ProductionQuality `json:"videoquality,omitempty"`
	// Type of content (game, video, text, etc.).
	Context adcom1.ContentContext `json:"context,omitempty"`
	// Content rating (e.g., MPAA).
	ContentRating string `json:"contentrating,omitempty"`
	// User rating of the content (e.g., number of stars, likes, etc.).
	UserRating string `json:"userrating,omitempty"`
	// Media rating per IQG guidelines.
	QAGMediaRating adcom1.MediaRating `json:"qagmediarating,omitempty"`
	// Comma separated list of keywords describing the content.
	Keywords string `json:"keywords,omitempty"`
	// Array of keywords about the site.
	KwArray []string `json:"kwarray,omitempty"`
	// 0 = not scheduled (VOD or user initiated), 1 = scheduled linear broadcast.
	LiveStream *int8 `json:"livestream,omitempty"`
	// 0 = indirect, 1 = direct.
	SourceRelationship *int8 `json:"sourcerelationship,omitempty"`
	// Length of content in seconds; appropriate for video or audio.
	Len int64 `json:"len,omitempty"`
	// Content language using ISO-639-1-alpha-2.
	Language string `json:"language,omitempty"`
	// Content language using IETF BCP 47.
	LangB string `json:"langb,omitempty"`
	// Indicator of whether or not the content is embeddable (e.g., an embeddable video player), where 0 = no, 1 = yes.
	Embeddable *int8 `json:"embeddable,omitempty"`
	// Additional content data.
	Data []Data `json:"data,omitempty"`
	// Details about the network (Section 3.2.23) the content is on.
	Network *Network `json:"network,omitempty"`
	// Details about the channel (Section 3.2.24) the content is on.
	Channel *Channel `json:"channel,omitempty"`
	// 0 = replay or otherwise not real time, 1 = happening in real time.
	RealTime *int8 `json:"realtime,omitempty"`
	// 0 = not the first broadcast, 1 = first time the content is available.
	FirstBroadcast *int8 `json:"firstbroadcast,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
