package native1

// The FORMAT of the ad you are purchasing, separate from the surrounding context (OpenRTB Native 1.2 §7.3).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#7-3
// Values of 500 and greater are exchange-specific.
type PlacementType int16

const (
	PlacementTypeFeed                 = 1 // In the feed of content - for example as an item inside the organic feed/grid/listing/carousel.
	PlacementTypeAtomicContentUnit    = 2 // In the atomic unit of the content - IE in the article page or single image page
	PlacementTypeOutsideCoreContent   = 3 // Outside the core content - for example in the ads section on the right rail, as a banner-style placement near the content, etc.
	PlacementTypeRecommendationWidget = 4 // Recommendation widget, most commonly presented below the article content.
)
