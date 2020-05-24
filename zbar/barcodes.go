package zbar

type Barcodes struct {
	Sources []Source `xml:"source"`
}

type Source struct {
	Href    string  `xml:"href,attr"`
	Indexes []Index `xml:"index"`
}

type Index struct {
	Num     int8     `xml:"num,attr"`
	Symbols []Symbol `xml:"symbol"`
}

type Symbol struct {
	Type    string `xml:"type,attr"`
	Quality string `xml:"quality,attr"`
	Data    string `xml:"data"`
}

func (b *Barcodes) Symbols() map[string][]Symbol {
	texts := make(map[string][]Symbol)
	for _, source := range b.Sources {
		if len(source.Href) == 0 {
			continue
		}
		for _, index := range source.Indexes {
			for _, symbol := range index.Symbols {
				if _, ok := texts[source.Href]; !ok {
					texts[source.Href] = make([]Symbol, 0)
				}
				texts[source.Href] = append(texts[source.Href], symbol)
			}
		}
	}

	return texts
}
