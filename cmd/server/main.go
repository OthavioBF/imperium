package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/othavioBF/imperium/internal/core/di"
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

	api := di.InjectDependencies(pool, s)

	api.BindRoutes()

	fmt.Println("Starting Server on port :3080")
	if err := http.ListenAndServe("localhost:3080", api.Router); err != nil {
		panic(err)
	}
}
