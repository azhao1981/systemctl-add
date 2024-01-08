package cfg

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/creasty/defaults"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initEnv() {
	// 如果两个 key 用一个组件，但是其中一个默认值不一样，需要手动设置
	// setNewDefault("api.port", "API_PORT", "8081")
	viper.AutomaticEnv()
	autoBindEnv()
}

type StructField struct {
	Key    string
	EnvKey string
	Type   string
	Value  string
}

var (
	DEFAULT_TAG  = "default"
	ENV_TAG      = "env"
	ENV_IGNORE   = []string{"MUTEX", "BUILD", "TEST"}
	StructFields = []StructField{}
	EnvCmd       = &cobra.Command{
		Use:   "env",
		Short: "show available env vars and bind keys",
		Long: `
Struct 中的字段,可以通过环境变量来设置,例如: VAULT_ADDR -> vault.addr 
可以通过设置 tag 来指定环境变量的名称,例如: env:"VAULT_ADDR"
default:"默认值"也会setDefault给viper`,
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set("show.env", true)
		},
	}
)

// viper.BindEnv vault.addr -> VAULT_ADDR
func autoBindEnv() {
	// 遍历结构体的字段
	app := App{}
	defaults.Set(app)

	t := reflect.TypeOf(app)
	val := reflect.ValueOf(app)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := val.Field(i)
		StructFields = append(StructFields, fieldInfo("", field, fieldValue)...)
	}
	for _, env := range StructFields {
		viper.SetDefault(env.Key, env.Value)
		viper.BindEnv(env.Key, env.EnvKey) // 绑定到环境变量
	}
}

func fieldInfo(prefix string, field reflect.StructField, fieldValue reflect.Value) []StructField {
	retFields := []StructField{}
	// 遍历嵌套结构体的字段
	if field.Type.Kind() == reflect.Struct {
		for j := 0; j < fieldValue.NumField(); j++ {
			nestedFieldType := field.Type.Field(j)
			keyPath := strings.ToLower(keyJoin(prefix, field.Name))
			fields := fieldInfo(keyPath, nestedFieldType, fieldValue.Field(j))
			retFields = append(retFields, fields...)
		}
	} else {
		keyPath := strings.ToLower(keyJoin(prefix, field.Name))
		envVarName := strings.ReplaceAll(strings.ToUpper(keyPath), ".", "_")
		if field.Tag.Get(ENV_TAG) != "" {
			envVarName = field.Tag.Get(ENV_TAG)
		}
		if IsStartWith(envVarName, ENV_IGNORE...) {
			return retFields
		}
		defaultVal := field.Tag.Get(DEFAULT_TAG)
		viperVal := viper.GetString(keyPath)
		if viperVal != "" {
			defaultVal = viperVal
		}
		retFields = append(retFields, StructField{Key: keyPath, EnvKey: envVarName, Value: defaultVal, Type: field.Type.String()})
	}
	return retFields
}

func keyJoin(prefix string, key string) string {
	if prefix == "" {
		return key
	}
	return prefix + "." + key
}

// 比较两个Gin的端口配置，需要不同
// setNewDefault("api.port", "API_PORT", "8081")
// setNewDefault("internal.port", "API_PORT", "8080")
func setNewDefault(viperKey, EnvKey, defaultValue string) {
	viper.SetDefault(viperKey, defaultValue)
	viper.BindEnv(viperKey, EnvKey)
}

func CmdShowEnv() {
	if viper.GetBool("show.env") {
		showEnv()
		os.Exit(1)
	}
}
func showEnv() {
	fmt.Println("    environment:")
	for _, env := range StructFields {
		if env.Type == "string" {
			fmt.Printf("      %s: \"%s\" \t\t# %s %s\n", env.EnvKey, viper.GetString(env.Key), env.Key, env.Type)
		} else {
			fmt.Printf("      %s: %s \t\t# %s %s\n", env.EnvKey, viper.GetString(env.Key), env.Key, env.Type)
		}
	}
}

func IsStartWith(s string, prefix ...string) bool {
	for _, p := range prefix {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
