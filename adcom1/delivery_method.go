package adcom1

// DeliveryMethod represents options for the delivery of video or audio content (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_deliverymethods
type DeliveryMethod int8

// Options for the delivery of video or audio content.
const (
	DeliveryStreaming   DeliveryMethod = 1 // Streaming
	DeliveryProgressive DeliveryMethod = 2 // Progressive
	DeliveryDownload    DeliveryMethod = 3 // Download
)
