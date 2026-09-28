package native1

// EventType is defined by OpenRTB Native 1.2 (OpenRTB Native 1.2 §7.6).
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md#7-6
// Values of 500 and greater are exchange-specific.
type EventType int16

const (
	EventTypeImpression      EventType = 1 // Impression
	EventTypeViewableMRC50   EventType = 2 // Visible impression using MRC definition at 50% in view for 1 second
	EventTypeViewableMRC100  EventType = 3 // 100% in view for 1 second (ie GroupM standard)
	EventTypeViewableVideo50 EventType = 4 // Visible impression for video using MRC definition at 50% in view for 2 seconds
)
