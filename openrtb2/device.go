package openrtb2

import (
	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
)

// This object provides information pertaining to the device through which the user is interacting (OpenRTB 2.6 §3.2.18).
// https://github.com/InteractiveAdvertisingBureau/openrtb2.x/blob/main/2.6.md#objectdevice
type Device struct {
	// Location of the device assumed to be the user’s current location defined by a Geo object (Section 3.2.19).
	Geo *Geo `json:"geo,omitempty"`
	// Standard “Do Not Track” flag as set in the header by the browser, where 0 = tracking is unrestricted, 1 = do not track.
	DNT *int8 `json:"dnt,omitempty"`
	// “Limit Ad Tracking” signal commercially endorsed (e.g., iOS, Android), where 0 = tracking is unrestricted, 1 = tracking must be limited per commercial guidelines.
	Lmt *int8 `json:"lmt,omitempty"`
	// Browser user agent string.
	UA string `json:"ua,omitempty"`
	// Structured user agent information defined by a UserAgent object (see Section 3.2.29).
	SUA *UserAgent `json:"sua,omitempty"`
	// IPv4 address closest to device.
	IP string `json:"ip,omitempty"`
	// IP address closest to device as IPv6.
	IPv6 string `json:"ipv6,omitempty"`
	// The general type of device.
	DeviceType adcom1.DeviceType `json:"devicetype,omitempty"`
	// Device make (e.g., “Apple”).
	Make string `json:"make,omitempty"`
	// Device model (e.g., “iPhone”).
	Model string `json:"model,omitempty"`
	// Device operating system (e.g., “iOS”).
	OS string `json:"os,omitempty"`
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
	JS *int8 `json:"js,omitempty"`
	// Indicates if the geolocation API will be available to JavaScript code running in the banner, where 0 = no, 1 = yes.
	GeoFetch *int8 `json:"geofetch,omitempty"`
	// Version of Flash supported by the browser.
	FlashVer string `json:"flashver,omitempty"`
	// Browser language using ISO-639-1-alpha-2.
	Language string `json:"language,omitempty"`
	// Content language using IETF BCP 47.
	LangB string `json:"langb,omitempty"`
	// Carrier or ISP (e.g., “VERIZON”) using exchange curated string names which should be published to bidders a priori.
	Carrier string `json:"carrier,omitempty"`
	// Mobile carrier as the concatenated MCC-MNC code (e.g., "310-005" identifies Verizon Wireless CDMA in the USA).
	MCCMNC string `json:"mccmnc,omitempty"`
	// Network connection type.
	ConnectionType *adcom1.ConnectionType `json:"connectiontype,omitempty"`
	// ID sanctioned for advertiser use in the clear (i.e., not hashed).
	IFA string `json:"ifa,omitempty"`
	// Hardware device ID (e.g., IMEI); hashed via SHA1.
	DIDSHA1 string `json:"didsha1,omitempty"`
	// Hardware device ID (e.g., IMEI); hashed via MD5.
	DIDMD5 string `json:"didmd5,omitempty"`
	// Platform device ID (e.g., Android ID); hashed via SHA1.
	DPIDSHA1 string `json:"dpidsha1,omitempty"`
	// Platform device ID (e.g., Android ID); hashed via MD5.
	DPIDMD5 string `json:"dpidmd5,omitempty"`
	// MAC address of the device; hashed via SHA1.
	MACSHA1 string `json:"macsha1,omitempty"`
	// MAC address of the device; hashed via MD5.
	MACMD5 string `json:"macmd5,omitempty"`
	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext common.Ext `json:"ext,omitempty"`
}
