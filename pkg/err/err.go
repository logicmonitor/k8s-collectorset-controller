package err

import (
	log "github.com/sirupsen/logrus"
)

func RecoverError() {
	if err := recover(); err != nil {
		log.Errorf("recover error: %s", err)
	}
}
