package main

import (
	"errors"
	"slices"

	"github.com/bndrmrtn/go-gale"
	"github.com/bottyaneu/goalcounter/config"
	"github.com/bottyaneu/goalcounter/handlers"
	"github.com/bottyaneu/goalcounter/middlewares"
)

func NewApiServer() *gale.Gale {
	conf := config.Api()
	app := gale.New(&conf)
	server := NewWSServer(app)
	store := gale.NewMemStorage()

	// Register custom validators
	registerValidators(app)

	r := app.Group("/", middlewares.CORSMiddleware)

	// Returns the list of teams
	r.Get("/", handlers.HandleGetTeams).Name("teams")
	// Increments the goal of the team
	r.Get("/teams/{team@team}/increment", handlers.HandleIncrementTeam(store, server)).Name("teams.increment")
	// Returns the scoreboard
	r.Get("/scoreboard", handlers.HandleGetScoreBoard(store)).Name("scoreboard")
	// Resets the scoreboard
	r.Get("/reset", handlers.HandleResetScoreBoard(store, server)).Name("reset")

	if conf.Mode == gale.Development {
		app.Use(gale.NewUIDevtools())
	}

	// Dump the routes to the console
	app.Dump()
	return app
}

func registerValidators(r gale.RouterParamValidator) {
	// Checking if route param is a team or not
	r.RegisterRouteParamValidator("team", func(value string) (string, error) {
		if !slices.Contains(config.Teams(), value) {
			return "", errors.New("invalid team")
		}
		return value, nil
	})
}
