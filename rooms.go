package spark

import (
    "errors"
    "strings"
    "encoding/json"
    "github.com/roporter/go-libs/as"
    "github.com/roporter/go-libs/jmespath"
)

func CreateRoom(token string, name string, teamid string) (string, string, string,error) {
    var createRoom transaction
    if name == "" {
        return "", "", "", errors.New("A name needs to be given for the new room")
    }
    createRoom.url = "https://api.ciscospark.com/v1/rooms"
    createRoom.method = POST
    body := "{\"title\":\""+name+"\""
    if teamid != "" {
        body = body + ",\"teamId\":\"" + teamid + "\""
    }
    body = body + "}"
    createRoom.body = body
    createRoom.token = token
    createRoom,err := createRoom.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return createRoom.responseStatus, strings.Join(createRoom.responseHeaders[:],responseSeparator), createRoom.responseBody, nil
}
func FindRoomID(token string, name string) (string, string, string, error) {
    status, headers, body, err := ListRooms(token, "", 0, "")
    if err != nil {
        return "", "", "", err
    }
    newBody := "{\"ID\":\"Not Found\"}"
    var data interface{}
    expression := "items.length(@)"
    json.Unmarshal([]byte(body), &data)
    result, _ := jmespath.Search(expression, data)
    for x := 0; x < as.Int(result); x++ {
        expression := "items[" + as.String(x) + "].title"
        result2, _ := jmespath.Search(expression, data)
        if strings.TrimSpace(as.String(result2)) == strings.TrimSpace(name) {
            expression = "items[" + as.String(x) + "].id"
            result2, _ = jmespath.Search(expression, data)
            newBody = "{\"id\":\"" + as.String(result2) + "\"}"
        }
    }
    return status, headers, newBody, nil
}

func ListRooms(token string, teamid string, max int, typed string) (string, string, string, error) {
    var listRooms transaction
    listRooms.url = "https://api.ciscospark.com/v1/rooms"
    listRooms.method = GET
    items := []string{}
    body := "{"
    if teamid != "" { items = append(items, "\"teamId\":\"" + teamid + "\"")}
    if max > 0 { items = append(items, "\"max\":" + as.String(max))}
    if typed != "" { items = append(items, "\"type\":\"" + typed + "\"")}
    for x := range items {
        var toAdd = ""
        if(x < len(items)-1) {
            toAdd = ","
        }
        body += items[x] + toAdd
    }
    listRooms.body = body + "}"
    listRooms.token = token
    listRooms, err := listRooms.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return listRooms.responseStatus, strings.Join(listRooms.responseHeaders[:],responseSeparator), listRooms.responseBody, nil
}
func GetRoomDetails(token string, roomid string) (string, string, string,error) {
    var roomDetails transaction
    if roomid == "" {
        return "", "", "", errors.New("A room ID needs to be provided.")
    }
    roomDetails.url = "https://api.ciscospark.com/v1/rooms/" + roomid
    roomDetails.method = GET
    roomDetails.body = ""
    roomDetails.token = token
    roomDetails, err := roomDetails.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return roomDetails.responseStatus, strings.Join(roomDetails.responseHeaders[:],responseSeparator), roomDetails.responseBody, nil
}
func UpdateRoomDetails(token string, roomid string, title string) (string, string, string,error) {
    var updateRoom transaction
    if roomid == "" {
        return "", "", "", errors.New("A room ID needs to be provided.")
    }
    if title == "" {
        return "", "", "", errors.New("A new name needs to be provided for the room.")
    }
    updateRoom.url = "https://api.ciscospark.com/v1/rooms/" + roomid
    updateRoom.method = PUT
    updateRoom.body = "{\"title\":\""+title+"\"}"
    updateRoom.token = token
    updateRoom, err := updateRoom.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return updateRoom.responseStatus, strings.Join(updateRoom.responseHeaders[:],responseSeparator), updateRoom.responseBody, nil
}
func DeleteRoom(token string, roomid string) (string, string, string,error) {
    var deleteRoom transaction
    if roomid == "" {
        return "", "", "", errors.New("A room ID needs to be provided.")
    }
    deleteRoom.url = "https://api.ciscospark.com/v1/rooms/" + roomid
    deleteRoom.method = DELETE
    deleteRoom.body = ""
    deleteRoom.token = token
    deleteRoom, err := deleteRoom.doTransaction()
    if err != nil {
        return "", "", "", err
    }
    return deleteRoom.responseStatus, strings.Join(deleteRoom.responseHeaders[:],responseSeparator), deleteRoom.responseBody, nil
}
