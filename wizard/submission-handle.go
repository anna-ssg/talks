func (ws *WizardServer) handleSubmit(w http.ResponseWriter, r *http.Request) {
	// snip

	// got the form data, now ask theme.go to unzip and place in current dir
	var config parser.LayoutConfig
	err := json.NewDecoder(r.Body).Decode(&config)
	// Writing encoded json to FS to replace the current config.json

	err = DownloadTheme(config.ThemeURL)
	if err != nil {
		// snip
		return
	}

	FormSubmittedCh <- struct{}{}
}
