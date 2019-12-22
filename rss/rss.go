package rss

import (
	"encoding/xml"
	"io"
	"time"
)

const PubDateFormat = "Mon, 2 Jan 2006 15:04:05 -0700"

type Guid struct {
	XMLName     xml.Name `xml:"guid"`
	IsPermalink bool     `xml:"isPermalink,attr"`
	Id          string   `xml:",chardata"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Guid        Guid     `xml:"guid"`
	Author      string   `xml:"author"`
	PubDate     PubDate  `xml:"pubDate"`
	Comments    string   `xml:"comments"`
	Description string   `xml:"description"`
	Categories  []string `xml:"category"`
}

type PubDate time.Time

func (pd *PubDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(PubDateFormat, v)
	if err != nil {
		return err
	}
	*pd = PubDate(parse)
	return err
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	AtomLink    AtomLink `xml:"atom:link"`
	Items       []Item   `xml:"item"`
}

type AtomLink struct {
	XMLName   xml.Name `xml:"atom:link"`
	Reference string   `xml:"href,attr"`
	Relation  string   `xml:"rel,attr"`
	Type      string   `xml:"type,attr"`
}

type Rss struct {
	XMLName   xml.Name `xml:"rss"`
	Version   string   `xml:"version,attr"`
	Namespace string   `xml:"xmlns:atom,attr"`
	Channel   Channel  `xml:"channel"`
}

func DecodeItems(reader io.Reader) ([]Item, error) {
	decoder := xml.NewDecoder(reader)
	feed := Rss{}
	err := decoder.Decode(&feed)
	if err != nil {
		return nil, err
	}
	return feed.Channel.Items, nil
}

func PubDateToTime(pd PubDate) time.Time {
	return time.Time(pd)
}
