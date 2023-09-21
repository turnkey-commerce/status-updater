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
