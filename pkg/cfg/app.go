package cfg

import (
	"os"
	"path/filepath"
	"strings"
)

// 本文件由项目各自修改
var (
	Cfg           *App
	DEFAULT_NAMES = "system-add,base"
	DEFAULT_PATH  = "~/.system-add/"
)

type App struct {
	MustConfig bool    `mapstructure:"must.config" default:"false"`
	Service    Service `mapstructure:"service"`
}

func Init(name string, paths ...string) *App {
	Cfg = &App{}

	initDefault(Cfg)
	initEnv()
	initConfigFile(Cfg, name, paths...)

	// 只显示当前配置文件的变量
	CmdShowEnv()

	return Cfg
}

func Load(paths ...string) *App {
	env_path := os.Getenv("CONFIG_FILE")
	if env_path != "" {
		paths = append([]string{env_path}, paths...)
	}
	if len(paths) == 0 {
		return Init("")
	}
	newPaths := make([]string, 0)
	names := make([]string, 0)
	for _, path := range paths {
		dir, name, _ := splitPath(path)
		names = append(names, name)
		newPaths = append(newPaths, dir)
	}
	return Init(strings.Join(names, ","), newPaths...)
}

func splitPath(path string) (dir, fileName, fileType string) {
	dir, file := filepath.Split(path)
	ext := filepath.Ext(file)
	fileName = strings.TrimSuffix(file, ext)
	fileType = strings.TrimPrefix(ext, ".")
	return dir, fileName, fileType
}
