package adcom1

// PodDedupe Types of pod deduplication settings (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_poddedupe
type PodDedupe int8

const (
	PodDedupeADomain      PodDedupe = 1
	PodDedupeIABCategory  PodDedupe = 2
	PodDedupeCreativeID   PodDedupe = 3
	PodDedupeMediafileURL PodDedupe = 4
)
