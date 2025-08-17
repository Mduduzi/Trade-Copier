package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	firebase "firebase.google.com/go"
	"your_module_path/api" // Import the api package for middleware and context functions
	"trade-copier/firestore" // Import the firestore package
	"your_module_path/models" // Replace with your actual module path
	"trade-copier/api"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/accounts/link", linkAccountHandler).Methods("POST")

	// Add other routes here as needed

	return r
}

// SetupRouter initializes and configures the API router with the Firebase app instance.
func SetupRouter(app *firebase.App) *mux.Router {
	r := mux.NewRouter()

	// Apply the authentication middleware to routes that require authentication
	r.Use(api.AuthMiddleware(app))

	r.HandleFunc("/api/v1/accounts/link", linkAccountHandler(app)).Methods("POST")

	return r
}

func linkAccountHandler(app *firebase.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	// Retrieve the authenticated user ID from the request context
	userID, ok := api.GetUserIDFromContext(r.Context())
	if !ok {
		// This case should ideally be handled by the middleware, but as a fallback
		http.Error(w, `{"error": "User not authenticated"}`, http.StatusUnauthorized)
		return
	}

	if account.UserID == "" {
		// Input validation (excluding UserID, as we get it from context)
	}

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request payload: %v", err), http.StatusBadRequest)
		return
	}

	if account.Platform == "" {
		http.Error(w, `{"error": "Platform is required"}`, http.StatusBadRequest)
		return
	}
	if account.Login == "" && account.APIKey == "" {
		http.Error(w, `{"error": "Either Login or APIKey is required"}`, http.StatusBadRequest)
		return
	}

	// Override the UserID from the request body with the authenticated user ID
	account.UserID = userID

	// Save the account data to Firestore
	accountID, err := firestore.SaveAccount(app, account) // Changed to use imported firestore
	if err != nil {
		// Log the error for debugging purposes
		fmt.Printf("Error saving account to Firestore: %v\n", err)
		http.Error(w, `{"error": "Failed to save account"}`, http.StatusInternalServerError)
		return
	}

	// Return a success response with the actual account ID
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"accountId": accountID, // Use the actual document ID from Firestore
		"status":    "linked",
	}
	json.NewEncoder(w).Encode(response)
	}
}