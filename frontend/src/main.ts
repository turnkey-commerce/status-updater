import './style.css';
import './app.css';

import {AmOk} from '../wailsjs/go/main/App';
import {CallMe} from '../wailsjs/go/main/App';

let timeout = 7000;

// Setup the amOK function
window.amOk = function () {
    greenButtonElement.disabled = true;
    // Call App.AmOk()
    try {
        AmOk()
            .then((result) => {
                // Update result with data back from App.AmOK()
                resultElement!.innerText = result;
                setTimeout(() => {
                    resultElement!.innerText = ""
                    greenButtonElement.disabled = false;
                },
                timeout);
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
    yellowButtonElement.disabled = true;
    // Call App.CallMe()
    try {
        CallMe()
            .then((result) => {
                // Update result with data back from App.CallMe()
                resultElement!.innerText = result;
                setTimeout(() => {
                    resultElement!.innerText = ""
                    yellowButtonElement.disabled = false
                },
                timeout);
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
        <button id="greenButton" class="button-xlarge margin-right green-button" onclick="amOk()">I'm OK</button>
        <button id="yellowButton" class="button-xlarge yellow-button" onclick="callMe()">Call Me</button>
    </div>
    <div class="result" id="result"></div>
`;

let resultElement = document.getElementById("result");
let greenButtonElement = <HTMLInputElement> document.getElementById("greenButton")
let yellowButtonElement = <HTMLInputElement> document.getElementById("yellowButton")

declare global {
    interface Window {
        amOk: () => void;
        callMe: () => void;
    }
}

