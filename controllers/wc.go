package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

var database *sql.DB

func SetDB(db *sql.DB) {
	database = db
}

func Get(ctx context.Context) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		type Workcenter struct {
			ID              int
			Name            string
			CurrentProduct  string
			Status          int
			EscalationLevel int
			StatusSetAt     time.Time
		}

		result, err := database.QueryContext(ctx,
			`SELECT
			id, wc_name, current_product, wc_status, escalation_level
		FROM workcenters
		ORDER BY wc_name`)
		if err != nil {
			msg := fmt.Sprintf("failed to retrieve workcententers from database: %v", err)
			log.Println(msg)
			fmt.Errorf(msg)
		}
		var wc Workcenter
		for result.Next() {
			err := result.Scan(&wc.ID, &wc.Name, &wc.CurrentProduct, &wc.Status, &wc.EscalationLevel)
			fmt.Println(wc.ID, wc.Name, wc.CurrentProduct, wc.Status, wc.EscalationLevel)
			if err != nil {
				fmt.Errorf("failed to retrieve workcenter fields database: %v", err)
			}
		}
	}
}
