package openrtb3

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// Item object represents a unit of goods being offered for sale either on the open market or in relation to a private marketplace deal (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_item
type Item struct {
	// A unique identifier for this item within the context of the offer (typically starts with “1” and increments).
	ID string `json:"id"`
	// Quantity of billable events, as a whole number. Only one of qty or qtyflt may be present.
	Qty int64 `json:"qty,omitempty"`
	// Quantity of billable events when it is not a whole number. Only one of qty or qtyflt may be present.
	QtyFlt float64 `json:"qtyflt,omitempty"`
	// If multiple items are offered in the same bid request, the sequence number allows for the coordinated delivery.
	Seq int64 `json:"seq,omitempty"`
	// Minimum bid price for this item expressed in CPM.
	Flr float64 `json:"flr,omitempty"`
	// Currency of the flr attribute specified using ISO-4217 alpha codes.
	FlrCur string `json:"flrcur,omitempty"`
	// Advisory as to the number of seconds that may elapse between auction and fulfilment.
	Exp int64 `json:"exp,omitempty"`
	// Timestamp when the item is expected to be fulfilled (e.g when a DOOH impression will be displayed) in Unix format (i.e., milliseconds since the epoch).
	DT int64 `json:"dt,omitempty"`
	// Item (e.g., an Ad object) delivery method required, where 0 = either method, 1 = the item must be sent as part of the transaction (e.g., by value in the bid itself, fetched by URL included in the bid), and 2 = an item previously uploaded to the exchange must be referenced by its ID.
	Dlvy int8 `json:"dlvy,omitempty"`
	// An array of Metric objects.
	Metric []Metric `json:"metric,omitempty"`
	// Array of Deal objects that convey special terms applicable to this item.
	Deal []Deal `json:"deal,omitempty"`
	// Indicator of auction eligibility to seats named in Deal objects, where 0 = all bids are accepted, 1 = bids are restricted to the deals specified and the terms thereof.
	Private int8 `json:"private,omitempty"`
	// Layer-4 domain object structure that provides specifies the item being offered conforming to the specification and version referenced in openrtb.domainspec and openrtb.domainver.
	Spec *adcom1.ItemSpec `json:"spec"`
	// Popunder marks an item that intentionally has no placement subtype.
	Popunder bool `json:"-"`
	// Optional exchange-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
