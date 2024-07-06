// slice it
var posts []parser.TemplateData
for _, templateData := range e.DeepDataMerge.Templates {
	if !templateData.Frontmatter.Draft {
		posts = append(posts, templateData)
	}
}

// sort by publication date
sort.Slice(posts, func(i, j int) bool {
	return posts[i].Date > posts[j].Date // assuming Date is Unix timestamp
})
