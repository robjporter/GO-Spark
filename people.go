package spark

import (
    "errors"
    "strings"
)

func GetMyDetail(token string) (string, string, string,error) {
    var myPeople transaction
    myPeople.url = "https://api.ciscospark.com/v1/people/me"
    myPeople.method = GET
    myPeople.body = ""
    myPeople.token = token
    myPeople, err := myPeople.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return myPeople.responseStatus, strings.Join(myPeople.responseHeaders[:],responseSeparator), myPeople.responseBody, nil
}

func GetPeopleDetail(token string, peopleid string) (string, string, string,error) {
    var peopleDetail transaction
    peopleDetail.url = "https://api.ciscospark.com/v1/people/"+peopleid
    peopleDetail.method = GET
    peopleDetail.body = ""
    peopleDetail.token = token
    peopleDetail, err := peopleDetail.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return peopleDetail.responseStatus, strings.Join(peopleDetail.responseHeaders[:],responseSeparator), peopleDetail.responseBody, nil
}

func ListPeople(token string, email string, displayname string, max int) (string, string, string,error) {
    var listPeople transaction
    if email == "" && displayname == "" {
        return "", "", "", errors.New("A email or a display name needs to be provided.")
    }
    listPeople.url = "https://api.ciscospark.com/v1/people?"
    listPeople.method = GET
    first := true
    if email != "" {
        listPeople.url += "email="+email
        first = false
    }
    if displayname != "" {
        if first {
            listPeople.url += "displayName="+urlEncode(displayname)
        } else {
            listPeople.url += "&displayName="+urlEncode(displayname)
        }
    }
    listPeople.body = ""
    listPeople.token = token
    listPeople, err := listPeople.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return listPeople.responseStatus, strings.Join(listPeople.responseHeaders[:],responseSeparator), listPeople.responseBody, nil
}
