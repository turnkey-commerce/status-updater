# Status Updater

The Status Updater provides a simple interface for a person
with limited vision to routinely update their status to
remote caregivers. They can choose "I'm OK" or "Call Me"
and the application will automatically send an email to one
or more caregivers.

The two buttons presented by the app are very large with
large fonts with high contrast (black on yellow) so that
they are easy to see.

The app should run as a desktop application of platforms
as supported by the [Wails](https://wails.io/) framework
(currently supporting Windows, MacOS, and Linux).

The simple and accessible interface looks as follows when
the app is started from the desktop:

![App at Startup](https://github.com/turnkey-commerce/status-updater/blob/master/screenshots/status-updater-screen1.png)

Once the large "I'm OK" button is clicked
then the button is temporarily disabled
and a status message will show that the
email was successfully sent:

![App after button pressed](https://github.com/turnkey-commerce/status-updater/blob/master/screenshots/status-updater-screen2.png)

## Running the Application

Currently the application is available as a binary
executable for the Windows 64-bit environment. You can
download and run the program as follows:

1. Download the latest executable from the
[Releases](https://github.com/turnkey-commerce/status-updater/releases).

2. Extract the zip to a folder.

3. Copy example.config.toml file to config.toml.

4. Edit the config.toml file for the email SMTP
settings, the recipients email list, and the name of
the user that will use the program. Optionally you can
edit the button labels that will be seen by the user.

5. Start the program by clicking on the status-updater.exe
file. Use the Configuration -> Configure menu to set the
SMTP User and Password for the email settings. This will
write the encrypted settings back to the config.toml so that
they are not visible as plain text.

6. The program can then be tested by clicking on the
status-updater.exe file.  If all works as expected
it can be distributed to the user, being sure to
keep it with the config.toml file.

7. For better ease of use for the end user you can
right-mouse drag the status-updater.exe to the user's
desktop and create a shortcut to the program.

## Configuration File

The config.toml configuration file needs to be present with the
executable file for the configuration of the SMTP server, the
email recipients, and the name of the end user. Additionally
the button labels can be changed for alternate languages or to
send different messages than the default "I'm OK" and "Call Me".

The contents of the configuration file are as follows and
are available in the example.config.toml.

Also, it is recommended to change the Encryption key to a
different one before the program is first run. The website
[Random Key Generator](https://acte.ltd/utils/randomkeygen)
can be used to randomly generate it using the "Encryption key
256" option for a 32 byte key.

```text
# SMTP credentials for sending email notifications
SmtpUser     = ""  # Will be set as an encrypted string by the Configuration menu.
SmtpPassword = ""  # Will be set as an encrypted string by the Configuration menu.
EncryptionKey = "p04lCUCXBjDIlpiN1dIjRauOghtmL8f1"  # 32 bytes https://acte.ltd/utils/randomkeygen
SmtpServer   = "smtp.gmail.com"
SmtpPort     = "587"
FromEmail    = "sender@example.org"

# The comma-delimited list of email recipients to receive the updates.
Recipients = ["JohnDoe@example.com", "JaneDoe@example.com"]

# The name of the person that will use the program and send the updates.
Name = "Bill"

# The labels on the buttons. Can be changed to support alternate languages or different purposes.
# Try to keep the length short because of the button size.
Button1Label = "I'm OK"
Button2Label = "Call Me"
```

## Development

If you wish to modify the way the application works you can use the following steps to set up a development environment.

1. Clone the code from the Github repository:
<https://github.com/turnkey-commerce/status-updater.git>

2. Install Wails as described in the Wails documentation:
[Wails Installation](https://wails.io/docs/gettingstarted/installation)

3. Copy the file example.config.toml to config.toml and
put in the correct settings for your environment.

4. Modify the program as desired.

5. Use the following command to start the iterative
development environment:

    ```text
        wails dev
    ```

6. An application window will pop up and it will
automatically update on code changes.

7. Once you are satisfied with the changes and test that
it works then build the application as follows:

    ```text
        wails build
    ```

8. The executable will be built in the /build/bin subfolder.
Be sure and also copy the config.toml to distribute with the
executable as they must be in the same directory.
