package openrtb3

// OpenRTB top-level object is the root for both request and response payloads (OpenRTB 3.0).
// https://github.com/InteractiveAdvertisingBureau/openrtb/blob/main/OpenRTB%20v3.0%20FINAL.md#object_openrtb
type OpenRTB[A any] struct {
	// Version of the Layer-3 OpenRTB specification (e.g., "3.0").
	Ver string `json:"ver,omitempty"`
	// Identifier of the Layer-4 domain model used to define items for sale, media associated with bids, etc.
	DomainSpec string `json:"domainspec,omitempty"`
	// Specification version of the Layer-4 domain model referenced in the domainspec attribute.
	DomainVer string `json:"domainver"`
	// Bid request container.
	Request *Request `json:"request,omitempty"`
	// Bid response container.
	Response *Response[A] `json:"response,omitempty"`
}
