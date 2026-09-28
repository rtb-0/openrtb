package openrtb2

import (
	"errors"
	"strconv"

	"github.com/rtb-0/openrtb/errs"
)

var (
	// ErrRequired is a missing required field.
	ErrRequired = errs.Err("required")
	// ErrFormatMissing means an impression has no banner, video, audio, or native object.
	ErrFormatMissing = errs.Err("banner, video, audio, or native is required")
	// ErrSiteApp means site and app were both set.
	ErrSiteApp = errs.Err("site and app are mutually exclusive")
	// ErrMIMEs means a video or audio impression has no MIME types.
	ErrMIMEs = errs.Err("mimes is required")
	// ErrNativeRequest means a native impression has an empty request.
	ErrNativeRequest = errs.Err("request is required")
	// ErrNegative means a price is below zero.
	ErrNegative = errs.Err("must not be negative")
)

// Option changes Validate.
type Option func(*validateConfig)

type validateConfig struct {
	allowMissingFormat bool
}

// AllowMissingFormat accepts impressions that have no banner, video, audio, or native object.
// Popunder requests from the wire use this. PopunderImp sets the same exception on the impression itself.
func AllowMissingFormat() Option {
	return func(c *validateConfig) { c.allowMissingFormat = true }
}

// Validate checks the required OpenRTB 2.6 request fields.
func (r *BidRequest) Validate(opts ...Option) error {
	if r == nil {
		return ErrRequired.WithMessage("BidRequest")
	}
	cfg := validateConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}
	var errs []error
	if r.ID == "" {
		errs = append(errs, ErrRequired.WithMessage("id"))
	}
	if r.Site != nil && r.App != nil {
		errs = append(errs, ErrSiteApp.WithMessage("site"))
	}
	if len(r.Imp) == 0 {
		errs = append(errs, ErrRequired.WithMessage("imp"))
	}
	for i := range r.Imp {
		errs = append(errs, r.Imp[i].validate("imp["+strconv.Itoa(i)+"]", cfg.allowMissingFormat)...)
	}
	return errors.Join(errs...)
}

func (imp *Imp) validate(path string, allowMissing bool) []error {
	var errs []error
	if imp.ID == "" {
		errs = append(errs, ErrRequired.WithMessage(path+".id"))
	}
	formats := 0
	if imp.Banner != nil {
		formats++
	}
	if imp.Video != nil {
		formats++
		if len(imp.Video.MIMEs) == 0 {
			errs = append(errs, ErrMIMEs.WithMessage(path+".video.mimes"))
		}
	}
	if imp.Audio != nil {
		formats++
		if len(imp.Audio.MIMEs) == 0 {
			errs = append(errs, ErrMIMEs.WithMessage(path+".audio.mimes"))
		}
	}
	if imp.Native != nil {
		formats++
		if imp.Native.Request.IsZero() {
			errs = append(errs, ErrNativeRequest.WithMessage(path+".native.request"))
		}
	}
	if formats == 0 && !allowMissing && !imp.Popunder {
		errs = append(errs, ErrFormatMissing.WithMessage(path))
	}
	return errs
}

// Validate checks required bid fields when a seat bid is present.
// A zero price is accepted: the field cannot tell a missing number from 0.
func (r *BidResponse[A]) Validate() error {
	if r == nil {
		return ErrRequired.WithMessage("BidResponse")
	}
	var errs []error
	if r.ID == "" {
		errs = append(errs, ErrRequired.WithMessage("id"))
	}
	for i := range r.SeatBid {
		seat := "seatbid[" + strconv.Itoa(i) + "]"
		if len(r.SeatBid[i].Bid) == 0 {
			errs = append(errs, ErrRequired.WithMessage(seat+".bid"))
		}
		for j := range r.SeatBid[i].Bid {
			bid := r.SeatBid[i].Bid[j]
			base := seat + ".bid[" + strconv.Itoa(j) + "]"
			if bid.ID == "" {
				errs = append(errs, ErrRequired.WithMessage(base+".id"))
			}
			if bid.ImpID == "" {
				errs = append(errs, ErrRequired.WithMessage(base+".impid"))
			}
			if bid.Price < 0 {
				errs = append(errs, ErrNegative.WithMessage(base+".price"))
			}
		}
	}
	return errors.Join(errs...)
}
