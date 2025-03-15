package log

import (
	"github.com/rs/zerolog/log"
)

func Boot(message string) {
	log.Info().Msg(message)
}
