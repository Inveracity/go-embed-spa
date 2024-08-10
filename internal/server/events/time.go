package events

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func TimeStream(c echo.Context) error {
	log.Printf("SSE client connected to time stream, ip: %v", c.RealIP())

	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.Request().Context().Done():
			log.Printf("SSE client disconnected from time stream, ip: %v", c.RealIP())
			return nil
		case <-ticker.C:
			err := sendTime(w)
			if err != nil {
				return err
			}
		}
	}
}

func sendTime(w *echo.Response) error {
	event := Event{
		Data: []byte(time.Now().Format(time.RFC1123)),
	}
	if err := event.MarshalTo(w); err != nil {
		return err
	}
	w.Flush()
	return nil
}
