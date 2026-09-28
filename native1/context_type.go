package native1

// ContextType is the context in which the ad appears (OpenRTB Native 1.2 §7.1).
// Values of 500 and greater are exchange-specific.
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#7-1
type ContextType int16

const (
	ContextTypeContent ContextType = 1 // Content-centric context such as newsfeed, article, image gallery, video gallery, or similar.
	ContextTypeSocial  ContextType = 2 // Social-centric context such as social network feed, email, chat, or similar.
	ContextTypeProduct ContextType = 3 // Product context such as product listings, details, recommendations, reviews, or similar.
)
