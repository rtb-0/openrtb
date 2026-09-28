package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
	nresp "github.com/rtb-0/openrtb/native1/response"
)

// A SeatBid object contains one or more Bid objects, each of which relates to a specific impression in the bid request via the impid attribute and constitutes an offer to buy that impression for a given price (OpenRTB 2.6 §4.2.3).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectbid
type Bid[A any] struct {
	// Bidder generated bid ID to assist with logging/tracking.
	ID string `json:"id"`
	// ID of the Imp object in the related bid request.
	ImpID string `json:"impid"`
	// Bid price expressed as CPM although the actual transaction is for a unit impression only.
	Price float64 `json:"price"`
	// Win notice URL called by the exchange if the bid wins (not necessarily indicative of a delivered, viewed, or billable ad); optional means of serving ad markup.
	NURL string `json:"nurl,omitempty"`
	// Billing notice URL called by the exchange when a winning bid becomes billable based on exchange-specific business policy (e.g., typically delivered, viewed, etc.).
	BURL string `json:"burl,omitempty"`
	// Loss notice URL called by the exchange when a bid is known to have been lost.
	LURL string `json:"lurl,omitempty"`
	// Optional means of conveying ad markup in case the bid wins; supersedes the win notice if markup is included in both.
	// JSON stores this as a string. A parses that string.
	Adm common.StringEncoded[A] `json:"adm,omitzero"`
	// Native markup as an object. The root native node is omitted (OpenRTB 2.6).
	AdMNative *nresp.Response `json:"adm_native,omitzero"`
	// ID of a preloaded ad to be served if the bid wins.
	AdID string `json:"adid,omitempty"`
	// Advertiser domain for block list checking (e.g., “ford.com”).
	ADomain []string `json:"adomain,omitempty"`
	// A platform-specific application identifier intended to be unique to the app and independent of the exchange.
	Bundle string `json:"bundle,omitempty"`
	// URL without cache-busting to an image that is representative of the content of the campaign for ad quality/safety checking.
	IURL string `json:"iurl,omitempty"`
	// Campaign ID to assist with ad quality checking; the collection of creatives for which iurl should be representative.
	CID string `json:"cid,omitempty"`
	// Creative ID to assist with ad quality checking
	CrID string `json:"crid,omitempty"`
	// Tactic ID to enable buyers to label bids for reporting to the exchange the tactic through which their bid was submitted.
	Tactic string `json:"tactic,omitempty"`
	// The taxonomy in use.
	CatTax adcom1.CategoryTaxonomy `json:"cattax,omitempty"`
	// IAB content categories of the creative.
	Cat []string `json:"cat,omitempty"`
	// Set of attributes describing the creative.
	Attr []adcom1.CreativeAttribute `json:"attr,omitempty"`
	// List of supported APIs for the markup.
	APIs []adcom1.APIFramework `json:"apis,omitempty"`
	// NOTE: Deprecated in favor of the apis integer array.
	API adcom1.APIFramework `json:"api,omitempty"`
	// Video response protocol of the markup if applicable.
	Protocol adcom1.MediaCreativeSubtype `json:"protocol,omitempty"`
	// Media rating per IQG guidelines.
	QAGMediaRating adcom1.MediaRating `json:"qagmediarating,omitempty"`
	// Language of the creative using ISO-639-1-alpha-2.
	Language string `json:"language,omitempty"`
	// Language of the creative using IETF BCP 47.
	LangB string `json:"langb,omitempty"`
	// Reference to the deal.id from the bid request if this bid pertains to a private marketplace direct deal.
	DealID string `json:"dealid,omitempty"`
	// Width of the creative in device independent pixels (DIPS).
	W int64 `json:"w,omitempty"`
	// Height of the creative in device independent pixels (DIPS).
	H int64 `json:"h,omitempty"`
	// Relative width of the creative when expressing size as a ratio.
	WRatio int64 `json:"wratio,omitempty"`
	// Relative height of the creative when expressing size as a ratio.
	HRatio int64 `json:"hratio,omitempty"`
	// Advisory as to the number of seconds the bidder is willing to wait between the auction and the actual impression.
	Exp int64 `json:"exp,omitempty"`
	// Duration of the video or audio creative in seconds.
	Dur int64 `json:"dur,omitempty"`
	// Type of the creative markup so that it can properly be associated with the right sub-object of the BidRequest.Imp.
	MType MarkupType `json:"mtype,omitempty"`
	// Indicates that the bid response is only eligible for a specific position within a video or audio ad pod (e.g. first position, last position, or any).
	SlotInPod adcom1.SlotPositionInPod `json:"slotinpod,omitempty"`
	// Placeholder for bidder-specific extensions to OpenRTB
	Ext common.Ext `json:"ext,omitempty"`
}
