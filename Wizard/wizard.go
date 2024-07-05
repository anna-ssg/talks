// if the webconsole flag is passed
if webconsole {
	server := anna.NewWizardServer(":8080")
	go server.Start()
	<-anna.FormSubmittedCh // wait for response
	if err := server.Stop(); err != nil {
		annaCmd.InfoLogger.Println(err)
	}
	annaCmd.VanillaRenderManager()
	annaCmd.LiveReloadManager()
}