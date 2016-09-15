package main

import (
    "os"
    "fmt"
    "time"
    "github.com/roporter/go-spark"
    "github.com/roporter/go-libs/as"
)

var lastID, roomid, peopleid  string

func main() {
    var demoFull = false
    var demoListRooms = false
    var demoGetRoomDetails = false
    var demoNewRoom = false
    var demoListPeople = false
    var demoMyPeople = false
    var demoListMembership = false
    var demoCreateMembership = false
    var demoMembershipDetail = true
    var demoListMessages = false
    var demoCreateMessage = false
    var demoGetMessageDetail = false
    var demoDeleteMessage = false

    fmt.Println("Initialising....")
    token := os.Getenv("SPARK_TOKEN")
    if token == "" {
        token = "<TOKEN>"
    }
    fmt.Println("Token configured: ", token)

    if demoFull {
        demoFullAction(token)
    }
    if demoListRooms {
        demoListRoomsAction(token)
    }
    if demoGetRoomDetails {
        demoGetRoomDetailsAction(token)
    }
    if demoNewRoom {
        demoNewRoomAction(token)
    }
    if demoListPeople {
        demoListPeopleAction(token)
    }
    if demoMyPeople {
        demoMyPeopleAction(token)
    }
    if demoListMembership {
        demoListMembershipAction(token)
    }
    if demoCreateMembership {
        demoCreateMembershipAction(token)
    }
    if demoMembershipDetail {
        demoMembershipDetailAction(token)
    }
    if demoListMessages {
        demoListMessagesAction(token)
    }
    if demoCreateMessage {
        demoCreateMessageAction(token)
    }
    if demoGetMessageDetail {
        demoGetMessageDetailAction(token)
    }
    if demoDeleteMessage {
        demoDeleteMessageAction(token)
    }
}
func demoDeleteMessageAction(token string) {
    messageid := "<MESSAGEID>"
    status, headers, body, err := spark.DeleteMessage(token,messageid)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    fmt.Println("BODY: ",body)
    if status == "204 No Content" {
        fmt.Println("Message deleted successfully from room.")
    } else {
        fmt.Println("Message deletion failed.")
    }
}
func demoGetMessageDetailAction(token string) {
    fmt.Println("GET MESSAGE DETAIL =================================================")
    messageid := "<MESSAGEID>"
    status, headers, body, err := spark.GetMessageDetail(token,messageid)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    fmt.Println("ID: ",spark.GetParameterByDetails(body,"id",0))
    fmt.Println("ROOMID: ",spark.GetParameterByDetails(body,"roomId",0))
    fmt.Println("ROOMTYPE: ",spark.GetParameterByDetails(body,"roomType",0))
    fmt.Println("TEXT: ",spark.GetParameterByDetails(body,"text",0))
    fmt.Println("PERSONID: ",spark.GetParameterByDetails(body,"personId",0))
    fmt.Println("PERSONEMAIL: ",spark.GetParameterByDetails(body,"personEmail",0))
    fmt.Println("CREATED: ",spark.GetParameterByDetails(body,"created",0))
}
func demoCreateMessageAction(token string) {
    fmt.Println("CREATE MESSAGE =====================================================")
    roomid := "<ROOMID>"
    topersonid := "<PERSONID>"
    topersonemail := "<PERSONEMAIL>"
    status, headers, _, err := spark.CreateMessage(token,roomid,"","","","**TEST MESSAGE**","") // token, roomid, personid, personemail, plain text message, markdown message, files
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    if status == "200 OK" {
        fmt.Println("Message posted successfully to room.")
    } else {
        fmt.Println("Message posting failed.")
    }
}
func demoListMessagesAction(token string) {
    fmt.Println("LIST MESSAGES ======================================================")
    roomid := "<ROOMID>"
    messageid := "<MESSAGEID>"
    timestamp := spark.GetISO8601Timestamp("2016","07","15","12","00")
    //status, headers, body, err := spark.ListMessages(token,roomid,"","",0)
    status, headers, body, err := spark.ListMessages(token,roomid,timestamp,"",4) // token, roomid, before timestamp, before messageid, max items
    //status, headers, body, err := spark.ListMessages(token,roomid,"",messageid,4)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)


    for x := 0; x < spark.GetSize(body); x++ {
        fmt.Println("ID: ",spark.GetParameterByDetails(body,"id",x))
        fmt.Println("ROOMID: ",spark.GetParameterByDetails(body,"roomId",x))
        fmt.Println("ROOMTYPE: ",spark.GetParameterByDetails(body,"roomType",x))
        fmt.Println("TEXT: ",spark.GetParameterByDetails(body,"text",x))
        fmt.Println("PERSONID: ",spark.GetParameterByDetails(body,"personId",x))
        fmt.Println("PERSONEMAIL: ",spark.GetParameterByDetails(body,"personEmail",x))
        fmt.Println("CREATED: ",spark.GetParameterByDetails(body,"created",x))
        fmt.Println("")
    }
}
func demoMembershipDetailAction(token string) {
    fmt.Println("GET MEMBERSHIP ======================================================")
    status, headers, body, err := spark.GetMembershipDetail(token,"<MEMBERSHIPID>")
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)

    fmt.Println("ID: ",spark.GetParameterByDetails(body,"id",0))
    fmt.Println("ROOMID: ",spark.GetParameterByDetails(body,"roomId",0))
    fmt.Println("PERSONID: ",spark.GetParameterByDetails(body,"personId",0))
    fmt.Println("PERSONEMAIL: ",spark.GetParameterByDetails(body,"personEmail",0))
    fmt.Println("PERSONDISPLAYNAME: ",spark.GetParameterByDetails(body,"personDisplayName",0))
    fmt.Println("ISMODERATOR: ",spark.GetParameterByDetails(body,"isModerator",0))
    fmt.Println("ISMONITOR: ",spark.GetParameterByDetails(body,"isMonitor",0))
    fmt.Println("CREATED: ",spark.GetParameterByDetails(body,"created",0))
}
func demoFullAction(token string) {
    fmt.Println("NEW ROOM ============================================================")
    status, headers, body, err := spark.CreateRoom(token,"<ROOMNAME>","<TEAMID>")
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    roomid = spark.GetParameterByDetails(body,"id",0)

    time.Sleep(2 * time.Second)

    //personid := "<PERSONID>" //John
    //status, headers, body, err = spark.CreateMembership(token,roomid,personid,"",false)
    status, headers, body, err = spark.CreateMembership(token,roomid,"<PERSONID>","<PERSONEMAIL>",false)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    var membershipid = spark.GetParameterByDetails(body,"id",0)
    if status == "200 OK" {
        fmt.Println("Addition of user to group succeeded.")
    } else {
        fmt.Println("Addition of user to group failed.")
    }

    status, headers, body, err = spark.CreateMessage(token,roomid,"","","","**TEST**","")
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    messageid := spark.GetParameterByDetails(body,"id",0)
    if status == "200 OK" {
        fmt.Println("Message posted successfully to room.")
    } else {
        fmt.Println("Message posting failed.")
    }

    time.Sleep(10 * time.Second)

    status, headers, body, err = spark.DeleteMessage(token,messageid)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    if status == "204 No Content" {
        fmt.Println("Message deleted successfully from room.")
    } else {
        fmt.Println("Message deletion failed.")
    }

    status, headers, body, err = spark.DeleteMembership(token,membershipid)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    if status == "204 No Content" {
        fmt.Println("Deletion of user from group succeeded.")
    } else {
        fmt.Println("Deletion of user from group failed.")
    }

    fmt.Println("DELETE ROOM =========================================================")
    status, headers, _, err = spark.DeleteRoom(token,as.String(roomid))
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    if status == "204 No Content" {
        fmt.Println("DELETED: SUCCESSFULLY")
    } else {
        fmt.Println("DELETED: FAILED")
    }
}
func demoListRoomsAction(token string) {
    fmt.Println("LIST ROOMS ==========================================================")
    status, headers, body, err := spark.ListRooms(token,"",0,"")
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)

    fmt.Println("===============")
    for x := 0; x < spark.GetSize(body); x++ {
        lastID = spark.GetParameterByDetails(body,"id",x)
        fmt.Println("TITLE: ",spark.GetParameterByDetails(body,"title",x))
        fmt.Println("ID: ",lastID)
        fmt.Println("TYPE: ",spark.GetParameterByDetails(body,"type",x))
        fmt.Println("LOCKED: ",spark.GetParameterByDetails(body,"isLocked",x))
        fmt.Println("LASTACTIVITY: ",spark.GetParameterByDetails(body,"lastActivity",x))
        fmt.Println("CREATED: ",spark.GetParameterByDetails(body,"created",x))
    }
}
func demoGetRoomDetailsAction(token string) {
    fmt.Println("GET ROOM DETAILS ====================================================")
    status, headers, body, err := spark.GetRoomDetails(token,lastID)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)

    fmt.Println("TITLE: ",spark.GetParameterByDetails(body,"title",0))
    fmt.Println("ID: ",lastID)
    fmt.Println("TYPE: ",spark.GetParameterByDetails(body,"type",0))
    fmt.Println("LOCKED: ",spark.GetParameterByDetails(body,"isLocked",0))
    fmt.Println("LASTACTIVITY: ",spark.GetParameterByDetails(body,"lastActivity",0))
    fmt.Println("CREATED: ",spark.GetParameterByDetails(body,"created",0))
}
func demoNewRoomAction(token string) {
    fmt.Println("NEW ROOM ============================================================")
    status, headers, body, err := spark.CreateRoom(token,"<ROOMNAME>","")
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    roomid = spark.GetParameterByDetails(body,"id",0)
    fmt.Println("ID: ",roomid)

    time.Sleep(2 * time.Second)

    fmt.Println("UPDATE ROOM ==========================================================")
    status, headers, body, err = spark.UpdateRoomDetails(token,as.String(roomid),"<NEWROOMNAME>")
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    if status == "200 OK" {
        fmt.Println("UDPATED: SUCCESSFULLY")
    } else {
        fmt.Println("UDPATED: FAILED")
    }

    time.Sleep(2 * time.Second)

    fmt.Println("FIND ROOM ============================================================")
    status, headers, body, err = spark.FindRoomID(token,"<ROOMNAME>")
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    fmt.Println("ID: ",spark.GetParameterByDetails(body,"id",0))

    time.Sleep(2 * time.Second)

    fmt.Println("DELETE ROOM =========================================================")
    status, headers, _, err = spark.DeleteRoom(token,as.String(roomid))
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    if status == "204 No Content" {
        fmt.Println("DELETED: SUCCESSFULLY")
    } else {
        fmt.Println("DELETED: FAILED")
    }

    time.Sleep(2 * time.Second)
}
func demoListPeopleAction(token string) {
    fmt.Println("LIST PEOPLE ==========================================================")
    status, headers, body, err := spark.ListPeople(token,"<PERSONEMAIL>","<PERSONNAME>",0)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)

    fmt.Println("===============")
    for x := 0; x < spark.GetSize(body); x++ {
        fmt.Println("DISPLAYNAME: ",spark.GetParameterByDetails(body,"displayName",x))
        fmt.Println("ID: ",spark.GetParameterByDetails(body,"id",x))
        fmt.Println("EMAILS: ",spark.GetParameterByDetails(body,"emails",x))
        fmt.Println("AVATAR: ",spark.GetParameterByDetails(body,"avatar",x))
        fmt.Println("CREATED: ",spark.GetParameterByDetails(body,"created",x))
    }
}
func demoMyPeopleAction(token string) {
    fmt.Println("MY PEOPLE ==========================================================")
    status, headers, body, err := spark.GetMyDetail(token)
    peopleid = spark.GetParameterByDetails(body,"id",0)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    fmt.Println("DISPLAYNAME: ",spark.GetParameterByDetails(body,"displayName",0))
    fmt.Println("ID: ",peopleid)
    fmt.Println("EMAILS: ",spark.GetParameterByDetails(body,"emails",0))
    fmt.Println("AVATAR: ",spark.GetParameterByDetails(body,"avatar",0))

    fmt.Println("GET PEOPLE DETAIL ==================================================")
    status, headers, body, err = spark.GetPeopleDetail(token,peopleid)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    fmt.Println("DISPLAYNAME: ",spark.GetParameterByDetails(body,"displayName",0))
    fmt.Println("ID: ",peopleid)
    fmt.Println("EMAILS: ",spark.GetParameterByDetails(body,"emails",0))
    fmt.Println("AVATAR: ",spark.GetParameterByDetails(body,"avatar",0))
    fmt.Println("CREATED: ",spark.GetParameterByDetails(body,"created",0))
}
func demoListMembershipAction(token string) {
    //status, headers, body, err := spark.ListMemberships(token,"","","",0)  // my subscriptions
    status, headers, body, err := spark.ListMemberships(token,"<ROOMID>","","",0) // room subscriptions
    //status, headers, body, err := spark.ListMemberships(token,"","<PERSONID>","",0) // user subscription by id
    //status, headers, body, err := spark.ListMemberships(token,"","","<PERSONEMAIL>",0) // user subscription by email
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)

    fmt.Println("===============")
    for x := 0; x < spark.GetSize(body); x++ {
        fmt.Println("ID: ",spark.GetParameterByDetails(body,"id",x))
        fmt.Println("ROOMID: ",spark.GetParameterByDetails(body,"roomId",x))
        fmt.Println("PERSONID: ",spark.GetParameterByDetails(body,"personId",x))
        fmt.Println("PERSONEMAIL: ",spark.GetParameterByDetails(body,"personEmail",x))
        fmt.Println("PERSONDISPLAYNAME: ",spark.GetParameterByDetails(body,"personDisplayName",x))
        fmt.Println("ISMODERATOR: ",spark.GetParameterByDetails(body,"isModerator",x))
        fmt.Println("ISMONITOR: ",spark.GetParameterByDetails(body,"isMonitor",x))
        fmt.Println("CREATED: ",spark.GetParameterByDetails(body,"created",x))
    }
}
func demoCreateMembershipAction(token string) {
    roomid := "<ROOMID>"
    personid := "<PERSONID>"
    status, headers, body, err := spark.CreateMembership(token,roomid,personid,"",false)
    fmt.Println("STATUS: ",status)
    fmt.Println("HEADERS: ",headers)
    fmt.Println("ERROR: ",err)
    fmt.Println("BODY: ",body)
}
