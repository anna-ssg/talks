func eventsHandler(w http.ResponseWriter, r *http.Request) {
	if !reloadPageBool.Load() {
		return
	}

	event := "event:\ndata:\n\n"
	_, err := w.Write([]byte(event))
	if err != nil {
		log.Fatal(err)
	}
	w.(http.Flusher).Flush()

	reloadPageBool.CompareAndSwap(true, false)
}
