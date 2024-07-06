func DownloadTheme(themeName string) error {
	if err := downloadFile(zipURL, zipFile); err != nil {
		return fmt.Errorf("error downloading theme '%s': %v", themeName, err)
	}

	// Extracting files too a destination folder
	// and ensuring that the paths are safe before copying to
	// the desination path
	if err := unzip(zipFile, destDir); err != nil {
		return fmt.Errorf("error extracting files: %v", err)
	}

	if err := os.Remove(zipFile); err != nil {
		fmt.Printf("Warning: error deleting zip file '%s': %v\n", zipFile, err)
	}

	return nil
}
