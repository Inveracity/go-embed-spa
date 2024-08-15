package events

import (
	"encoding/json"
	"log"

	"github.com/hpcloud/tail"
	"github.com/labstack/echo/v4"
)

func StreamSyslog(c echo.Context) error {
	log.Printf("SSE client connected to syslog stream, ip: %v", c.RealIP())
	w := c.Response()
	w.Header().Set("Content-Type", "application/stream+json")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	syslogchan := make(chan string)
	go syslogRead(syslogchan)

	for {
		select {
		case <-c.Request().Context().Done():
			log.Printf("SSE client disconnected from syslog stream, ip: %v", c.RealIP())
			return nil
		case val := <-syslogchan:
			err := syslogSend(val, w)
			if err != nil {
				return err
			}
		}
	}
}

func syslogRead(c chan string) {
	t, err := tail.TailFile("/var/log/syslog", tail.Config{Follow: true})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		// time.Sleep(80 * time.Microsecond)
		c <- line.Text
	}
}

func syslogSend(result string, w *echo.Response) error {
	enc := json.NewEncoder(w)
	if err := enc.Encode(Msg{Msg: result, EventType: "syslog"}); err != nil {
		return err
	}
	w.Flush()
	return nil
}
