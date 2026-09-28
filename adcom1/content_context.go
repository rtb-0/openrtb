package adcom1

// ContentContext represents options for indicating the type of content being used or consumed by the user in which ads may appear (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_contentcontexts
type ContentContext int8

// Options for indicating the type of content being used or consumed by the user in which ads may appear.
const (
	ContentVideo   ContentContext = 1 // 1 Video (i.e., video file or stream such as Internet TV broadcasts)
	ContentGame    ContentContext = 2 // 2 Game (i.e., an interactive software game)
	ContentMusic   ContentContext = 3 // 3 Music (i.e., audio file or stream such as Internet radio broadcasts)
	ContentApp     ContentContext = 4 // 4 Application (i.e., an interactive software application)
	ContentText    ContentContext = 5 // 5 Text (i.e., primarily textual document such as a web page, eBook, or news article)
	ContentOther   ContentContext = 6 // 6 Other (i.e., none of the other categories applies)
	ContentUnknown ContentContext = 7 // 7 Unknown
)
