package rss

import (
	"encoding/xml"
	"time"
)

type RSS struct {
	XMLName        xml.Name `xml:"rss"`
	XMLXmlnsAtom   string   `xml:"xmlns:atom,attr"`
	XMLXmlnsItunes string   `xml:"xmlns:itunes,attr"`
	XMLVersion     string   `xml:"version,attr"`
	Channel        *Channel `xml:"channel,omitempty"`
}

type Channel struct {
	Title          string        `xml:"title,omitempty"`
	Description    string        `xml:"description,omitempty"`
	Link           string        `xml:"link,omitempty"`
	Language       string        `xml:"language,omitempty"`
	Copyright      string        `xml:"copyright,omitempty"`
	ChannelImage   *ChannelImage `xml:"image,omitempty"`
	LastBuildDate  string        `xml:"lastBuildDate,omitempty"`
	AtomLink       *AtomLink     `xml:"atom:link,omitempty"`
	ItunesSubtitle string        `xml:"itunes:subtitle,omitempty"`
	ItunesAuthor   string        `xml:"itunes:author,omitempty"`
	ItunesSummary  string        `xml:"itunes:summary,omitempty"`
	ItunesKeywords string        `xml:"itunes:keywords,omitempty"`
	ItunesExplicit string        `xml:"itunes:explicit,omitempty"`
	ItunesOwner    *ItunesOwner  `xml:"itunes:owner,omitempty"`
	ItunesImage    *ItunesImage
	Item           []*Item `xml:"item,omitempty"`
}

type ItunesOwner struct {
	ItunesName  string `xml:"itunes:name,omitempty"`
	ItunesEmail string `xml:"itunes:mail,omitempty"`
}

type ItunesImage struct {
	XMLName xml.Name `xml:"itunes:image,omitempty"`
	Href    string   `xml:"href,attr"`
}

type ChannelImage struct {
	URL   string `xml:"url,omitempty"`
	Title string `xml:"title,omitempty"`
	Link  string `xml:"link,omitempty"`
}

type AtomLink struct {
	Href string `xml:"href,attr,omitempty"`
	Rel  string `xml:"rel,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
}

type Item struct {
	Title                   string     `xml:"title,omitempty"`
	Description             string     `xml:"description,omitempty"`
	Guid                    string     `xml:"guid,omitempty"`
	PubDate                 string     `xml:"pubDate,omitempty"`
	Enclosure               *Enclosure `xml:"enclosure,omitempty"`
	ItunesAuthor            string     `xml:"itunes:author,omitempty"`
	ItunesSubtitle          string     `xml:"itunes:subtitle,omitempty"`
	ItunesSummary           string     `xml:"itunes:summary,omitempty"`
	ItunesDuration          string     `xml:"itunes:duration,omitempty"`
	ItunesExplicit          string     `xml:"itunes:explicit,omitempty"`
	ItunesOrder             string     `xml:"itunes:order,omitempty"`
	ItunesisClosedCaptioned string     `xml:"itunes:isClosedCaptioned,omitempty"`
	ItunesBlock             string     `xml:"itunes:block,allowempty"`
}

type Enclosure struct {
	URL    string `xml:"url,attr,omitempty"`
	Type   string `xml:"type,attr,omitempty"`
	Length int64  `xml:"length,attr,omitempty"`
}

// Sort Interface

type ByPubDate []*Item

func (p ByPubDate) Len() int      { return len(p) }
func (p ByPubDate) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p ByPubDate) Less(i, j int) bool {
	iDate, _ := time.Parse(time.RFC1123, p[i].PubDate)
	jDate, _ := time.Parse(time.RFC1123, p[j].PubDate)
	return iDate.Unix() < jDate.Unix()
}
