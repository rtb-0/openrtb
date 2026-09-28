package openrtb3

import (
	"github.com/rtb-0/openrtb/common"
)

// Source object carries data about the source of the transaction including the unique ID of the transaction itself, source authentication information, and the chain of custody (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_source
type Source struct {
	// Transaction ID that must be common across all participants throughout the entire supply chain of this transaction.
	TID string `json:"tid,omitempty"`
	// Timestamp when the request originated at the beginning of the supply chain in Unix format (i.e., milliseconds since the epoch).
	TS int64 `json:"ts,omitempty"`
	// Digital signature used to authenticate the origin of this request computed by the publisher or its trusted agent from a digest string composed of a set of immutable attributes found in the bid request.
	DS string `json:"ds,omitempty"`
	// An ordered list of identifiers that indicates the attributes used to create the digest.
	DSMap string `json:"dsmap,omitempty"`
	// File name of the certificate (i.e., the public key) used to generate the digital signature in the ds attribute.
	Cert string `json:"cert,omitempty"`
	// The full digest string that was signed to produce the digital signature.
	Digest string `json:"digest,omitempty"`
	// Supply chain links and whether the chain is complete.
	SChain *SupplyChain `json:"schain,omitempty"`
	// Payment ID chain string containing embedded syntax described in the TAG Payment ID Protocol.
	PChain string `json:"pchain,omitempty"`
	// Optional exchange-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
