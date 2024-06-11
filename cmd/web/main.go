package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	app := &application{
		logger: logger,
	}

	conn, err := pgx.Connect(context.Background(), "postgresql://root:secret@localhost:5432/smartsheets")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server on ", "addr", *addr)

	err = srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
