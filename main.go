package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dre1080/recovr"
	"github.com/gangozero/hackzurich2020-be/components/server"
	"github.com/gangozero/hackzurich2020-be/generated/restapi"
	"github.com/gangozero/hackzurich2020-be/generated/restapi/operations"
	"github.com/go-openapi/loads"

	"github.com/jackc/pgx/v4/pgxpool"
)

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	recovery := recovr.New()
	return recovery(handler)
}
func main() {

	log.Printf("App started")

	// init DB connection
	connectionString := os.Getenv("POSTGRESQLCONNSTR_URL")
	if connectionString == "" {
		fmt.Fprintln(os.Stderr, "Unable to find connection string in POSTGRES_URL:")
		os.Exit(1)
	}

	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to connect to database:", err)
		os.Exit(1)
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	srv := server.NewServer(pool)

	api := operations.NewHackzurich2020BeAPI(swaggerSpec)
	api.BearerAuth = srv.Validate

	api.UserLoginHandler = srv.UserLoginHandler()
	api.UserGetProductListHandler = srv.UserGetProductListHandler()
	api.UserDeleteProductHandler = srv.UserDeleteProductHandler()
	api.UserGetRecipeMatchListHandler = srv.UserGetRecipeMatchListHandler()
	api.UserGetRecipeDetailsHandler = srv.UserGetRecipeDetailsHandler()
	api.UserGetChatItemsHandler = srv.UserGetChatItemsHandler()

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080
	server.EnabledListeners = []string{"http"}

	server.SetHandler(setupGlobalMiddleware(server.GetHandler()))

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
