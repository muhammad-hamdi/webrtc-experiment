// let p = new RTCPeerConnection(); // lol

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