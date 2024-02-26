package evnt

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Service() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{topic}", subscribe)
	return mux
}

type Event struct {
	Name string
	HTML string
}

var listeners = map[string][]chan Event{}

func Broadcast(topic, name, html string) {
	for _, events := range listeners[topic] {
		log.Println("broadcasting ", topic, name)
		events <- Event{name, html}
	}
}

func subscribe(w http.ResponseWriter, r *http.Request) {
	topic, events := r.PathValue("topic"), make(chan Event)
	listeners[topic] = append(listeners[topic], events)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	log.Println("Waiting for events")
	fmt.Fprintf(w, "event: ping\ndata: {\"foo\":\"bar\"}\n\n")
	flusher.Flush()
	for {
		select {
		case event := <-events:
			log.Printf("event: %s\ndata: \"%s\"\n\n", event.Name, strings.ReplaceAll(event.HTML, "\n", ""))
			fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event.Name, strings.ReplaceAll(event.HTML, "\n", ""))
			flusher.Flush()
		case <-r.Context().Done():
			for i, l := range listeners[topic] {
				if l == events {
					listeners[topic] = append(listeners[topic][:i], listeners[topic][i+1:]...)
					break
				}
			}
			close(events)
			return
		}
	}
}
