package native1

// Next-level context in which the ad appears (OpenRTB Native 1.2 §7.2).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#7-2
// Values of 500 and greater are exchange-specific.
type ContextSubType int16

const (
	ContextSubTypeGeneral       ContextSubType = 10 // General or mixed content.
	ContextSubTypeArticle       ContextSubType = 11 // Primarily article content (which of course could include images, etc as part of the article)
	ContextSubTypeVideo         ContextSubType = 12 // Primarily video content
	ContextSubTypeAudio         ContextSubType = 13 // Primarily audio content
	ContextSubTypeImage         ContextSubType = 14 // Primarily image content
	ContextSubTypeUserGenerated ContextSubType = 15 // User-generated content - forums, comments, etc
	ContextSubTypeSocial        ContextSubType = 20 // General social content such as a general social network
	ContextSubTypeEmail         ContextSubType = 21 // Primarily email content
	ContextSubTypeChat          ContextSubType = 22 // Primarily chat/IM content
	ContextSubTypeSelling       ContextSubType = 30 // Content focused on selling products, whether digital or physical
	ContextSubTypeAppStore      ContextSubType = 31 // Application store/marketplace
	ContextSubTypeProductReview ContextSubType = 32 // Product reviews site primarily (which may sell product secondarily)
)
