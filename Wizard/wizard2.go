// Called after successful form submission
err := DownloadTheme(config.ThemeURL)
if err != nil {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	ws.ErrorLogger.Println("Error downloading and extracting theme:", err)
	return
}
