buffer.WriteString("<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\"?>\n")
buffer.WriteString("<?xml-stylesheet href=\"/static/styles/feed.xsl\" type=\"text/xsl\"?>\n")
buffer.WriteString("<rss version=\"2.0\" xmlns:atom=\"http://www.w3.org/2005/Atom\">\n")
buffer.WriteString("  <channel>\n")
buffer.WriteString("   <title>")
xml.EscapeText(&buffer, []byte(e.DeepDataMerge.LayoutConfig.SiteTitle))
buffer.WriteString("</title>\n")
/*
   so on
*/
buffer.WriteString("   <lastBuildDate>" + time.Now().Format(time.RFC1123Z) + "</lastBuildDate>\n")
buffer.WriteString("   <atom:link href=\"" + e.DeepDataMerge.LayoutConfig.BaseURL + "/feed.xml\" rel=\"self\" type=\"application/rss+xml\" />\n")
