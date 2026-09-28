package response

// Corresponds to the Video Object in the request, yet containing a value of a conforming VAST tag as a value (OpenRTB Native 1.2 §5.6).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#5-6
type Video struct {
	// VAST XML
	VASTTag string `json:"vasttag"`
}
