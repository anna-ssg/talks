func (e *Engine) RenderTags(fileOutPath string, templ *template.Template) {
	var tagsBuffer bytes.Buffer

	// Extracting tag titles
	tags := make([]template.URL, 0, len(e.DeepDataMerge.TagsMap))
	for tag := range e.DeepDataMerge.TagsMap {
		tags = append(tags, tag)
	}

	slices.SortFunc(tags, func(a, b template.URL) int {
		return cmp.Compare(strings.ToLower(string(a)), strings.ToLower(string(b)))
	})

	tagNames := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagString := string(tag)
		tagString, _ = strings.CutPrefix(tagString, "tags/")
		tagString, _ = strings.CutSuffix(tagString, ".html")

		tagNames = append(tagNames, tagString)
	}

	tagRootTemplataData := parser.TemplateData{
		Frontmatter: parser.Frontmatter{Title: "Tags"},
	}

	tagTemplateData := TagRootTemplateData{
		DeepDataMerge: e.DeepDataMerge,
		PageURL:       "tags.html",
		TemplateData:  tagRootTemplataData,
		TagNames:      tagNames,
	}

	// Rendering the page displaying all tags
	err := templ.ExecuteTemplate(&tagsBuffer, "all-tags", tagTemplateData)
	if err != nil {
		e.ErrorLogger.Fatal(err)
	}

	// Flushing 'tags.html' to the disk
	err = os.WriteFile(fileOutPath+"rendered/tags.html", tagsBuffer.Bytes(), 0666)
	if err != nil {
		e.ErrorLogger.Fatal(err)
	}

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	e.DeepDataMerge.Tags = make(map[template.URL]parser.TemplateData)

	for tag := range e.DeepDataMerge.TagsMap {
		slices.SortFunc(e.DeepDataMerge.TagsMap[tag], func(a, b parser.TemplateData) int {
			return cmp.Compare(b.Date, a.Date)
		})
		tagString := string(tag)
		tagString, _ = strings.CutPrefix(tagString, "tags/")
		tagString, _ = strings.CutSuffix(tagString, ".html")

		e.DeepDataMerge.Tags[tag] = parser.TemplateData{
			Frontmatter: parser.Frontmatter{
				Title: tagString,
			},
		}
	}

	// Rendering the subpages with merged tagged posts
	// for tag, taggedTemplates := range e.DeepDataMerge.TagsMap {
	// 	wg.Add(1)
	// 	go func(tag template.URL, taggedTemplates []parser.TemplateData) {
	// 		defer wg.Done()

	// 		e.RenderPage(fileOutPath, tag, templ, "tag-subpage")
	// 	}(tag, taggedTemplates)
	// }

	// Wait for all goroutines to finish
	wg.Wait()
}
