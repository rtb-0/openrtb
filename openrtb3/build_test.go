package openrtb3_test

import (
	"encoding/json"
	"testing"

	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/common"
	"github.com/rtb-0/openrtb/openrtb3"
)

func TestItemBuilder(t *testing.T) {
	req := openrtb3.NewRequest("req").
		Auction(openrtb3.SecondPricePlus).
		AddItem(openrtb3.BannerItem("1", 300, 250).Floor(0.5, "USD")).
		AddItem(openrtb3.VideoItem("2", "video/mp4")).
		AddItem(openrtb3.PopunderItem("3"))
	if err := req.Validate(); err != nil {
		t.Fatal(err)
	}
	if req.Item[0].Spec.Placement.Display.W != 300 || !req.Item[2].Popunder {
		t.Fatalf("%+v", req.Item)
	}

	body := []byte(`{"id":"req","seatbid":[{"bid":[{"item":"1","price":1.2,"media":{"ad":{"id":"cr1","display":{"adm":"<b>hi</b>","w":300,"h":250}}}}]}]}`)
	var resp openrtb3.Response[string]
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatal(err)
	}
	if err := resp.Validate(); err != nil {
		t.Fatal(err)
	}
	adm := resp.SeatBid[0].Bid[0].Media.Ad.Display.Adm.Value
	if adm != "<b>hi</b>" {
		t.Fatalf("adm %q", adm)
	}

	var ad adcom1.Ad[string]
	if err := json.Unmarshal([]byte(`"{\"id\":\"cr1\",\"display\":{\"adm\":\"<b>hi</b>\"}}"`), &ad); err != nil {
		t.Fatal(err)
	}
	if ad.ID != "cr1" || ad.Display.Adm.Value != "<b>hi</b>" {
		t.Fatalf("%+v", ad)
	}
	_ = common.Protocol{}
}
