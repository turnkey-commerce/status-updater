import './style.css';
import './app.css';
import { EventsOn } from "../wailsjs/runtime/runtime.js"; // the runtime for Wails2

import {Button1Action, Button2Action} from '../wailsjs/go/main/App';

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
    // Popup Configuration window.
    configurationElement!.innerText = "Configure";
    setTimeout(() => {
        configurationElement!.innerText = ""
    },
    resultShowTimeMs);
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

document.querySelector('#app')!.innerHTML = `
    <div class="configuration-container" id="configuration">
    </div>
    <div class="button-container">
        <button id="button1" class="button-xlarge margin-right green-button" onclick="button1Action()">`+ button1Text + `</button>
        <button id="button2" class="button-xlarge yellow-button" onclick="button2Action()">` + button2Text + `</button>
    </div>
    <div class="result" id="result"></div>
`;

let resultElement = document.getElementById("result");
let button1Element = <HTMLInputElement> document.getElementById("button1")
let button2Element = <HTMLInputElement> document.getElementById("button2")
let configurationElement = document.getElementById("configuration");

declare global {
    interface Window {
        button1Action: () => void;
        button2Action: () => void;
    }
}

