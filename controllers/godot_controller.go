package controllers

import (
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func GodotPollingHandler(app *pocketbase.PocketBase, e *core.RequestEvent) error {

	transaction_id := e.Request.PathValue("id")

	record, err := app.FindFirstRecordByData("transactions", "transaction_id", transaction_id)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Transaction not found",
		})
	}

	if record.GetString("status") == "COMPLETED" {
		// Here you can add any additional data you want to send back to the Godot game
		return e.JSON(http.StatusForbidden, map[string]string{
			"message": "Token already shared",
		})
	}

	record.Set("status", "COMPLETED")
	err = app.Save(record)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error saving transaction",
		})
	}

	return e.JSON(http.StatusOK, map[string]string{
		"token": record.GetString("token"),
	})
}
