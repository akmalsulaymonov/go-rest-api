package main

import (
	"context"
	"fmt"

	"github.com/akmalsulaymonov/go-rest-api/internal/comment"
	"github.com/akmalsulaymonov/go-rest-api/internal/db"
	transportHttp "github.com/akmalsulaymonov/go-rest-api/internal/transport/http"
)

// Run - is going to be responsible for the instantiation and startup
// of our go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate the database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully connected and pinged database")

	cmtService := comment.NewService(db)

	/*
		// PostComment - posts comment
		cmtService.PostComment(
			context.Background(),
			comment.Comment{
				ID:     "2a746118-d01e-11ed-afa1-0242ac120011",
				Slug:   "slug-aki",
				Author: "Azizjon",
				Body:   "Hello mama",
			},
		)

		// GetComment - get comment by uuid
		fmt.Println(cmtService.GetComment(
			context.Background(),
			"2a115dec-b4d8-4074-987d-c47575a3065f",
		))
	*/

	// transport
	fmt.Println("Transport starts...")
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
