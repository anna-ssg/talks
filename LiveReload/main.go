//START OMIT
type liveReload struct {
	errorLogger   *log.Logger
	fileTimes     map[string]time.Time
	rootDirs      []string // Directories to monitor, so add or remove as needed
	extensions    []string // File extensions to monitor
	serverRunning bool
	siteDataPath  string
}

//END OMIT