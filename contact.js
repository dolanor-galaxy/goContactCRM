
function savePerson() {
    console.log("It fucking works!!", personForm.children[1].objectName);
    var obj = new Object();
    
    var pFormLength = personForm.children.length;
    console.log("pFormLength: ", pFormLength);

    obj[personForm.children[1].objectName] = personForm.children[1].text
    obj[personForm.children[3].objectName] = personForm.children[3].text
    obj[personForm.children[5].objectName] = personForm.children[5].text
    obj[personForm.children[7].objectName] = personForm.children[7].text
    obj[personForm.children[9].objectName] = personForm.children[9].text
    obj[personForm.children[11].objectName] = personForm.children[11].text

        /*
    for(var i=0;i<=pFormLength;i++) {

        console.log(i);

        obj.personForm.children[i].objectName = personForm.children[i].text
        obj[personForm.children[i].objectName] = personForm.children[i].text

            /*
        if(personForm.children[i].objectName !== undefined) {
            obj[personForm.children[i].objectName] = personForm.children[i].text
        } else {
            console.log("objectName is really really undefined!");
        }
    }
    */
    for (var property in obj) {
        console.log(property);
        console.log(obj[property]);
    }
    person.savePerson(personForm);



    //personMap.set(personForm.children.name)
    /*
    console.log(personForm.children.type)
    console.log(personForm.children[1].name)
    console.log(personForm.children[1].objectName)
    console.log(personForm.children.length)
    var keys = Object.keys(personForm.children)
    console.log(keys)
    /*
    person.setTitle(titleField.text)
    person.setFamilyName(familyNameField.text)
    person.setGivenName(givenNameField.text)
    person.setAdditionalName(additionalNameField.text)
    person.setNickNames(nickNameField.text)
    person.setRole(roleField.text)
    // Number of entities in the personForm
    console.log(personForm.children.length)
    console.log(personForm.children[1].text)
    console.log(function() {
        return "Hello World"
    })
    */
}
