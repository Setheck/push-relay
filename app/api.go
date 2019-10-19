package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Api struct {
	Port int

	server *http.Server
	router *mux.Router
	wg     sync.WaitGroup
}

func (a *Api) AddHandler(p string, fn http.HandlerFunc) *mux.Route {
	if a.router == nil {
		a.router = mux.NewRouter()
	}
	return a.router.HandleFunc(p, fn)
}

func (a *Api) ListenAndServe() {
	if a.server == nil {
		go a.listenAndServe()
	}
}

func (a *Api) listenAndServe() {
	a.wg.Add(1)
	defer a.wg.Done()
	a.server = &http.Server{
		Addr:           fmt.Sprintf(":%d", a.Port),
		Handler:        a.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	endpoints := make([]string, 0)
	a.router.Walk(func(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {
		if r, err := route.URL(); err != nil {
			return err
		} else {
			endpoints = append(endpoints, r.String())
		}
		return nil
	})
	log.Printf("api starting on %q endpoints %q\n", a.server.Addr, endpoints)
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}

func (a *Api) Stop() error {
	if a.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if err := a.server.Shutdown(ctx); err != nil {
			return err
		}
		a.wg.Wait()
	}
	a.wg = sync.WaitGroup{}
	return nil
}
