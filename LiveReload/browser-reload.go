func StartLiveReload() {
	// ...
	if lr.traverseDirectory(rootDir) {
		cmd.VanillaRender(lr.siteDataPath)
		reloadPageBool.CompareAndSwap(false, true)
	}
	// ...
}
