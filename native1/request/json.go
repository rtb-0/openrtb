package request

import (
	"encoding/json"

	"github.com/rtb-0/openrtb/common"
)

// UnmarshalJSON accepts a native request object, a JSON string containing that object,
// or a Native 1.0 document wrapped in {"native": ...}.
// https://github.com/InteractiveAdvertisingBureau/Native-Ads/blob/main/OpenRTB-Native-Ads-Specification-Final-1.2.md
func (r *Request) UnmarshalJSON(data []byte) error {
	raw, err := common.UnwrapNativeRoot(data)
	if err != nil {
		return err
	}
	type alias Request
	var v alias
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*r = Request(v)
	return nil
}
