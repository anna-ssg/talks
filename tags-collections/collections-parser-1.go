func (p *Parser) collectionsParser(page TemplateData) {
	// Iterating over all sets of collections defined in the frontmatter
	for _, collectionSet := range page.Frontmatter.Collections {
		var collections []string
		// Collections will be nested using > as the separator - "posts>tech>Go"
		for _, item := range strings.Split(collectionSet, ">") {
			collections = append(collections, strings.TrimSpace(item))
		}
		for i := range len(collections) {
			collectionKey := "collections/"
			for j := range i + 1 {
				collectionKey += collections[j]
				if j != i {
					collectionKey += "/"
				}
			}
