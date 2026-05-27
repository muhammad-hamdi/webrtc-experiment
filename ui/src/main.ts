// let p = new RTCPeerConnection(); // lol

function setHandlers() {
    let createBtn = document.getElementById("create-btn")
    let joinBtn = document.getElementById("join-btn")
    let joinInput = document.getElementById("join-input") as HTMLInputElement

    createBtn?.addEventListener("click", () => {
        // create room message/logic
    });

    joinBtn?.addEventListener("click", () => {
        if(joinInput.value.length) {
            // join room message/logic
        }
    });
}

async function init() {
    let stream = await navigator.mediaDevices.getUserMedia({
        audio:true,
        video:true
    })

    let localVideo = (document.getElementById("localVideo") as HTMLVideoElement);
    localVideo.srcObject = stream;
    localVideo.play();
}

init();