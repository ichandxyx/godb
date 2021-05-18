package main

import (
	"context"
	"flag"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/ichandxyx/godb/ent"
	"github.com/ichandxyx/godb/pkg/api"
	"github.com/ichandxyx/godb/pkg/store"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
)

func main() {

	var dbPath, address string
	flag.StringVar(&dbPath, "db", "godb.db", "Path to the sqlite database file.")
	flag.StringVar(&address, "address", "0.0.0.0:8080", "The address to bind the HTTP server.")
	flag.Parse()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	client, err := ent.Open("sqlite3", "file:"+dbPath+"?cache=shared&mode=rwc&_fk=1")
	if err != nil {
		level.Error(logger).Log("err", err)
		return
	}
	level.Info(logger).Log("msg", "successfully connected to database")
	if err := client.Schema.Create(context.Background()); err != nil {
		level.Error(logger).Log("msg", "auto migration", "err", err)
	}

	str := store.New(client)
	apy := api.New(str, log.With(logger, "component", "api"))
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		apy.Register(r)
	})
	level.Info(logger).Log("msg", "starting web server", "addr", address)
	if err := http.ListenAndServe(address, r); err != nil && err != http.ErrServerClosed {
		level.Error(logger).Log("msg", "http server listen", "err", err)
	}

}
