package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/NoteToScreen/maily-go/maily"
	"github.com/willbarkoff/simplemailer/config"
)

func main() {
	configFile := flag.String("config", "config.toml", "The configuration file to use for sending mail")
	addresses := flag.String("addresses", "addresses.txt", "The list of email addresses to send to")
	templatePath := flag.String("templatePath", "templates/", "The path to the template directory to use when sending")
	template := flag.String("template", "message", "The template to send")
	failFast := flag.Bool("stopOnFail", false, "Whether to stop sending on a send failure")

	flag.Parse()

	conf := config.Configure(*configFile)

	context := maily.Context{
		FromAddress:  conf.FromAddress,
		FromDisplay:  conf.FromDisplay,
		SendDomain:   conf.SendDomain,
		SMTPHost:     conf.SMTPHost,
		SMTPPort:     conf.SMTPPort,
		SMTPUsername: conf.SMTPUsername,
		SMTPPassword: conf.SMTPPassword,
		TemplatePath: *templatePath,
	}

	file, err := os.Open(*addresses)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		email := scanner.Text()
		result, err := context.SendMail("", email, *template, nil, nil, nil)
		if err != nil {
			if *failFast {
				fmt.Fprintf(os.Stderr, "An error occured sending to %s:\n", email)
				panic(err)
			}

			fmt.Fprintf(os.Stderr, "An error occured sending to %s: %s\n", email, err.Error())
		} else {
			fmt.Printf("Sent to %s, message ID: %s\n", email, result.MessageID)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
