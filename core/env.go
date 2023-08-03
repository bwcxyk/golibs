package core

import (
	"os"

	"github.com/bwcxyk/golibs/global/consts"
)

func GetDevBasePath() string {
	env, _ := os.LookupEnv(consts.DevBasePathKey)
	return env
}
