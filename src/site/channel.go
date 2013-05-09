package site

type Channel struct {
	Sitemap Sitemap
	Variables map[string] string
	Components map[string] Component
	Templates map[string] Template
}
