package cfg

type Service struct {
	Description      string `mapstructure:"description"`
	ExecStart        string `mapstructure:"exec_start"`
	WorkingDirectory string `mapstructure:"working_directory"`
	User             string `mapstructure:"user"`
	Group            string `mapstructure:"group"`
	UMask            string `mapstructure:"umask" default:"0022"`

	Dir string `mapstructure:"dir" default:"/etc/systemd/system/"`
}
