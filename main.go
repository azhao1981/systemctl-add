package main

import (
	"fmt"
	"os"

	"github.com/azhao1981/systemctl-add/internal/service"
	"github.com/azhao1981/systemctl-add/pkg/cfg"
	"github.com/spf13/viper"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: systemctl-add <name> '<command>'")
		os.Exit(1)
	}

	name := os.Args[1]
	command := os.Args[2]
	viper.Set("service.description", name)
	viper.Set("service.exec.start", command)

	cfg.Init("")
	serviceData := cfg.Cfg.Service

	filename := serviceData.Dir + name + ".service"

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBody, err := service.Execute(serviceData)
	if err != nil {
		panic(err)
	}
	file.WriteString(fileBody)

	fmt.Printf("created %s\ntry:\nsudo systemctl daemon-reload\nsudo systemctl start %s\n\nsudo systemctl enable %s\n", filename, name, name)
}
