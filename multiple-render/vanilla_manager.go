func (cmd *Cmd) VanillaRenderManager() {
	// ...
	if cmd.RenderSpecificSite == "" {
		for _, path := range annaConfig.SiteDataPaths {
			// call vanilla render on all directories in anna.json
			cmd.VanillaRender(path)
		}
		// If no site has been rendered due to empty "anna.yml", render the default "site/" path
		if !siteRendered {
			cmd.VanillaRender("site/")
		}
	} else {
		siteRendered := false
		for _, sitePath := range annaConfig.SiteDataPaths {
			// call vanilla on passed directory render the path is valid
		}
	}
}
