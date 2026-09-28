package openrtb3

import "github.com/rtb-0/openrtb/adcom1"

// NewRequest starts an OpenRTB 3.0 request.
func NewRequest(id string) *Request {
	return &Request{ID: id}
}

// Auction sets the auction type.
func (r *Request) Auction(at AuctionType) *Request {
	r.AT = at
	return r
}

// AddItem appends an item for sale.
func (r *Request) AddItem(item *Item) *Request {
	if item != nil {
		r.Item = append(r.Item, *item)
	}
	return r
}

// BannerItem builds an item whose spec is a display placement of the given size.
func BannerItem(id string, w, h int64) *Item {
	return &Item{
		ID: id,
		Spec: &adcom1.ItemSpec{Placement: &adcom1.Placement{
			Display: &adcom1.DisplayPlacement{W: w, H: h},
		}},
	}
}

// VideoItem builds an item whose spec is a video placement.
func VideoItem(id string, mime ...string) *Item {
	return &Item{
		ID: id,
		Spec: &adcom1.ItemSpec{Placement: &adcom1.Placement{
			Video: &adcom1.VideoPlacement{MIME: mime},
		}},
	}
}

// PopunderItem builds an item with no placement subtype.
// Validate accepts the missing spec subtype because Popunder is set.
func PopunderItem(id string) *Item {
	return &Item{ID: id, Popunder: true, Spec: &adcom1.ItemSpec{}}
}

// Floor sets the item floor and its currency.
func (item *Item) Floor(flr float64, cur string) *Item {
	item.Flr = flr
	item.FlrCur = cur
	return item
}
