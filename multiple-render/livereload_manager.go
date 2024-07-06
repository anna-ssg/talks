func (cmd *Cmd) LiveReloadManager() {
	// ...
	if cmd.ServeSpecificSite == "" {
		cmd.StartLiveReload("site/")
	} else {
		for _, sitePath := range annaConfig.SiteDataPaths {
			if strings.Compare(cmd.ServeSpecificSite, sitePath) == 0 {
				cmd.StartLiveReload(sitePath)
				return
			}
		}

		cmd.ErrorLogger.Fatal("Invalid site path to serve")

	}
}
