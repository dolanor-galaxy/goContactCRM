import QtQuick 2.3
import QtQuick.Controls 1.1
import QtQuick.Layouts 1.0
//import "contact.js" as Person 


ApplicationWindow {
    visible: true
    title: "Contact Form"
    property int margin: 11
    width: 300
    height: 300

    GridLayout {
        id: personForm 
        anchors.margins: margin
        anchors.fill: parent
        columns:2

        Label {
            text: "Title:"
        }
        TextField {
            id: titleField
            placeholderText: "Title"
            Layout.fillWidth: true
            objectName: "title"
        }
        Label {
            text: "Family Name(s):"
        }
        TextField {
            id: familyNameField
            placeholderText: "Family Name(s)"
            Layout.fillWidth: true
            objectName: "familyName"
        }
        Label {
            text: "Given Name(s):"
        }
        TextField {
            id: givenNameField
            placeholderText: "Given Name(s)"
            Layout.fillWidth: true
            objectName: "givenName"
        }
        Label {
            text: "Additional Name(s):"
        }
        TextField {
            id: additionalNameField
            placeholderText: "Additional Name(s)"
            Layout.fillWidth: true
            objectName: "additionalName"
        }
        Label {
            text: "Nick Name(s):"
        }
        TextField {
            id: nickNameField
            placeholderText: "Nick Name(s)"
            Layout.fillWidth: true
            objectName: "nickName"
        }
        Label {
            text: "Role:"
        }
        TextField {
            id: roleField
            placeholderText: "Role"
            Layout.fillWidth: true
            objectName: "role"
        }
        Button {
            text: "Save"
            onClicked: person.savePerson(personForm)
        }
    }
}
