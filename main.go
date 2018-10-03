package walrus

import (
	log "github.com/sirupsen/logrus"
)

func GetWalrus() {
	//fmt.Println("Log Example")

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
