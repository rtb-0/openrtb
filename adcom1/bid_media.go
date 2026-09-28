package adcom1

// BidMedia represents a bunch of data for OpenRTB 3 Bid.media field (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_bid
type BidMedia[A any] struct {
	Ad *Ad[A] `json:"ad,omitempty"`
}
