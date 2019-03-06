package main

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/sirupsen/logrus"

	"github.com/uhayate/datastore-keyloader-demo/database"
)

func main() {
	var (
		ctx   = context.Background()
		err   error
		users []*database.User
	)

	client, err := datastore.NewClient(ctx, "yourProjectID")
	if err != nil {
		logrus.Errorf("fail to create datastore client")
		return
	}

	query := datastore.NewQuery("User")
	_, err = client.GetAll(ctx, query, &users)
	if err != nil {
		logrus.Errorf("fail to get users: %s", err)
	}

	logrus.Infof("Total users: %d\n", len(users))
}
