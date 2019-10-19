package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/setheck/push-relay/push"
	"github.com/setheck/push-relay/util"
)

type PushRelay struct {
	c       *Config
	api     Api
	pushers map[string]push.PusherSender
}

func NewPushRelay(c *Config) *PushRelay {
	return &PushRelay{
		c:       c,
		pushers: make(map[string]push.PusherSender),
	}
}

func (p *PushRelay) RelayHandler(pusherName string) http.HandlerFunc {
	pusher, ok := p.pushers[pusherName]
	if !ok {
		return func(w http.ResponseWriter, r *http.Request) {
			data, _ := httputil.DumpRequest(r, false)
			log.Println(strings.Fields(string(data)))
			http.NotFound(w, r)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := httputil.DumpRequest(r, false)
		log.Println(strings.Fields(string(data)))
		body, _ := ioutil.ReadAll(r.Body)
		var m push.Message
		if err := json.Unmarshal(body, &m); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if resp, err := pusher.Send(m); err != nil {
			http.Error(w, err.Error(), 400)
		} else {
			w.WriteHeader(resp.StatusCode)
		}
	}
}

func (p *PushRelay) Init() error {
	log.Println("push-relay initializing")
	var po *push.Pushover
	if err := p.c.Load("pushover", &po); err != nil {
		return err
	}
	p.pushers["pushover"] = po

	if err := p.c.Load("api", &p.api); err != nil {
		return err
	}
	p.api.AddHandler("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := httputil.DumpRequest(r, false)
		log.Println(strings.Fields(string(data)))
		if _, err := fmt.Fprintf(w, "Alive"); err != nil {
			log.Println(err)
		}
	})
	p.api.AddHandler("/push", p.RelayHandler("pushover"))
	p.api.ListenAndServe()
	log.Println("push-relay initialization complete")
	return nil
}

func (p *PushRelay) Shutdown() {
	log.Println("push-relay shutting down")
	if err := p.api.Stop(); err != nil {
		log.Println(err)
	}
	log.Println("shutdown complete")
}

func Main() {
	pr := NewPushRelay(NewConfig())
	if err := pr.Init(); err != nil {
		panic(err)
	}

	//po := push.NewPushOver("", "")
	//resp,err := po.Send(push.Message{
	//	Title:       "his",
	//	Message:     "another test",
	//})
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(resp)
	<-util.Signal()
	pr.Shutdown()
}
