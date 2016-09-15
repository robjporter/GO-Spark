package spark

import (
    //"fmt"
    "errors"
    "strings"
    //"encoding/json"
    "github.com/roporter/go-libs/as"
    //"github.com/roporter/go-libs/jmespath"
)

func ListMessages(token string, roomid string, before string, beforemessageid string, max int) (string, string, string, error) {
    var listMessages transaction
    if roomid == "" {
        return "", "", "", errors.New("A room ID needs to be provided.")
    }
    listMessages.url = "https://api.ciscospark.com/v1/messages?"
    if roomid != "" {
        listMessages.url += "roomId=" + roomid
    }
    if before != "" {
        listMessages.url += "&before=" + urlEncode(before)
    }
    if beforemessageid != "" {
        listMessages.url += "&beforeMessage=" + beforemessageid
    }
    if max > 0 {
        listMessages.url += "&max=" + as.String(max)
    }
    listMessages.method = GET
    listMessages.body = ""
    listMessages.token = token
    listMessages, err := listMessages.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return listMessages.responseStatus, strings.Join(listMessages.responseHeaders[:],responseSeparator), listMessages.responseBody, nil
}
func CreateMessage(token string, roomid string, topersonid string, topersonemail string, text string, markdown string, files string) (string, string, string, error) {
    var createMessage transaction
    if roomid == "" && topersonid == "" && topersonemail == "" {
        return "", "", "", errors.New("A room ID, To Person ID or a To Person Email address needs to be provided when creating a new message.")
    }
    createMessage.url = "https://api.ciscospark.com/v1/messages"
    createMessage.method = POST
    body := "{"
    items := []string{}
    if roomid != "" { items = append(items, "\"roomId\":\"" + roomid + "\"")}
    if topersonid != "" { items = append(items, "\"toPersonId\":\"" + topersonid + "\"")}
    if topersonemail != "" { items = append(items, "\"toPersonEmail\":\"" + topersonemail + "\"")}
    if text != "" { items = append(items, "\"text\":\"" + text + "\"")}
    if markdown != "" { items = append(items, "\"markdown\":\"" + markdown + "\"")}
    if files != "" { items = append(items, "\"files\":\"" + files + "\"")}
    for x := range items {
        var toAdd = ""
        if(x < len(items)-1) {
            toAdd = ","
        }
        body += items[x] + toAdd
    }
    createMessage.body = body + "}"
    createMessage.token = token
    createMessage, err := createMessage.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return createMessage.responseStatus, strings.Join(createMessage.responseHeaders[:],responseSeparator), createMessage.responseBody, nil
}
func GetMessageDetail(token string, messageid string) (string, string, string, error) {
    var getMessage transaction
    if messageid == "" {
        return "", "", "", errors.New("A message ID needs to be provided.")
    }
    getMessage.url = "https://api.ciscospark.com/v1/messages/" + messageid
    getMessage.method = GET
    getMessage.body = ""
    getMessage.token = token
    getMessage, err := getMessage.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return getMessage.responseStatus, strings.Join(getMessage.responseHeaders[:],responseSeparator), getMessage.responseBody, nil
}
func DeleteMessage(token string, messageid string) (string, string, string, error) {
    var deleteMessage transaction
    if messageid == "" {
        return "", "", "", errors.New("A message ID needs to be provided.")
    }
    deleteMessage.url = "https://api.ciscospark.com/v1/messages/" + messageid
    deleteMessage.method = DELETE
    deleteMessage.body = ""
    deleteMessage.token = token
    deleteMessage, err := deleteMessage.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return deleteMessage.responseStatus, strings.Join(deleteMessage.responseHeaders[:],responseSeparator), deleteMessage.responseBody, nil
}
