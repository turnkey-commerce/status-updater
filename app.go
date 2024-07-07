package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/wailsapp/wails/v2/pkg/menu"
	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

var Settings struct {
	EncryptionKey string `valid:"-"`
	SmtpUser      string `valid:"-"`
	SmtpPassword  string `valid:"-"`
	SmtpServer    string `valid:"-"`
	SmtpPort      string `valid:"int,required"`
	FromEmail     string `valid:"-"`
	Recipients    []string
	Name          string `valid:"-"`
	Button1Label  string `valid:"-"`
	Button2Label  string `valid:"-"`
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

func (a *App) Menu() *menu.Menu {
	AppMenu := menu.NewMenu()
	ConfigureMenu := AppMenu.AddSubmenu("Configuration")
	ConfigureMenu.AddText("&Configure", nil, menuCallbackEmit(a, "configure"))
	return AppMenu
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

func menuCallbackEmit(a *App, eventName string, data ...interface{}) func(cd *menu.CallbackData) {
	return func(cd *menu.CallbackData) {
		rt.EventsEmit(a.ctx, eventName, data...)
	}
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

// Button1Action sends the message with the Button 2 labels.
func (a *App) Button1Action() string {
	smtpStatus := SendMessage("Status Updater - "+Settings.Button1Label, Settings.Name+" sent the message \""+Settings.Button1Label+"\".")
	return fmt.Sprintf(smtpStatus)
}

// Button2Action sends the message with the Button 2 labels.
func (a *App) Button2Action() string {
	smtpStatus := SendMessage("Status Updater - "+Settings.Button2Label, Settings.Name+" sent the message \""+Settings.Button2Label+"\".")
	return fmt.Sprintf(smtpStatus)
}

// SaveAction saves the configuration
func (a *App) SaveAction(smtpUserName string, smtpPassword string) string {
	err := EncryptCredentials(smtpUserName, smtpPassword)
	saveStatus := "Saved!"
	if err != nil {
		saveStatus = fmt.Sprint("Problem Saving Configuration: " + err.Error())
	}
	return fmt.Sprintf(saveStatus)
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
	smtpUser, smtpPassword, err := DecryptCredentials()
	if err != nil {
		return err
	}
	smtpServer := Settings.SmtpServer
	smtpPort := Settings.SmtpPort
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpServer)
	server := smtpServer + ":" + smtpPort
	fromEmail := Settings.FromEmail

	msg := []byte("Subject: " + subject + "\r\n\r\n" +
		message + "\r\n")

	err = smtp.SendMail(server, auth, fromEmail, Settings.Recipients, msg)
	if err != nil {
		return err
	}

	return nil
}

func DecryptCredentials() (smtpUser string, smtpPassword string, err error) {
	smtpUserEncrypted, err := hex.DecodeString(Settings.SmtpUser)
	smtpPasswordEncrypted, err := hex.DecodeString(Settings.SmtpPassword)
	if err != nil {
		return "", "", err
	}

	smtpUserDecrypted, err := decrypt([]byte(Settings.EncryptionKey), smtpUserEncrypted)
	smtpPasswordDecrypted, err := decrypt([]byte(Settings.EncryptionKey), smtpPasswordEncrypted)
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf("%s", smtpUserDecrypted), fmt.Sprintf("%s", smtpPasswordDecrypted), nil
}

func EncryptCredentials(smtpUserName string, smtpPassword string) (err error) {
	smtpUserNameByte, err := encrypt([]byte(Settings.EncryptionKey), []byte(smtpUserName))
	smtpPasswordByte, err := encrypt([]byte(Settings.EncryptionKey), []byte(smtpPassword))
	if err != nil {
		return err
	}

	Settings.SmtpUser = hex.EncodeToString(smtpUserNameByte)
	Settings.SmtpPassword = hex.EncodeToString(smtpPasswordByte)

	configWriter, err := os.Create(configFile)
	if err != nil {
		return err
	}
	if err := toml.NewEncoder(configWriter).Encode(&Settings); err != nil {
		return err
	}
	if err := configWriter.Close(); err != nil {
		return err
	}
	return nil
}
