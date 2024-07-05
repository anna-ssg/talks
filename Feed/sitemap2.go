for _, templateData := range e.DeepDataMerge.Templates {
	url := e.DeepDataMerge.LayoutConfig.BaseURL + "/" + string(templateData.CompleteURL)
	buffer.WriteString("\t<url>\n")
	buffer.WriteString("\t\t<loc>" + url + "</loc>\n")
	buffer.WriteString("\t\t<lastmod>" + templateData.Frontmatter.Date + "</lastmod>\n")
	buffer.WriteString("\t</url>\n")
}
buffer.WriteString("</urlset>\n")
