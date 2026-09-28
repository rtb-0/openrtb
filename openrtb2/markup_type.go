package openrtb2

// MarkupType defines the type of the creative markup so that it can properly be associated with the right sub-object of the BidRequest.Imp (OpenRTB 2.6).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectbid
type MarkupType int8

const (
	MarkupBanner MarkupType = 1
	MarkupVideo  MarkupType = 2
	MarkupAudio  MarkupType = 3
	MarkupNative MarkupType = 4
)
