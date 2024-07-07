import './style.css';
import './app.css';
import { EventsOn } from "../wailsjs/runtime/runtime.js"; // the runtime for Wails2

import {Button1Action, Button2Action, SaveAction} from '../wailsjs/go/main/App';

let resultShowTimeMs = 7000;

let button1Text = "I'm OK"
let button2Text = "Call Me"

// Listen for events from the backend.
//
EventsOn("setButtonText", (button1, button2) => {
    button1Element.innerText = button1
    button2Element.innerText = button2
    console.log(button1, button2)
});

EventsOn("configure", () => {
    // Create Configuration window.
    configurationElement!.innerHTML = `
        <div class="row">
            <div class="col-25">
                <label for="userName">SMTP Username</label>
            </div>
            <div class="col-75">
                <input type="text" id="userName" name="userName">
            </div>
        </div>
        <div class="row">
            <div class="col-25">
                <label for="password">SMTP Password</label>
            </div>
            <div class="col-75">
                <input type="password" id="password" name="password">
            </div>
        </div>
        <div class="row">
            <button class="button-small" id="saveButton" onclick="saveAction()">Save</button>
        </div>
    `;
    saveElement = <HTMLInputElement> document.getElementById("saveButton");

    buttonsElement.style.display = "none";
    saveElement.disabled = false;
});

// Setup the button1Action function
window.button1Action = function () {
    button1Element.disabled = true;
    // Call App.Button1Action()
    try {
        Button1Action()
            .then((result) => {
                // Update result with data back from App.Button1Action()
                resultElement!.innerText = result;
                setTimeout(() => {
                    resultElement!.innerText = ""
                    button1Element.disabled = false;
                },
                resultShowTimeMs);
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

// Setup the button2Action function
window.button2Action = function () {
    button2Element.disabled = true;
    // Call App.Button2Action()
    try {
        Button2Action()
            .then((result) => {
                // Update result with data back from App.Button2Action()
                resultElement!.innerText = result;
                setTimeout(() => {
                    resultElement!.innerText = ""
                    button2Element.disabled = false
                },
                resultShowTimeMs);
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

// Setup the Configuration Save function
window.saveAction = function () {
    saveElement.disabled = true;
    let passwordElement = <HTMLInputElement> document.getElementById("password");
    let password = passwordElement.value;
    let userNameElement = <HTMLInputElement> document.getElementById("userName");
    let userName = userNameElement.value;
    try {
        SaveAction(userName, password)
            .then((result) => {
                // Update result with data back from App.SaveAction()
                resultElement!.innerText = result;
                setTimeout(() => {
                    resultElement!.innerText = "";
                    saveElement.disabled = false;
                    configurationElement!.innerHTML = "";
                    buttonsElement!.style.display = "block";
                },
                resultShowTimeMs);
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
};

document.querySelector('#app')!.innerHTML = `
    <div class="configuration" id="configuration">
    </div>
    <div class="button-container" id="buttons">
        <button id="button1" class="button-xlarge margin-right green-button" onclick="button1Action()">`+ button1Text + `</button>
        <button id="button2" class="button-xlarge yellow-button" onclick="button2Action()">` + button2Text + `</button>
    </div>
    <div class="result" id="result"></div>
`;

let resultElement = document.getElementById("result");
let button1Element = <HTMLInputElement> document.getElementById("button1")
let button2Element = <HTMLInputElement> document.getElementById("button2")
let buttonsElement = <HTMLInputElement> document.getElementById("buttons")
let configurationElement = document.getElementById("configuration");
let saveElement = <HTMLInputElement> document.getElementById("saveButton")

declare global {
    interface Window {
        button1Action: () => void;
        button2Action: () => void;
        saveAction: () => void;
    }
}

