package native1

// Deprecated/removed in version 1.2 (OpenRTB Native 1.2 §7.1).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#4-1
// Values of 500 and greater are exchange-specific.
type Layout int16

const (
	LayoutContentWall   Layout = 1 // Content Wall
	LayoutAppWall       Layout = 2 // App Wall
	LayoutNewsFeed      Layout = 3 // News Feed
	LayoutChatList      Layout = 4 // Chat List
	LayoutCarousel      Layout = 5 // Carousel
	LayoutContentStream Layout = 6 // Content Stream
	LayoutGrid          Layout = 7 // Grid adjoining the content
)
