package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hpcloud/tail"
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

	sumchan := make(chan string)
	go syslog(sumchan)

	for {
		select {
		case <-c.Request().Context().Done():
			log.Printf("SSE client disconnected, ip: %v", c.RealIP())
			return nil
		case <-ticker.C:
			err := a.sendTime(w)
			if err != nil {
				return err
			}
			err = a.sendMemory(w)
			if err != nil {
				return err
			}
		case val := <-sumchan:
			err := a.sendSyslog(val, w)
			if err != nil {
				return err
			}
		}
	}
}

func (a *Api) sendTime(w *echo.Response) error {
	event := Event{
		Data: []byte(time.Now().Format(time.RFC1123)),
	}
	if err := event.MarshalTo(w); err != nil {
		return err
	}
	w.Flush()
	return nil
}

func (a *Api) sendMemory(w *echo.Response) error {
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

func syslog(c chan string) {
	t, err := tail.TailFile("/var/log/syslog", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		c <- line.Text
	}
}

func (a *Api) sendSyslog(result string, w *echo.Response) error {
	event := Event{
		Event: []byte("syslog"),
		Data:  []byte(fmt.Sprint(result)),
	}
	if err := event.MarshalTo(w); err != nil {
		return err
	}
	w.Flush()
	return nil
}
