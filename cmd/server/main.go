package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/othavioBF/imperium/internal/core"
	"github.com/othavioBF/imperium/internal/infra/pgstore"
	"github.com/othavioBF/imperium/internal/session"
)

func main() {
	gob.Register(uuid.UUID{})

	ctx := context.Background()

	pool, err := pgstore.Init(ctx)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	s := session.InitSessionManager(pool)

	api := core.InjectDependencies(pool, s)

	api.BindRoutes()

	port := os.Getenv("APP_PORT")

	fmt.Println("Starting Server on port :%s\n", port)
	if err := http.ListenAndServe(":"+port, api.Router); err != nil {
		panic(err)
	}
}
