package adcom1

// RequestContext represents a bunch of data for OpenRTB 3 Request.context field (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_request
type RequestContext struct {
	Site         *Site         `json:"site,omitempty"`
	App          *App          `json:"app,omitempty"`
	DOOH         *DOOH         `json:"dooh,omitempty"`
	User         *User         `json:"user,omitempty"`
	Device       *Device       `json:"device,omitempty"`
	Regs         *Regs         `json:"regs,omitempty"`
	Restrictions *Restrictions `json:"restrictions,omitempty"`
}
