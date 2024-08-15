package events

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shirou/gopsutil/v4/mem"
)

func StreamMemory(c echo.Context) error {
	log.Printf("SSE client connected to memory stream, ip: %v", c.RealIP())

	w := c.Response()
	w.Header().Set("Content-Type", "application/stream+json")
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
	v, err := mem.VirtualMemory()
	if err != nil {
		return err
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(Msg{Msg: fmt.Sprintf("%v", v.UsedPercent), EventType: "memory"}); err != nil {
		return err
	}
	w.Flush()
	return nil
}
