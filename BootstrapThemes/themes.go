func DownloadTheme(themeName string) error {
	zipURL := fmt.Sprintf("%s/%s/%s.zip", baseURL, tagVer, themeName)
	zipFile := fmt.Sprintf("%s.zip", themeName)

	fmt.Printf("Downloading %s...\n", zipURL)
	if err := downloadFile(zipURL, zipFile); err != nil {
		return fmt.Errorf("error downloading theme '%s': %v", themeName, err)
	}

	fmt.Println("Extracting files...")
	if err := unzip(zipFile, destDir); err != nil {
		return fmt.Errorf("error extracting files: %v", err)
	}

	fmt.Printf("Theme '%s' extracted to '%s' directory successfully.\n", themeName, destDir)
	if err := os.Remove(zipFile); err != nil {
		fmt.Printf("Warning: error deleting zip file '%s': %v\n", zipFile, err)
	}
}
