package openrtb2

// AuctionType is the auction type (OpenRTB 2.6 §3.2.1).
// Values of 500 and greater are exchange-specific.
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectbidrequest
type AuctionType int16

const (
	FirstPrice  AuctionType = 1
	SecondPrice AuctionType = 2 // Second Price Plus.
	DealPrice   AuctionType = 3
)
