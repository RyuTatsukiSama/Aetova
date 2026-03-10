// ─────────────────────────────────────────────
//  LAUNCH BUTTON – plug your game launch logic here
// ─────────────────────────────────────────────
const btnLaunch = document.getElementById('btn-launch');
var clientBridge = null;

new QWebChannel(qt.webChannelTransport, function (channel) {
  console.log("QWebChannel connected");
  clientBridge = channel.objects.clientBridge;

  // Listen qt signal
  if (clientBridge.monitoringSignal) {
    clientBridge.monitoringSignal.connect(function (dlPrc, dlSpeed, wrPrc, wrSpeed) {
      console.log("Download Pourcent: " + dlPrc + "% at speed " + dlSpeed + "kB/s");
      console.log("Write Pourcent: " + wrPrc + "% at speed " + wrSpeed + "kB/s");
    })
  }
});

btnLaunch.addEventListener('click', () => {
  btnLaunch.classList.add('loading');
  btnLaunch.textContent = 'Launching…';

  clientBridge.download();
});