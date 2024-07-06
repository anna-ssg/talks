func (e *Engine) RenderCollections(fileOutPath string, templ *template.Template) {
	var collectionsBuffer bytes.Buffer

	// Extracting collection titles
	collections := make([]template.URL, 0, len(e.DeepDataMerge.CollectionsMap))
	for collection := range e.DeepDataMerge.CollectionsMap {
		collections = append(collections, collection)
	}

	slices.SortFunc(collections, func(a, b template.URL) int {
		return cmp.Compare(strings.ToLower(string(a)), strings.ToLower(string(b)))
	})

	collectionNames := make([]string, 0, len(collections))
	for _, collection := range collections {
		collectionString := string(collection)
		collectionString, _ = strings.CutPrefix(collectionString, "collections/")
		collectionString, _ = strings.CutSuffix(collectionString, ".html")

		collectionNames = append(collectionNames, string(collectionString))
	}

	collectionRootTemplataData := parser.TemplateData{
		Frontmatter: parser.Frontmatter{Title: "Collections"},
	}

	collectionTemplateData := CollectionRootTemplateData{
		DeepDataMerge:   e.DeepDataMerge,
		PageURL:         "collections.html",
		TemplateData:    collectionRootTemplataData,
		CollectionNames: collectionNames,
	}

	// Rendering the page displaying all collections
	err := templ.ExecuteTemplate(&collectionsBuffer, "all-collections", collectionTemplateData)
	if err != nil {
		e.ErrorLogger.Fatal(err)
	}

	// Flushing 'collections.html' to the disk
	err = os.WriteFile(fileOutPath+"rendered/collections.html", collectionsBuffer.Bytes(), 0666)
	if err != nil {
		e.ErrorLogger.Fatal(err)
	}

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	e.DeepDataMerge.Collections = make(map[template.URL]parser.TemplateData)

	for collection := range e.DeepDataMerge.CollectionsMap {
		slices.SortFunc(e.DeepDataMerge.CollectionsMap[collection], func(a, b parser.TemplateData) int {
			return cmp.Compare(b.Date, a.Date)
		})

		collectionString := string(collection)
		collectionString, _ = strings.CutPrefix(collectionString, "collections/")
		collectionString, _ = strings.CutSuffix(collectionString, ".html")

		e.DeepDataMerge.Collections[collection] = parser.TemplateData{
			Frontmatter: parser.Frontmatter{
				Title: collectionString,
			},
		}
	}

	// Rendering the subpages with merged tagged posts
	for collection, collectionTemplates := range e.DeepDataMerge.CollectionsMap {
		wg.Add(1)
		go func(collection template.URL, collectionTemplates []parser.TemplateData) {
			defer wg.Done()

			layoutName := e.DeepDataMerge.CollectionsSubPageLayouts[collection]
			if layoutName == "" {
				layoutName = "collection-subpage"
			}

			e.RenderPage(fileOutPath, collection, templ, layoutName)
		}(collection, collectionTemplates)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
