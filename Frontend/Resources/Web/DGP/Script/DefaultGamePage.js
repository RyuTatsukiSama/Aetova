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
      const dlP = parseFloat(dlPrc).toFixed(1);
      const wrP = parseFloat(wrPrc).toFixed(1);
      const dlS = parseFloat(dlSpeed).toFixed(1);
      const wrS = parseFloat(wrSpeed).toFixed(1);

      document.getElementById("dl-bar").style.width = dlP + "%";
      document.getElementById("dl-info").textContent = dlP + "% — " + dlS + " kB/s";

      document.getElementById("wr-bar").style.width = wrP + "%";
      document.getElementById("wr-info").textContent = wrP + "% — " + wrS + " kB/s";
    })
  }

  if (clientBridge.bindFunctionToButton) {
    clientBridge.bindFunctionToButton.connect(function (text, name) {
      console.log("[JS/NewBind] Func name " + name + "\n");
      btnLaunch.textContent = text;
      btnLaunch.onclick = function () {
        clientBridge.CallFuncByName(name);
      }
    })
  }

  clientBridge.StartBinding();
});