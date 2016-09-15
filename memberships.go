package spark

import (
    "errors"
    "strings"
    "github.com/roporter/go-libs/as"
)

func UpdateMembership(token string, membershipid string, ismod bool) (string, string, string, error) {
    var updateMembership transaction
    if membershipid == "" {
        return "", "", "", errors.New("A membership ID needs to be provided.")
    }
    updateMembership.url = "https://api.ciscospark.com/v1/memberships/" + membershipid
    updateMembership.method = DELETE
    updateMembership.body = "{\"isModerator\":" + as.String(ismod) + "}"
    updateMembership.token = token
    updateMembership, err := updateMembership.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return updateMembership.responseStatus, strings.Join(updateMembership.responseHeaders[:],responseSeparator), updateMembership.responseBody, nil
}
func DeleteMembership(token string, membershipid string) (string, string, string, error) {
    var deleteMembership transaction
    if membershipid == "" {
        return "", "", "", errors.New("A membership ID needs to be provided.")
    }
    deleteMembership.url = "https://api.ciscospark.com/v1/memberships/" + membershipid
    deleteMembership.method = DELETE
    deleteMembership.body = ""
    deleteMembership.token = token
    deleteMembership, err := deleteMembership.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return deleteMembership.responseStatus, strings.Join(deleteMembership.responseHeaders[:],responseSeparator), deleteMembership.responseBody, nil
}
func CreateMembership(token string, roomid string, personid string, personemail string, ismod bool) (string, string, string, error) {
    var createMembership transaction
    createMembership.url = "https://api.ciscospark.com/v1/memberships"
    createMembership.method = POST
    if roomid == "" {
        return "", "", "", errors.New("room ID must be set for creating any new memberships.")
    }
    body := "{"
    items := []string{}
    items = append(items, "\"isModerator\":" + as.String(ismod))
    if roomid != "" { items = append(items, "\"roomId\":\"" + roomid + "\"")}
    if personid != "" { items = append(items, "\"personId\":\"" + personid + "\"")}
    if personemail != "" { items = append(items, "\"personEmail\":\"" + personemail + "\"")}
    for x := range items {
        var toAdd = ""
        if(x < len(items)-1) {
            toAdd = ","
        }
        body += items[x] + toAdd
    }
    createMembership.body = body + "}"
    createMembership.token = token
    createMembership, err := createMembership.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return createMembership.responseStatus, strings.Join(createMembership.responseHeaders[:],responseSeparator), createMembership.responseBody, nil
}

func GetMembershipDetail(token string, membershipid string) (string, string, string, error) {
    var getMembership transaction
    if membershipid == "" {
        return "", "", "", errors.New("A membership ID is required to query the membership status.")
    }
    getMembership.url = "https://api.ciscospark.com/v1/memberships/"+membershipid
    getMembership.method = GET
    getMembership.body = ""
    getMembership.token = token
    getMembership, err := getMembership.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return getMembership.responseStatus, strings.Join(getMembership.responseHeaders[:],responseSeparator), getMembership.responseBody, nil
}

func ListMemberships(token string, roomid string, personid string, personemail string, max int) (string, string, string, error) {
    var listMemberships transaction
    listMemberships.url = "https://api.ciscospark.com/v1/memberships?"
    listMemberships.method = GET
    first := true
    if roomid != "" {
        listMemberships.url += "roomId=" + roomid
        first = false
    }
    if personid != "" {
        if first {
            listMemberships.url += "personId=" + personid
            first = false
        } else {
            listMemberships.url += "&personId=" + personid
        }
    }
    if personemail != "" {
        if first {
            listMemberships.url += "personEmail=" + personemail
            first = false
        } else {
            listMemberships.url += "&personEmail=" + personemail
        }
    }
    listMemberships.token = token
    listMemberships, err := listMemberships.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return listMemberships.responseStatus, strings.Join(listMemberships.responseHeaders[:],responseSeparator), listMemberships.responseBody, nil
}
