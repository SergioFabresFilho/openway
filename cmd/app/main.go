package main

import (
	"github.com/sirupsen/logrus"
	"ssff.com.br/openway/common/db"
	"ssff.com.br/openway/controlplane"
	"ssff.com.br/openway/dataplane"
)

func main() {
	logrus.SetFormatter(
		&logrus.TextFormatter{FullTimestamp: true},
	)

	db.InitDB()
	controlplane.InitAsync()
	dataplane.Init()
}
