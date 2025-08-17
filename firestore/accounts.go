package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"your_module_path/models" // Replace with your actual module path
)

// SaveAccount saves an Account struct to the 'accounts' collection in Firestore.
func SaveAccount(app *firebase.App, account models.Account) (string, error) {
	ctx := context.Background()

	client, err := app.Firestore(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get Firestore client: %w", err)
	}
	// Defer closing the client. In a real application, you might manage
	// the client lifecycle differently, perhaps keeping one instance
	// for the application's lifetime.
	defer client.Close()

	// Add the account data as a new document to the 'accounts' collection.
	// Firestore will automatically generate a unique ID for the document.
	ref, _, err := client.Collection("accounts").Add(ctx, account)
	if err != nil {
		return "", fmt.Errorf("failed to add account to Firestore: %w", err)
	}

	return ref.ID, nil
}