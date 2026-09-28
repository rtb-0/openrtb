package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
	nreq "github.com/rtb-0/openrtb/native1/request"
)

// NewBidRequest starts a bid request with the given auction id (OpenRTB 2.6 §3.2.1).
func NewBidRequest(id string) *BidRequest {
	return &BidRequest{ID: id}
}

// Auction sets the auction type.
func (r *BidRequest) Auction(at AuctionType) *BidRequest {
	r.AT = at
	return r
}

// SetSite sets the publisher site. Site and App together fail validation.
func (r *BidRequest) SetSite(site Site) *BidRequest {
	r.Site = &site
	return r
}

// SetApp sets the publisher app.
func (r *BidRequest) SetApp(app App) *BidRequest {
	r.App = &app
	return r
}

// SetDevice sets the device.
func (r *BidRequest) SetDevice(device Device) *BidRequest {
	r.Device = &device
	return r
}

// AddImp appends an impression.
func (r *BidRequest) AddImp(imp *Imp) *BidRequest {
	if imp != nil {
		r.Imp = append(r.Imp, *imp)
	}
	return r
}

// BannerImp builds a banner impression. The constructor is not named Banner because that name is the object type.
func BannerImp(id string, w, h int64) *Imp {
	return &Imp{ID: id, Banner: &Banner{W: Int64Ptr(w), H: Int64Ptr(h)}}
}

// VideoImp builds a video impression.
func VideoImp(id string, video Video) *Imp {
	return &Imp{ID: id, Video: &video}
}

// AudioImp builds an audio impression.
func AudioImp(id string, audio Audio) *Imp {
	return &Imp{ID: id, Audio: &audio}
}

// NativeImp builds a native impression. The request is stored as the JSON string OpenRTB expects.
func NativeImp(id string, req nreq.Request) *Imp {
	return &Imp{ID: id, Native: &Native{Request: common.Encode(req)}}
}

// PopunderImp builds an impression with no banner, video, audio, or native object.
// Instl is set, and Validate accepts the missing format.
func PopunderImp(id string) *Imp {
	return &Imp{ID: id, Instl: 1, Popunder: true}
}

// Pos sets the banner or video position.
func (imp *Imp) Pos(pos int8) *Imp {
	p := adcom1.PlacementPosition(pos)
	if imp.Banner != nil {
		imp.Banner.Pos = &p
	}
	if imp.Video != nil {
		imp.Video.Pos = &p
	}
	return imp
}

// SetSecure requires HTTPS markup.
func (imp *Imp) SetSecure() *Imp {
	v := int8(1)
	imp.Secure = &v
	return imp
}

// SetBidFloor sets the impression floor and its currency.
func (imp *Imp) SetBidFloor(floor float64, cur string) *Imp {
	imp.BidFloor = floor
	imp.BidFloorCur = cur
	return imp
}
