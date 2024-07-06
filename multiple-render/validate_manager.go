func (cmd *Cmd) ValidateHTMLManager() {
	// Rendering all sites
	cmd.VanillaRenderManager()

	// Check if the configuration file exists
	// If it does not, validate only the site/ directory

	_, err := os.Stat("anna.json")
	if os.IsNotExist(err) {
		cmd.VanillaRender("site/")
		return
	}

	// Read and parse the configuration file
	annaConfigFile, err := os.ReadFile("anna.json")
	if err != nil {
		cmd.ErrorLogger.Fatal(err)
	}

	var annaConfig AnnaConfig

	err = json.Unmarshal(annaConfigFile, &annaConfig)
	if err != nil {
		cmd.ErrorLogger.Fatal(err)
	}

	// Validating sites
	validatedSites := false

	for _, sitePath := range annaConfig.SiteDataPaths {
		cmd.ValidateHTMLContent(sitePath)
		if !validatedSites {
			validatedSites = true
		}
	}

	// If no site has been validated due to empty "anna.yml", validate the default "site/" path
	if !validatedSites {
		cmd.ValidateHTMLContent("site/")
	}

}
