package events

import (
	"encoding/json"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

type msg struct {
	Msg       string `json:"msg"`
	EventType string `json:"eventtype"`
}

func TimeStream(c echo.Context) error {
	log.Printf("SSE client connected to time stream, ip: %v", c.RealIP())

	w := c.Response()
	w.Header().Set("Content-Type", "application/json")
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
	message := msg{
		Msg:       time.Now().Format(time.RFC1123),
		EventType: "time",
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(message); err != nil {
		return err
	}

	w.Flush()
	return nil
}
