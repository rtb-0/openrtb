package adcom1

// EventTrackingMethod represents methods of tracking of ad events (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_eventtrackingmethods
type EventTrackingMethod int16

// Available methods of tracking of ad events.
const (
	TrackingImagePixel EventTrackingMethod = 1 // Image-Pixel: URL provided will be inserted as a 1x1 pixel at the time of the event.
	TrackingJS         EventTrackingMethod = 2 // JavaScript: URL provided will be inserted as a JavaScript tag at the time of the event.
)
