package adcom1

// DistributionChannel is an abstraction of the various types of entities or channels through which ads are distributed (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#abstract_distributionchannel
type DistributionChannel struct {
	// Vendor-specific unique identifier of the distribution channel.
	ID string `json:"id,omitempty"`
	// Displayable name of the distribution channel.
	Name string `json:"name,omitempty"`
	// Details about the publisher of the distribution channel.
	Pub *Publisher `json:"pub,omitempty"`
	// Details about the content within the distribution channel.
	Content *Content `json:"content,omitempty"`
}
