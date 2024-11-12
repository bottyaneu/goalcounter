package handlers

import (
	"encoding/json"
	"math/big"

	"github.com/bndrmrtn/go-gale"
	"github.com/bottyaneu/goalcounter/config"
	"github.com/sirupsen/logrus"
)

// HandleGetTeams returns all available teams
func HandleGetTeams(c gale.Ctx) error {
	return c.JSON(gale.Map{
		"teams": config.Teams(),
	})
}

// HandleIncrementTeam increments the goals of a team
func HandleIncrementTeam(store gale.SessionStore, ws gale.WSServer) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		// Get team from route param
		team := c.Param("team")
		var goals int64

		// Check if goals exists in store
		if ok := store.Exists("goals." + team); ok {
			rawGoals, err := store.Get("goals." + team)
			if err != nil {
				return err
			}

			// Convert raw byte value to int64
			goals = big.NewInt(0).SetBytes(rawGoals).Int64()
		}

		// Increment goals
		goals++

		// Set new goals value to store or return error
		if err := store.Set("goals."+team, big.NewInt(goals).Bytes()); err != nil {
			return err
		}

		// Spin up a new goroutine to broadcast new goals value
		go func() {
			data, err := json.Marshal(gale.Map{
				"event": "increment",
				"goals": goals,
				"team":  team,
			})
			if err != nil {
				logrus.Error("Failed to marshal data: ", err)
				return
			}

			// Broadcast new goals value to all connected clients
			ws.Broadcast(data)
		}()

		// Return new goals value
		return c.JSON(gale.Map{
			"increment": true,
		})
	}
}

// HandleGetScoreBoard returns the current score board
func HandleGetScoreBoard(store gale.SessionStore) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		teams := config.Teams()
		scores := make(map[string]int64, len(teams))

		// Get goals for each team
		for _, team := range teams {
			var goals int64
			if ok := store.Exists("goals." + team); ok {
				rawGoals, err := store.Get("goals." + team)
				if err != nil {
					return err
				}

				goals = big.NewInt(0).SetBytes(rawGoals).Int64()
			}
			scores[team] = goals
		}

		// Return scores
		return c.JSON(scores)
	}
}

// HandleResetScoreBoard resets the score board
func HandleResetScoreBoard(store gale.SessionStore, ws gale.WSServer) gale.HandlerFunc {
	return func(c gale.Ctx) error {
		teams := config.Teams()

		// Delete goals for each team
		for _, team := range teams {
			if err := store.Del("goals." + team); err != nil {
				return err
			}
		}

		// Spin up a new goroutine to broadcast reset event
		go func() {
			data, err := json.Marshal(gale.Map{
				"event": "reset",
			})
			if err != nil {
				logrus.Error("Failed to marshal data: ", err)
				return
			}

			// Broadcast reset event to all connected clients
			ws.Broadcast(data)
		}()

		// Return reset event
		return c.JSON(gale.Map{
			"reset": true,
		})
	}
}
