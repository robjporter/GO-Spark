package spark

import (
    "fmt"
    "bytes"
    "io/ioutil"
    "net/url"
    "net/http"
    "encoding/json"
    "github.com/roporter/go-libs/as"
    "github.com/roporter/go-libs/jmespath"
)

const (
    GET             = 0
    POST            = 1
    PUT             = 2
    DELETE          = 3
)

const (
    responseSeparator = ","
)

type transaction struct {
    url             string
    body            string
    method          int
    token           string
    responseBody    string
    responseStatus  string
    responseHeaders []string
}

var successResponses = []int{200,204}
var failureResponses = []int{400,401,403,404,409,500,503}

func getHTTPMethodName(method int) string {
    switch(method) {
    case 0:
        return "GET"
    case 1:
        return "POST"
    case 2:
        return "PUT"
    case 3:
        return "DELETE"
    default:
        return "GET"
    }
}

func GetISO8601Timestamp(year,month,day,hour,minute string) string {
    return year + "-" + month + "-" + day + "T" + hour + ":" + minute + ":" + "00.000+01:00"
}

func getErrorCodeDescription(code int) string {
    switch(code) {
    case 200:
        return "OK"
    case 204:
        return "Deleted"
    case 400:
        return "The request was invalid or cannot be otherwise served. An accompanying error message will explain further."
    case 401:
        return "Authentication credentials were missing or incorrect."
    case 403:
        return "The request is understood, but it has been refused or access is not allowed."
    case 404:
        return "The URI requested is invalid or the resource requested, such as a user, does not exist. Also returned when the requested format is not supported by the requested method."
    case 409:
        return "The request could not be processed because it conflicts with some established rule of the system. For example, a person may not be added to a room more than once."
    case 500:
        return "Something went wrong on the server."
    case 503:
        return "Server is overloaded with requests. Try again later."
    default:
        return "Something went wrong and we recevied an unhandled error code."
    }
}

func headerToArray(header http.Header) (res []string) {
    for name, values := range header {
        for _, value := range values {
            res = append(res, fmt.Sprintf("%s: %s", name, value))
        }
    }
    return
}

func urlEncode(str string) string {
    return url.QueryEscape(str)
}

func GetSize(body string) int {
    var data interface{}
    expression := "items.length(@)"
    json.Unmarshal([]byte(body), &data)
    result, err := jmespath.Search(expression, data)
    if err == nil {
        return as.Int(result)
    }
    return 0
}

func GetParameterByDetails(body string, parameter string, pos int) string {
    var data interface{}
    expression := "items.length(@)"
    json.Unmarshal([]byte(body), &data)
    result, err := jmespath.Search(expression, data)
    if err != nil {
        expression = parameter
        result2, _ := jmespath.Search(expression, data)
        return as.String(result2)
    } else {
        if as.Int(result) > pos {
            expression = "items[" + as.String(pos) + "]." + parameter
            result2, _ := jmespath.Search(expression, data)
            return as.String(result2)
        }
    }
    return ""
}

func (t transaction) doTransaction() (transaction,error) {
    var jsonStr =[]byte{}
    if t.body != "" {
        jsonStr = []byte(t.body)
    }
    req, err := http.NewRequest(getHTTPMethodName(t.method), t.url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    req.Header.Set("Authorization", "Bearer " + t.token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body,_ := ioutil.ReadAll(resp.Body)
    t.responseStatus = string(resp.Status)
    t.responseHeaders = headerToArray(resp.Header)
    t.responseBody = string(body)
    return t,nil
}
