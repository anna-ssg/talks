			collectionKey += ".html"
			var found bool
			for _, map_page := range p.CollectionsMap[template.URL(collectionKey)] {
				if map_page.CompleteURL == page.CompleteURL {
					found = true
				}
			}
			if !found {
				p.CollectionsMap[template.URL(collectionKey)] = append(p.CollectionsMap[template.URL(collectionKey)], page)
			}
		}
	}
}
