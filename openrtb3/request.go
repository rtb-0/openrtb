package openrtb3

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// Request object contains a globally unique bid request ID (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_request
type Request struct {
	// Unique ID of the bid request; provided by the exchange.
	ID string `json:"id"`
	// Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode.
	Test int8 `json:"test,omitempty"`
	// Maximum time in milliseconds the exchange allows for bids to be received including Internet latency to avoid timeout.
	TMax int64 `json:"tmax,omitempty"`
	// Auction type, where 1 = First Price, 2 = Second Price Plus.
	AT AuctionType `json:"at,omitempty"`
	// Array of accepted currencies for bids on this bid request using ISO-4217 alpha codes.
	Cur []string `json:"cur,omitempty"`
	// Restriction list of buyer seats for bidding on this item.
	Seat []string `json:"seat,omitempty"`
	// Flag that determines the restriction interpretation of the seat array, where 0 = block list, 1 = whitelist.
	WSeat int8 `json:"wseat,omitempty"`
	// Allows bidder to retrieve data set on its behalf in the exchange’s cookie (refer to cdata in Object: Response) if supported by the exchange.
	CData string `json:"cdata,omitempty"`
	// A Source object that provides data about the inventory source and which entity makes the final decision Refer to Object: Source.
	Source *Source `json:"source,omitempty"`
	// Array of Item objects (at least one) that constitute the set of goods being offered for sale.
	Item []Item `json:"item"`
	// Flag to indicate if the Exchange can verify that the items offered represent all of the items available in context (e.g., all impressions on a web page, all video spots such as pre/mid/post roll) to support road-blocking, where 0 = no, 1 = yes.
	Package int8 `json:"package,omitempty"`
	// Layer-4 domain object structure that provides context for the items being offered conforming to the specification and version referenced in openrtb.domainspec and openrtb.domainver.
	Context *adcom1.RequestContext `json:"context,omitempty"`
	// Optional exchange-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
