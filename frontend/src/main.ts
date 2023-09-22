import './style.css';
import './app.css';
import { EventsOn } from "../wailsjs/runtime/runtime.js"; // the runtime for Wails2

import {AmOk, CallMe} from '../wailsjs/go/main/App';

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


// Setup the amOK function
window.amOk = function () {
    button1Element.disabled = true;
    // Call App.AmOk()
    try {
        AmOk()
            .then((result) => {
                // Update result with data back from App.AmOK()
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

// Setup the callMe function
window.callMe = function () {
    button2Element.disabled = true;
    // Call App.CallMe()
    try {
        CallMe()
            .then((result) => {
                // Update result with data back from App.CallMe()
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
    <div class="button-container">
        <button id="button1" class="button-xlarge margin-right green-button" onclick="amOk()">`+ button1Text + `</button>
        <button id="button2" class="button-xlarge yellow-button" onclick="callMe()">` + button2Text + `</button>
    </div>
    <div class="result" id="result"></div>
`;

let resultElement = document.getElementById("result");
let button1Element = <HTMLInputElement> document.getElementById("button1")
let button2Element = <HTMLInputElement> document.getElementById("button2")

declare global {
    interface Window {
        amOk: () => void;
        callMe: () => void;
    }
}

