package core

import (
	"os"

	"github.com/it-sos/golibs/global/consts"
)

func GetDevBasePath() string {
	env, _ := os.LookupEnv(consts.DevBasePathKey)
	return env
}
