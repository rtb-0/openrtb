package openrtb2

import (
	"github.com/rtb-0/openrtb/common"
)

// The data and segment objects together allow additional data about the related object (e.g., user, content) to be specified (OpenRTB 2.6 §3.2.21).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectdata
type Data struct {
	// Exchange-specific ID for the data provider.
	ID string `json:"id,omitempty"`
	// Exchange-specific name for the data provider.
	Name string `json:"name,omitempty"`
	// Extended content IDs from the source named in name.
	CIds []string `json:"cids,omitempty"`
	// Array of Segment (Section 3.2.22) objects that contain the actual data values.
	Segment []Segment `json:"segment,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
