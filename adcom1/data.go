package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Data and segment objects together allow additional data about the related object (e.g., user, content) to be specified (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_data
type Data struct {
	// Vendor-specific ID for the data provider.
	ID string `json:"id,omitempty"`
	// Vendor-specific displayable name for the data provider.
	Name string `json:"name,omitempty"`
	// Extended content IDs from the source named in name.
	CIds []string `json:"cids,omitempty"`
	// Array of Segment objects that contain the actual data values.
	Segment []Segment `json:"segment,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
