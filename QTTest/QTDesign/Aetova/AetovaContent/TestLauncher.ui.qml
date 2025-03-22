

/*
This is a UI file (.ui.qml) that is intended to be edited in Qt Design Studio only.
It is supposed to be strictly declarative and only uses a subset of QML. If you edit
this file manually, you might introduce QML code that is not supported by Qt Design Studio.
Check out https://doc.qt.io/qtcreator/creator-quick-ui-forms.html for details on .ui.qml files.
*/
import QtQuick
import QtQuick.Controls
import Aetova

Rectangle {
    id: rectangle
    width: Constants.width
    height: Constants.height

    color: Constants.backgroundColor

    Text {
        text: qsTr("TestLauncher")
        font.pointSize: 70
        anchors.verticalCenterOffset: -386
        anchors.horizontalCenterOffset: 0
        anchors.centerIn: parent
        font.family: Constants.font.family
    }

    Button {
        id: button
        x: 760
        y: 300
        width: 400
        height: 70
        text: qsTr("Template Button")
        font.pointSize: 30
        property string path: "path\\Test"
        property string exeName: "Test"

        background: Rectangle {
            color: "#8b64a3"
            radius: 8
        }

        Connections {
            target: button
            onClicked: gameLauncher.launchGame(button.path, button.exeName)
        }
    }

    Button {
        id: journeep
        x: 760
        y: 445
        width: 400
        height: 70
        text: qsTr("Journeep")
        property string path: "Journeep"
        property string exeName: "Journeep"
        font.pointSize: 30
        Connections {
            target: journeep
            function onClicked() {
                gameLauncher.launchGame(journeep.path, journeep.exeName)
            }
        }
        background: Rectangle {
            color: "#8b64a3"
            radius: 8
        }
    }

    Button {
        id: hydreal
        x: 760
        y: 590
        width: 400
        height: 70
        text: qsTr("Hydreal")
        property string path: "Hydreal\\Exe"
        property string exeName: "Hydreal"
        font.pointSize: 30
        Connections {
            target: hydreal
            function onClicked() {
                gameLauncher.launchGame(hydreal.path, hydreal.exeName)
            }
        }
        background: Rectangle {
            color: "#8b64a3"
            radius: 8
        }
    }

    Button {
        id: glitchHunter
        x: 760
        y: 735
        width: 400
        height: 70
        text: qsTr("Glitch Hunter")
        property string path: "Glitch Hunter\\Glitch Hunter"
        font.pointSize: 30
        property string exeName: "Glitch Hunter"
        Connections {
            target: glitchHunter
            function onClicked() {
                gameLauncher.launchGame(glitchHunter.path, glitchHunter.exeName)
            }
        }
        background: Rectangle {
            color: "#8b64a3"
            radius: 8
        }
    }
}
