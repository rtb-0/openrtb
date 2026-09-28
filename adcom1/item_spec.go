package adcom1

// ItemSpec represents a bunch of data for OpenRTB 3 Item.spec field (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_item
type ItemSpec struct {
	Placement *Placement `json:"placement,omitempty"`
}
