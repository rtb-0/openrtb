package openrtb3

// AuctionType defines an auction type (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_request
// Values of 500 and greater are exchange-specific.
type AuctionType int16

// AuctionType values.
const (
	FirstPrice      AuctionType = 1
	SecondPricePlus             = 2
	DealPrice                   = 3 // the value passed in flr is the agreed upon deal price
)
