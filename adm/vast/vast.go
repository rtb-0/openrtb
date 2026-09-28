// Package vast reads the fields needed to recognize a VAST document stored in adm.
// The original XML is kept in Raw. Unknown elements stay there and are not rewritten.
package vast

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"

	"github.com/rtb-0/openrtb/common"
)

// Document is a VAST payload carried in an OpenRTB adm string.
type Document struct {
	Raw     string
	Version string
	Ads     []Ad
}

// Ad is one VAST Ad element, either an InLine or a Wrapper.
type Ad struct {
	ID          string
	Inline      bool
	Wrapper     bool
	AdTagURI    string
	Impressions []string
	MediaFiles  []string
}

type xmlDoc struct {
	XMLName xml.Name `xml:"VAST"`
	Version string   `xml:"version,attr"`
	Ads     []xmlAd  `xml:"Ad"`
}

type xmlAd struct {
	ID      string   `xml:"id,attr"`
	InLine  *xmlBody `xml:"InLine"`
	Wrapper *xmlBody `xml:"Wrapper"`
}

type xmlBody struct {
	AdTagURI    string   `xml:"VASTAdTagURI"`
	Impressions []string `xml:"Impression"`
	Creatives   struct {
		Creative []struct {
			Linear struct {
				MediaFiles struct {
					MediaFile []string `xml:"MediaFile"`
				} `xml:"MediaFiles"`
			} `xml:"Linear"`
		} `xml:"Creative"`
	} `xml:"Creatives"`
}

// UnmarshalJSON reads a JSON string that holds a VAST XML document.
func (d *Document) UnmarshalJSON(data []byte) error {
	raw, err := common.UnwrapJSONString(data)
	if err != nil {
		return err
	}
	text := string(bytes.TrimSpace(raw))
	d.Raw = text
	dec := xml.NewDecoder(strings.NewReader(text))
	dec.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) {
		return input, nil
	}
	var doc xmlDoc
	if err := dec.Decode(&doc); err != nil {
		return err
	}
	d.Version = doc.Version
	d.Ads = d.Ads[:0]
	for _, src := range doc.Ads {
		ad := Ad{ID: src.ID}
		body := src.InLine
		if src.InLine != nil {
			ad.Inline = true
		}
		if src.Wrapper != nil {
			ad.Wrapper = true
			body = src.Wrapper
		}
		if body != nil {
			ad.AdTagURI = strings.TrimSpace(body.AdTagURI)
			ad.Impressions = trimAll(body.Impressions)
			for _, creative := range body.Creatives.Creative {
				ad.MediaFiles = append(ad.MediaFiles, trimAll(creative.Linear.MediaFiles.MediaFile)...)
			}
		}
		d.Ads = append(d.Ads, ad)
	}
	return nil
}

// MarshalJSON writes the original VAST text as a JSON string.
func (d Document) MarshalJSON() ([]byte, error) {
	return common.MarshalString(d.Raw)
}

func trimAll(in []string) []string {
	if len(in) == 0 {
		return nil
	}
	out := make([]string, 0, len(in))
	for _, s := range in {
		s = strings.TrimSpace(s)
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
