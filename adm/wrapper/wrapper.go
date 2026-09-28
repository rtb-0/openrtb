// Package wrapper detects which library markup type is stored in an OpenRTB adm string.
// Raw always keeps the original text so the value can be written back unchanged.
package wrapper

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/rtb-0/openrtb/adcom1"
	"github.com/rtb-0/openrtb/adm/popunder"
	"github.com/rtb-0/openrtb/adm/vast"
	"github.com/rtb-0/openrtb/common"
	nresp "github.com/rtb-0/openrtb/native1/response"
)

// Kind is the markup detected in an adm string.
type Kind uint8

const (
	KindRaw Kind = iota
	KindNative
	KindPopunder
	KindVAST
	KindAdCOM
)

// Markup is one adm value. At most one typed field is set.
type Markup struct {
	Kind     Kind
	Raw      string
	Native   *nresp.Response
	Popunder *popunder.Ad
	VAST     *vast.Document
	AdCOM    *adcom1.Ad[string]
}

// UnmarshalJSON classifies a JSON string stored in adm.
func (m *Markup) UnmarshalJSON(data []byte) error {
	raw, err := common.UnwrapJSONString(data)
	if err != nil {
		return err
	}
	text := string(bytes.TrimSpace(raw))
	*m = Markup{Raw: text}
	if text == "" {
		return nil
	}
	switch text[0] {
	case '{':
		if n, ok := tryNative(text); ok {
			m.Kind = KindNative
			m.Native = n
			return nil
		}
		if ad, ok := tryAdCOM(text); ok {
			m.Kind = KindAdCOM
			m.AdCOM = ad
			return nil
		}
	case '<':
		if strings.Contains(text, "<VAST") {
			var doc vast.Document
			if err := json.Unmarshal(data, &doc); err != nil {
				return err
			}
			m.Kind = KindVAST
			m.VAST = &doc
			return nil
		}
		if strings.Contains(text, "popunderAd") {
			var ad popunder.Ad
			if err := json.Unmarshal(data, &ad); err != nil {
				return err
			}
			m.Kind = KindPopunder
			m.Popunder = &ad
			return nil
		}
	}
	if strings.HasPrefix(text, "http://") || strings.HasPrefix(text, "https://") {
		ad := popunder.URLAd(text)
		m.Kind = KindPopunder
		m.Popunder = &ad
		return nil
	}
	return nil
}

// MarshalJSON writes the original adm text as a JSON string.
func (m Markup) MarshalJSON() ([]byte, error) {
	return common.MarshalString(m.Raw)
}

func tryNative(text string) (*nresp.Response, bool) {
	if !strings.Contains(text, `"assets"`) && !strings.Contains(text, `"link"`) && !strings.Contains(text, `"native"`) {
		return nil, false
	}
	var n nresp.Response
	if err := json.Unmarshal([]byte(text), &n); err != nil {
		return nil, false
	}
	if n.Link.URL == "" && len(n.Assets) == 0 {
		return nil, false
	}
	return &n, true
}

func tryAdCOM(text string) (*adcom1.Ad[string], bool) {
	if !strings.Contains(text, `"display"`) && !strings.Contains(text, `"video"`) && !strings.Contains(text, `"audio"`) {
		return nil, false
	}
	var ad adcom1.Ad[string]
	if err := json.Unmarshal([]byte(text), &ad); err != nil {
		return nil, false
	}
	if ad.Display == nil && ad.Video == nil && ad.Audio == nil {
		return nil, false
	}
	return &ad, true
}
