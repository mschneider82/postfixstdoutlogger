package main

import (
	"os"
	"os/exec"

	"github.com/alecthomas/kingpin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/mcuadros/go-syslog.v2"
)

var (
	socketfile = kingpin.Flag("socketfile", "create a unixsocket file").Default("/dev/log").String()
	overwrite  = kingpin.Flag("overwrite", "set to overwrite a socket file if it exists").Bool()
)

const timeFormat = "2006-01-02T15:04:05.000-07:00"

func main() {
	kingpin.Parse()
	zerolog.TimeFieldFormat = timeFormat

	if *overwrite {
		if _, err := os.Stat(*socketfile); err == nil {
			if err := os.Remove(*socketfile); err != nil {
				panic(err)
			}
		}
	}
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC3164)
	server.SetHandler(handler)
	server.ListenUnixgram(*socketfile)

	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			content := logParts["content"]
			hostname := logParts["hostname"]
			log.Printf("%s %s", hostname, content)
		}
	}(channel)

	cmd := exec.Command("postfix", "start-fg")
	log.Printf("Running postfix start-fg...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}
