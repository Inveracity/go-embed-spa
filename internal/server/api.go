package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v4/mem"
)

type Api struct {
}

func (a *Api) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "World!"})
}

func (a *Api) Stream(c echo.Context) error {
	log.Printf("SSE client connected, ip: %v", c.RealIP())

	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-c.Request().Context().Done():
			log.Printf("SSE client disconnected, ip: %v", c.RealIP())
			return nil
		case <-ticker.C:
			event := Event{
				Data: []byte(time.Now().Format(time.RFC1123)),
			}
			if err := event.MarshalTo(w); err != nil {
				return err
			}
			w.Flush()
		}
	}
}

func (a *Api) Diagnostics(c echo.Context) error {
	log.Printf("SSE client connected, ip: %v", c.RealIP())

	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-c.Request().Context().Done():
			log.Printf("SSE client disconnected, ip: %v", c.RealIP())
			return nil
		case <-ticker.C:
			v, _ := mem.VirtualMemory()
			b, _ := json.Marshal(v)
			event := Event{
				Event: []byte("diagnostics"),
				Data:  b,
			}
			if err := event.MarshalTo(w); err != nil {
				return err
			}
			w.Flush()
		}
	}
}
