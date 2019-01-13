package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"os"
)

func init() {

}

func connect(ctx context.Context) (*firestore.Client, error) {
	return firestore.NewClient(ctx, os.Getenv("FIREBASE_PROJECTID"))
}
