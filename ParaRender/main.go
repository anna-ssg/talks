func (e *Engine) RenderUserDefinedPages(fileOutPath string, templates *template.Template) {
	numTemplates := len(e.DeepDataMerge.Templates)
	concurrency := numCPU * 2 // Simple way to keep limit on concurrency based on System Resources
	if numTemplates < concurrency {
		concurrency = numTemplates
	}
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)
	for _, templateURL := range templateURLs {
		if templateURL == ".html" {
			continue
		}
		wg.Add(1)
		semaphore <- struct{}{}
		go func(templateURL string) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			e.RenderPage(fileOutPath, template.URL(templateURL), templates,
				e.DeepDataMerge.Templates[template.URL(templateURL)].Frontmatter.Layout)
		}(templateURL)
	}
	wg.Wait()
}
