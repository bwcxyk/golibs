package variable

import (
	"flag"
	"os"
	"strings"

	"github.com/bwcxyk/golibs/core"
	"github.com/bwcxyk/golibs/global/consts"
)

var (
	// BasePath 项目根目录
	BasePath = ""
	// ConfPath 配置目录
	ConfPath = ""
)

var workdir = flag.String("w", "", "指定工作目录 -w /workdir")
var confdir = flag.String("c", "", "指定配置目录 -c /confdir")

func init() {
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
		BasePath = core.GetDevBasePath()
		pwd, _ := os.Getwd()
		bases := []string{"/src/backend", "/golibs"}
		// 如果未设置开发根目录，则尝试自动识别根目录。
		for _, base := range bases {
			if strings.Contains(pwd, base) {
				BasePath = pwd[0 : strings.Index(pwd, base)+len(base)]
			}
		}
		if BasePath == "" {
			panic("Environment variable [" + consts.DevBasePathKey + "] not set. Specify the project root directory.")
		}
	} else {
		flag.Parse()
		if *workdir != "" {
			BasePath = *workdir
		} else if pwd, err := os.Getwd(); err == nil {
			BasePath = pwd
		}
	}
	ConfPath = BasePath
	if *confdir != "" {
		ConfPath = *confdir
	}
}
