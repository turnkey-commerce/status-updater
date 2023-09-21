package main

import (
	"context"
	"fmt"
	"log"
	"net/smtp"

	"github.com/BurntSushi/toml"
)

type Recipient string

var Settings struct {
	SmtpUser     string `valid:"-"`
	SmtpPassword string `valid:"-"`
	SmtpServer   string `valid:"-"`
	SmtpPort     string `valid:"int,required"`
	Recipients   []string
	Name         string `valid:"-"`
}

var configFile = "config.toml"

func init() {

}

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
	Init()
}

func Init() {
	if _, err := toml.DecodeFile(configFile, &Settings); err != nil {
		log.Println(err)
	}
	// log.Println(Settings)
}

// AmOK sends the email that all is OK.
func (a *App) AmOk() string {
	smtpStatus := SendMessage("Status Updater - I'm OK", Settings.Name+" sent the message \"I'm OK\".")
	return fmt.Sprintf(smtpStatus)
}

// AmOK sends the email to call the requester.
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
