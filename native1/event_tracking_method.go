package native1

// EventTrackingMethod is defined by OpenRTB Native 1.2 (OpenRTB Native 1.2 §7.7).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#7-7
// Values of 500 and greater are exchange-specific.
type EventTrackingMethod int16

const (
	EventTrackingMethodImage EventTrackingMethod = 1 // Image-pixel tracking - URL provided will be inserted as a 1x1 pixel at the time of the event.
	EventTrackingMethodJS    EventTrackingMethod = 2 // Javascript-based tracking - URL provided will be inserted as a js tag at the time of the event.
)
