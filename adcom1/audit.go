package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Audit objects represents the outcome of some form of review of the ad (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_audit
type Audit[A any] struct {
	// The audit status of the ad.
	Status AuditStatus `json:"status,omitempty"`
	// One or more human-readable explanations as to reasons for rejection or any changes to fields for ad quality reasons (e.g., adomain, cat, attr, etc.).
	Feedback []string `json:"feedback,omitempty"`
	// Timestamp of the original instantiation of this object in Unix format (i.e., milliseconds since the epoch).
	Init int64 `json:"init,omitempty"`
	// Timestamp of most recent modification to this object in Unix format (i.e., milliseconds since the epoch).
	LastMod int64 `json:"lastmod,omitempty"`
	// Correction object wherein the auditor can specify changes to attributes of the Ad object or its children they believe to be proper.
	Corr *Ad[A] `json:"corr,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
