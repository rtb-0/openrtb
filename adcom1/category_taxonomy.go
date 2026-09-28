package adcom1

// CategoryTaxonomy identifies the taxonomy in effect when content categories are listed (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_categorytaxonomies
type CategoryTaxonomy int16

// CategoryTaxonomy options.
const (
	CatTaxIABContent10  CategoryTaxonomy = 1 // IAB Tech Lab Content Category Taxonomy 1.0: Deprecated, and recommend NOT be used since it does not have SCD flags.
	CatTaxIABContent20  CategoryTaxonomy = 2 // IAB Tech Lab Content Category Taxonomy 2.0: Deprecated, and recommend NOT be used since it does not have SCD flags.
	CatTaxIABProduct10  CategoryTaxonomy = 3 // IAB Tech Lab Ad Product Taxonomy 1.0.
	CatTaxIABAudience11 CategoryTaxonomy = 4 // IAB Tech Lab Audience Taxonomy 1.1.
	CatTaxIABContent21  CategoryTaxonomy = 5 // IAB Tech Lab Content Category Taxonomy 2.1.
	CatTaxIABContent22  CategoryTaxonomy = 6 // IAB Tech Lab Content Category Taxonomy 2.2.
	CatTaxIABContent30  CategoryTaxonomy = 7 // IAB Tech Lab Content Category Taxonomy 3.0.
	CatTaxIABProduct20  CategoryTaxonomy = 8 // IAB Tech Lab Ad Product Taxonomy 2.0.
	CatTaxIABContent31  CategoryTaxonomy = 9 // IAB Tech Lab Content Taxonomy 3.1.
)
