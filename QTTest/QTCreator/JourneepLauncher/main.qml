import QtQuick 2.15
import QtQuick.Controls 2.15

import UntitledProject1 1.0 // Si Qt DS a généré un module

ApplicationWindow {
    visible: true
    width: 1280
    height: 720

    title: "Mon Application Qt"

    // Charge App.qml (ou un autre fichier QML principal de Qt DS)
    App {
        anchors.fill: parent
    }
}
