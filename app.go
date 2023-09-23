package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"

	"github.com/BurntSushi/toml"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

var Settings struct {
	SmtpUser     string `valid:"-"`
	SmtpPassword string `valid:"-"`
	SmtpServer   string `valid:"-"`
	SmtpPort     string `valid:"int,required"`
	Recipients   []string
	Name         string `valid:"-"`
	Button1Label string `valid:"-"`
	Button2Label string `valid:"-"`
}

var configFile = "config.toml"
var Button1Label = "I'm OK"
var Button2Label = "Call Me"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.Init()
}

func (a *App) domReady(ctx context.Context) {
	rt.EventsEmit(ctx, "setButtonText", Button1Label, Button2Label)
}

func (a *App) Init() {
	if _, err := toml.DecodeFile(configFile, &Settings); err != nil {
		log.Println(err)
	}
	// log.Println(Settings)

	if len(Settings.Button1Label) > 0 {
		Button1Label = Settings.Button1Label
	}

	if len(Settings.Button2Label) > 0 {
		Button2Label = Settings.Button2Label
	}
}

// AmOK sends the email that all is OK.
func (a *App) AmOk() string {
	smtpStatus := SendMessage("Status Updater - I'm OK", Settings.Name+" sent the message \"I'm OK\".")
	return fmt.Sprintf(smtpStatus)
}

// CallMe sends the email to call the requester.
func (a *App) CallMe() string {
	smtpStatus := SendMessage("Status Updater - Call Me", Settings.Name+" sent the message \"Call Me\".")
	return fmt.Sprintf(smtpStatus)
}

func SendMessage(subject string, message string) string {
	smtpStatus := "Email Sent"
	err := SendEmail(subject, message)
	if err != nil {
		smtpStatus = fmt.Sprint("Problem Sending Email: " + err.Error())
	}
	return fmt.Sprintf(smtpStatus)
}

// SendEmail provides the implementation of the EmailSender type for runtime usage.
func SendEmail(subject string, message string) error {
	// Set up authentication information.
	smtpUser := Settings.SmtpUser
	smtpServer := Settings.SmtpServer
	smtpPort := Settings.SmtpPort
	smtpPassword := Settings.SmtpPassword
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer)
	server := smtpServer + ":" + smtpPort
	from := "sender@example.org"

	msg := []byte("Subject: " + subject + "\r\n\r\n" +
		message + "\r\n")

	err := smtp.SendMail(server, auth, from, Settings.Recipients, msg)
	if err != nil {
		return err
	}

	return nil
}
