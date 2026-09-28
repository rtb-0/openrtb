package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Geo object encapsulates various methods for specifying a geographic location (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_geo
type Geo struct {
	// Source of location data; recommended when passing lat/lon.
	Type LocationType `json:"type,omitempty"`
	// Latitude from -90.0 to +90.0, where negative is south.
	Lat float64 `json:"lat,omitempty"`
	// Longitude from -180.0 to +180.0, where negative is west.
	Lon float64 `json:"lon,omitempty"`
	// Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device's location services (i.e., type = 1).
	Accur int64 `json:"accur,omitempty"`
	// Number of seconds since this geolocation fix was established.
	LastFix int64 `json:"lastfix,omitempty"`
	// Service or provider used to determine geolocation from IP address if applicable (i.e., type = 2).
	IPServ IPLocationService `json:"ipserv,omitempty"`
	// Country code using ISO-3166-1-alpha-2.
	Country string `json:"country,omitempty"`
	// Region code using ISO-3166-2; 2-letter state code if USA.
	Region string `json:"region,omitempty"`
	// Regional marketing areas such as Nielsen's DMA codes or other similar taxonomy to be agreed among vendors prior to use.
	Metro string `json:"metro,omitempty"`
	// City using United Nations Code for Trade & Transport Locations “UN/LOCODE” with the space between country and city suppressed (e.g., Boston MA, USA = “USBOS”).
	City string `json:"city,omitempty"`
	// ZIP or postal code.
	ZIP string `json:"zip,omitempty"`
	// Local time as the number +/- of minutes from UTC.
	UTCOffset int64 `json:"utcoffset,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
