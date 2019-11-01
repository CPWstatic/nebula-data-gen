package generator

import (
	"encoding/xml"
	"io/ioutil"
)

type Osm struct {
	XmlName     xml.Name `xml:"osm"`
	Generator   string   `xml:"generator,attr"`
	Copyright   string   `xml:"copyright, attr"`
	Attribution string   `xml:"attribution, attr"`
	License     string   `xml:"license, attr"`
	Nodes       []Node   `xml:"node"`
}

type Node struct {
	Id        string   `xml:"id,attr"`
	Visible   string   `xml:"visible,attr"`
	Version   string   `xml:"version,attr"`
	Changeset string   `xml:"changeset,attr"`
	Timestamp string   `xml:"timestamp,attr"`
	User      string   `xml:"user,attr"`
	Uid       string   `xml:"uid,attr"`
	Lat       string   `xml:"lat,attr"`
	Lon       string   `xml:"lon,attr"`
}

func ReadNode(cnt int64) ([]Node, error) {
	var nodes []Node
	var osm Osm
	if content, err := ioutil.ReadFile("./generator/map.osm"); err == nil {
		if err = xml.Unmarshal([]byte(content), &osm); err == nil {
			length := len(osm.Nodes)
			var interval int = int(int64(length) / cnt)
			for i := interval; i < length; i += interval {
				node := osm.Nodes[i]
				nodes = append(nodes, node)
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}

	return nodes, nil
}
