package adcom1

import (
	"encoding/json"

	"github.com/rtb-0/openrtb/common"
)

// UnmarshalJSON accepts an AdCOM ad object or a JSON string containing that object.
// OpenRTB stores adm as a string; OpenRTB 3 media.ad is an object.
func (a *Ad[A]) UnmarshalJSON(data []byte) error {
	raw, err := common.UnwrapJSONString(data)
	if err != nil {
		return err
	}
	type alias Ad[A]
	var v alias
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*a = Ad[A](v)
	return nil
}
