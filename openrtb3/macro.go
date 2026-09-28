package openrtb3

import (
	"github.com/rtb-0/openrtb/common"
)

// Macro object constitutes a buyer defined key/value pair used to inject dynamic values into media markup While they apply to any media markup irrespective of how it is conveyed, the principle use case is for media that was uploaded to the exchange prior to the transaction (e.g., pre-registered for creative quality review) and referenced in bid The full form of the macro to be substituted at runtime is ${CUSTOM_KEY}, where “KEY” is the name supplied in the key attribute This ensures no conflict with standard OpenRTB macros (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_macro
type Macro struct {
	// Name of a buyer specific macro.
	Key string `json:"key"`
	// Value to substitute for each instance of the macro found in markup.
	Value string `json:"value,omitempty"`
	// Optional demand source specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
