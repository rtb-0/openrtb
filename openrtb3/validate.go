package openrtb3

import (
	"errors"
	"strconv"

	"github.com/rtb-0/openrtb/errs"
)

var (
	// ErrRequired is a missing required field.
	ErrRequired = errs.Err("required")
	// ErrSpecMissing means an item has no placement spec.
	ErrSpecMissing = errs.Err("spec is required")
	// ErrNegative means a price is below zero.
	ErrNegative = errs.Err("must not be negative")
)

// Option changes Validate.
type Option func(*validateConfig)

type validateConfig struct {
	allowMissingFormat bool
}

// AllowMissingFormat accepts items whose spec has no placement subtype.
func AllowMissingFormat() Option {
	return func(c *validateConfig) { c.allowMissingFormat = true }
}

// Validate checks the required OpenRTB 3.0 request fields.
func (r *Request) Validate(opts ...Option) error {
	if r == nil {
		return ErrRequired.WithMessage("Request")
	}
	cfg := validateConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}
	var errs []error
	if r.ID == "" {
		errs = append(errs, ErrRequired.WithMessage("id"))
	}
	if len(r.Item) == 0 {
		errs = append(errs, ErrRequired.WithMessage("item"))
	}
	for i := range r.Item {
		item := &r.Item[i]
		base := "item[" + strconv.Itoa(i) + "]"
		if item.ID == "" {
			errs = append(errs, ErrRequired.WithMessage(base+".id"))
		}
		if item.Spec == nil && !cfg.allowMissingFormat && !item.Popunder {
			errs = append(errs, ErrSpecMissing.WithMessage(base+".spec"))
		}
	}
	return errors.Join(errs...)
}

// Validate checks required bid fields when a seat bid is present.
func (r *Response[A]) Validate() error {
	if r == nil {
		return ErrRequired.WithMessage("Response")
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
			if bid.Item == "" {
				errs = append(errs, ErrRequired.WithMessage(base+".item"))
			}
			if bid.Price < 0 {
				errs = append(errs, ErrNegative.WithMessage(base+".price"))
			}
		}
	}
	return errors.Join(errs...)
}
