func (ws *WizardServer) handleSubmit(w http.ResponseWriter, r *http.Request) {

	// Decode JSON data from the request body into the LayoutConfig struct
	var config parser.LayoutConfig
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		ws.ErrorLogger.Println("Error decoding JSON:", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = ws.writeConfigToFile(config)
	if err != nil {
		ws.ErrorLogger.Println("Error writing config to file:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	FormSubmittedCh <- struct{}{} // Signal that a form has been submitted successfully
}
