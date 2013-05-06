package site

type Template struct {
	Filename string
}

type Page struct {
	Template  Template
	Component Component
}

type Component struct {
	ObjectName string
}

type Sitemap struct {
	Mapping map[string] Page
}

type Channel struct {
	Hostname string
	Port     int32
	Prefix   string
	Sitemap  Sitemap
}
