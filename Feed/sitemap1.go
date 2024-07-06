buffer.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
buffer.WriteString("<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n")

// Sorting templates by key
keys := make([]string, 0, len(e.DeepDataMerge.Templates))
for k := range e.DeepDataMerge.Templates {
	keys = append(keys, string(k))
}
sort.Strings(keys)
tempTemplates := make(map[template.URL]parser.TemplateData)
for _, templateURL := range keys {
	tempTemplates[template.URL(templateURL)] = e.DeepDataMerge.Templates[template.URL(templateURL)]
}