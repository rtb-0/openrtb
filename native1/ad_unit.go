package native1

// Deprecated/removed in version 1.2 (OpenRTB Native 1.2 §7.2).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#4-1
// Values of 500 and greater are exchange-specific.
type AdUnit int16

const (
	AdUnitPaidSearch           AdUnit = 1 // Paid Search Units
	AdUnitRecommendationWidget AdUnit = 2 // Recommendation Widgets
	AdUnitPromotedListing      AdUnit = 3 // Promoted Listings
	AdUnitInAd                 AdUnit = 4 // In-Ad (IAB Standard) with Native Element Units
	AdUnitCustom               AdUnit = 5 // Custom /”Can’t Be Contained”
)
