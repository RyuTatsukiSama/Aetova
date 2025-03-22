

/*
This is a UI file (.ui.qml) that is intended to be edited in Qt Design Studio only.
It is supposed to be strictly declarative and only uses a subset of QML. If you edit
this file manually, you might introduce QML code that is not supported by Qt Design Studio.
Check out https://doc.qt.io/qtcreator/creator-quick-ui-forms.html for details on .ui.qml files.
*/
import QtQuick
import QtQuick.Controls

Item {
    id: root
    width: 1920
    height: 1080

    Rectangle {
        id: rectangle
        x: 0
        y: 0
        width: 1920
        height: 1080
        color: "#ffffff"

        Text {
            id: title
            x: 745
            y: 109
            text: qsTr("Aetova")
            font.pixelSize: 140
        }

        Button {
            id: button
            x: 189
            y: 329
            width: 314
            height: 507
            text: qsTr("Button")
            icon.height: 100
            icon.width: 100
            icon.cache: false
            icon.color: "#000000"
            icon.source: "../Dependencies/Components/examples/TransitionItem/images/svg/assistive-listening-systems.svg"
        }
    }
}
