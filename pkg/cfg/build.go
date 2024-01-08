package cfg

import "github.com/spf13/viper"

type Build struct {
	Version   string `mapstructure:"version"`
	BuildTime string `mapstructure:"build_time"`
	Builder   string `mapstructure:"builder"`
}

func SetBuild(version, buildTime, builder string) {
	viper.SetDefault("build.version", version)
	viper.SetDefault("build.build_time", buildTime)
	viper.SetDefault("build.builder", builder)
}

func ShowBuild() string {
	return "Build Verison " +
		viper.GetString("build.version") +
		" BuildTime " + viper.GetString("build.build_time") +
		" by " + viper.GetString("build.builder")
}
