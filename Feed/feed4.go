for _, templateData := range posts {
	buffer.WriteString("    <item>\n")
	buffer.WriteString("      <title>")
	xml.EscapeText(&buffer, []byte(templateData.Frontmatter.Title))
	buffer.WriteString("</title>\n")
	buffer.WriteString("      <link>" + e.DeepDataMerge.LayoutConfig.BaseURL + "/" + string(templateData.CompleteURL) + "/</link>\n")
	/*
		so on
	*/
	buffer.WriteString("      <description>")
	xml.EscapeText(&buffer, []byte(templateData.Body))
	buffer.WriteString("</description>\n")
	buffer.WriteString("    </item>\n")
}
buffer.WriteString("  </channel>\n")
buffer.WriteString("</rss>\n")
outputFile, err := os.Create(e.SiteDataPath + "rendered/feed.xml")
