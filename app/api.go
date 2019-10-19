package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Api struct {
	Port int

	server *http.Server
	doneCh chan struct{}
	wg sync.WaitGroup
}

func NewApi(port int) Api {
	return Api{
		Port:   port,
		doneCh: make(chan struct{}),
	}
}

func (a *Api) Start() {
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		log.Println("hit /")
		out,_ := json.Marshal(struct {
			Time int64
			Message string
		}{
			Time:time.Now().Unix(),
			Message:"testMessage",
		})
		if _,err := writer.Write(out); err != nil {
			log.Println(err)
		}
	})
	a.server = &http.Server{
		Addr:              fmt.Sprintf(":%d", a.Port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	go func() {
		<-a.doneCh
		if err := a.server.Close(); err != nil {
			log.Println(err)
		}
	}()
	go func() {
		a.wg.Add(1)
		defer a.wg.Done()
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed{
			log.Println(err)
		}
	}()
}

func (a *Api) Stop() {
	close(a.doneCh)
}