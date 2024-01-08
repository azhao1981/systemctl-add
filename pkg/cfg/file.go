package cfg

import (
	"fmt"
	"os"
	"strings"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
)

func initConfigFile(cfgV interface{}, name string, paths ...string) {
	if viper.ConfigFileUsed() != "" {
		err := viper.MergeInConfig()
		if err == nil {
			fmt.Println("ConfigLoaded: ", viper.ConfigFileUsed())
		}
	} else {
		setConfigFile(name, paths...)
	}
	if err := viper.Unmarshal(cfgV); err != nil {
		fmt.Printf("# Error unmarshaling config file: %s\n", err)
		os.Exit(1)
	}
	if viper.Get("debug") != nil {
		fmt.Printf("config init default %+v\n", cfgV)
	}
}

func setConfigFile(name string, paths ...string) {
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	if len(paths) == 0 {
		viper.AddConfigPath("./config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("../config")
		viper.AddConfigPath("../../config")
		viper.AddConfigPath("../../../config")
		viper.AddConfigPath(DEFAULT_PATH)
	}

	if name == "" {
		name = DEFAULT_NAMES
	}

	var errs []error
	var hasLoad bool

	for _, n := range strings.Split(name, ",") {
		viper.SetConfigName(n)
		err := viper.MergeInConfig()
		if err == nil {
			fmt.Println("ConfigLoaded: ", viper.ConfigFileUsed())
			hasLoad = true
		} else {
			errs = append(errs, err)
		}
	}
	if !hasLoad && viper.GetBool("must.config") {
		fmt.Println("# Error loading config files:")
		for _, err := range errs {
			fmt.Printf("  %s\n", err)
		}
		if !viper.GetBool("show.env") {
			os.Exit(1)
		}
	}
}

// 注意，Struct 中的default并不会被viper读取，所以需要手动设置
func initDefault(ptr interface{}) {
	if err := defaults.Set(ptr); err != nil {
		panic(err)
	}

	// viper 默认值,处理比较复杂的默认数据比如map,struct等
	// viper.SetDefault("redis.max_sub_goroutines", "200")
	// viper.SetDefault("api.port", "8081")
}
