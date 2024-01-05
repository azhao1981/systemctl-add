package main

import (
	"fmt"
	"os"
	"text/template"
)

const serviceTemplate = `[Unit]
Description={{.Description}}
After=network.target

[Service]
PermissionsStartOnly=true
Type=forking
ExecStart={{.ExecStart}}
ExecStartPost=
ExecStopPost=

; WorkingDirectory=
; User=
; Group=
; UMask=0022

[Install]
WantedBy=multi-user.target
`

type Service struct {
	Description string
	ExecStart   string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: systemctl-add <name> '<command>'")
		os.Exit(1)
	}

	name := os.Args[1]
	command := os.Args[2]
	filename := "/etc/systemd/system/" + name + ".service"

	service := Service{
		Description: name,
		ExecStart:   command,
	}

	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, service)
	if err != nil {
		panic(err)
	}

	fmt.Printf("created %s\ntry:\nsudo systemctl daemon-reload\nsudo systemctl start %s\n\nsudo systemctl enable %s\n", filename, name, name)
}
