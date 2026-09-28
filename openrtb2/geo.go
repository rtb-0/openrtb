package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object encapsulates various methods for specifying a geographic location (OpenRTB 2.6 §3.2.19).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectgeo
type Geo struct {
	// Latitude from -90.0 to +90.0, where negative is south.
	Lat *float64 `json:"lat,omitempty"`
	// Longitude from -180.0 to +180.0, where negative is west.
	Lon *float64 `json:"lon,omitempty"`
	// Source of location data; recommended when passing lat/lon.
	Type adcom1.LocationType `json:"type,omitempty"`
	// Estimated location accuracy in meters; recommended when lat/lon are specified and derived from a device’s location services (i.e., type = 1).
	Accuracy int64 `json:"accuracy,omitempty"`
	// Number of seconds since this geolocation fix was established.
	LastFix int64 `json:"lastfix,omitempty"`
	// Service or provider used to determine geolocation from IP address if applicable (i.e., type = 2).
	IPService adcom1.IPLocationService `json:"ipservice,omitempty"`
	// Country code using ISO-3166-1-alpha-3.
	Country string `json:"country,omitempty"`
	// Region code using ISO-3166-2; 2-letter state code if USA.
	Region string `json:"region,omitempty"`
	// Region of a country using FIPS 10-4 notation.
	RegionFIPS104 string `json:"regionfips104,omitempty"`
	// Google metro code; similar to but not exactly Nielsen DMAs.
	Metro string `json:"metro,omitempty"`
	// City using United Nations Code for Trade & Transport Locations.
	City string `json:"city,omitempty"`
	// Zip or postal code.
	ZIP string `json:"zip,omitempty"`
	// Local time as the number +/- of minutes from UTC.
	UTCOffset int64 `json:"utcoffset,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
