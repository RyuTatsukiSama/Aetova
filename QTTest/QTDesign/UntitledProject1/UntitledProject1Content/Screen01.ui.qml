

/*
This is a UI file (.ui.qml) that is intended to be edited in Qt Design Studio only.
It is supposed to be strictly declarative and only uses a subset of QML. If you edit
this file manually, you might introduce QML code that is not supported by Qt Design Studio.
Check out https://doc.qt.io/qtcreator/creator-quick-ui-forms.html for details on .ui.qml files.
*/
import QtQuick
import QtQuick.Controls
import UntitledProject1

Rectangle {
    id: rectangle
    width: 1024
    height: 720

    color: Constants.backgroundColor

    Button {
        id: button
        text: qsTr("Play")
        anchors.verticalCenter: parent.verticalCenter
        checkable: true
        property string gameName: "Journeep"
        anchors.horizontalCenter: parent.horizontalCenter

        Connections {
            target: button
            onClicked: animation.start()
        }

        Connections {
            target : button
            onClicked: gameLauncher.launchGame(button.gameName)
        }

        SequentialAnimation {
            id: animation

            ColorAnimation {
                id: colorAnimation1
                target: rectangle
                property: "color"
                to: "#2294c6"
                from: Constants.backgroundColor
            }

            ColorAnimation {
                id: colorAnimation2
                target: rectangle
                property: "color"
                to: Constants.backgroundColor
                from: "#2294c6"
            }
        }
    }

    Label {
        id: label
        x: 401
        y: 33
        width: 223
        height: 161
        text: qsTr("Journeep")
        horizontalAlignment: Text.AlignHCenter
        font.pointSize: 70
    }
    states: [
        State {
            name: "clicked"
            when: button.checked
        }
    ]
}
