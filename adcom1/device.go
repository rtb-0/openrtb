package adcom1

import (
	"github.com/rtb-0/openrtb/common"
)

// Device object provides information pertaining to the device through which the user is interacting (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#object_device
type Device struct {
	// The general type of device.
	Type DeviceType `json:"type,omitempty"`
	// Browser user agent string.
	UA string `json:"ua,omitempty"`
	// Structured user agent information defined by a Object: UserAgent.
	SUA *UserAgent `json:"sua,omitempty"`
	// ID sanctioned for advertiser use in the clear (i.e., not hashed).
	IFA string `json:"ifa,omitempty"`
	// Standard “Do Not Track” flag as set in the header by the browser, where 0 = tracking is unrestricted, 1 = do not track.
	DNT int8 `json:"dnt,omitempty"`
	// “Limit Ad Tracking” signal commercially endorsed (e.g., iOS, Android), where 0 = tracking is unrestricted, 1 = tracking must be limited per commercial guidelines.
	Lmt int8 `json:"lmt,omitempty"`
	// Device make (e.g., "Apple").
	Make string `json:"make,omitempty"`
	// Device model (e.g., “iPhone10,1” when the specific device model is known, “iPhone” otherwise).
	Model string `json:"model,omitempty"`
	// Device operating system.
	OS OperatingSystem `json:"os,omitempty"`
	// Device operating system version (e.g., “3.1.2”).
	OSV string `json:"osv,omitempty"`
	// Hardware version of the device (e.g., “5S” for iPhone 5S).
	HWV string `json:"hwv,omitempty"`
	// Physical height of the screen in pixels.
	H int64 `json:"h,omitempty"`
	// Physical width of the screen in pixels.
	W int64 `json:"w,omitempty"`
	// Screen size as pixels per linear inch.
	PPI int64 `json:"ppi,omitempty"`
	// The ratio of physical pixels to device independent pixels.
	PxRatio float64 `json:"pxratio,omitempty"`
	// Support for JavaScript, where 0 = no, 1 = yes.
	JS int8 `json:"js,omitempty"`
	// Browser language using ISO-639-1-alpha-2.
	Lang string `json:"lang,omitempty"`
	// Browser language using IETF BCP 47.
	LangB string `json:"langb,omitempty"`
	// IPv4 address closest to device.
	IP string `json:"ip,omitempty"`
	// IP address closest to device as IPv6.
	IPv6 string `json:"ipv6,omitempty"`
	// The value of the “x-forwarded-for” header.
	XFF string `json:"xff,omitempty"`
	// Indicator of truncation of any of the IP attributes (i.e., ip, ipv6, xff), where 0 = no, 1 = yes (e.g., from 1.2.3.4 to 1.2.3.0).
	IPTr int8 `json:"iptr,omitempty"`
	// Carrier or ISP (e.g., “VERIZON”) using exchange curated string names which should be published to bidders a priori.
	Carrier string `json:"carrier,omitempty"`
	// Mobile carrier as the concatenated MCC-MNC code (e.g., “310-005” identifies Verizon Wireless CDMA in the USA).
	MCCMNC string `json:"mccmnc,omitempty"`
	// MCC and MNC of the SIM card using the same format as mccmnc.
	MCCMNCSIM string `json:"mccmncsim,omitempty"`
	// Network connection type.
	ConType ConnectionType `json:"contype,omitempty"`
	// Indicates if the geolocation API will be available to JavaScript code running in display ad, where 0 = no, 1 = yes.
	GeoFetch int8 `json:"geofetch,omitempty"`
	// Location of the device (i.e., typically the user's current location).
	Geo *Geo `json:"geo,omitempty"`
	// Optional vendor-specific extensions.
	Ext common.Ext `json:"ext,omitempty"`
}
