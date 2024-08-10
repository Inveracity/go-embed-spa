package events

import (
	"encoding/json"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v4/mem"
)

func StreamMemory(c echo.Context) error {
	log.Printf("SSE client connected to memory stream, ip: %v", c.RealIP())

	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.Request().Context().Done():
			log.Printf("SSE client disconnected from memory stream, ip: %v", c.RealIP())
			return nil
		case <-ticker.C:
			err := sendMemory(w)
			if err != nil {
				return err
			}
		}
	}
}

func sendMemory(w *echo.Response) error {
	v, _ := mem.VirtualMemory()
	b, _ := json.Marshal(v)
	event := Event{
		Event: []byte("memory"),
		Data:  b,
	}
	if err := event.MarshalTo(w); err != nil {
		return err
	}
	w.Flush()
	return nil
}
