type DeepDataMerge struct {
	// Templates stores the template data of all the pages of the site
	// Access the data for a particular page by using the relative path to the file as the key
	Templates map[template.URL]parser.TemplateData
	// Templates stores the template data of all tag sub-pages of the site
	Tags map[template.URL]parser.TemplateData
	// K-V pair storing all templates corresponding to a particular tag in the site
	TagsMap map[template.URL][]parser.TemplateData
	// Stores data parsed from layout/config.yml
	LayoutConfig parser.LayoutConfig
	// Templates stores the template data of all collection sub-pages of the site
	Collections map[template.URL]parser.TemplateData
	// K-V pair storing all templates corresponding to a particular collection in the site
	CollectionsMap map[template.URL][]parser.TemplateData
	// K-V pair storing the template layout name for a particular collection in the site
	CollectionsSubPageLayouts map[template.URL]string
	// Stores the index generated for search functionality
	JSONIndex map[template.URL]JSONIndexTemplate
}
