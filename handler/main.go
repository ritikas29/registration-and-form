package handler 
import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/context"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    //db "github.com/form/db"
    model "github.com/form/models"
    "os"
    "io/ioutil"
    "io"
    "mime/multipart"
    "net/http"
    "bytes"
    "errors"

)
type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
    for _, adapter := range adapters {
        h = adapter(h)
    }
    return h
}
func handleLogin(w http.ResponseWriter, r *http.Request) {
    db := context.Get(r, "database").(*mgo.Session)

    // decode the request body
    var c loginPayload
    if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    email:= c.Email
    //password := c.Password
    //find in the db user with this email
    result := signupPayload{}
    err := db.DB("QCI").C("usersignup").Find(bson.M{"email": email}).Select(bson.M{"username": "", email: ""}).One(&result)
    if err != nil {
        panic(err)
    }
    fmt.Println("Results", result.Username)
    if result.Email != email {
        fmt.Println("Welcome", result.Username)
        fmt.Fprintln(w, "Login successful.")
    }

    // redirect to it
    //http.Redirect(w, r, "/welcome/"+c.ID.Hex(), http.StatusTemporaryRedirect)
}
func handleSignup(w http.ResponseWriter, r *http.Request) {
    db := context.Get(r, "database").(*mgo.Session)
    //decoder := json.NewDecoder(r.Body)
    var c signupPayload
    // to print the incoming variables

    //_ = decoder.Decode(&c)
    //
    //
     //fmt.Println("%v",c.Email)
       //fmt.Println("%v",c.Password)
     //  fmt.Println("%v",c.Username)

    // decode the request body

    if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := db.DB("QCI").C("usersignup").Insert(&c); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        fmt.Fprintln(w, "Signup unsuccessful.")
        return
    } else {
        fmt.Fprintln(w, "Signup successful.")
    }

    //signupConfirmLink := "HOST" + "/welcome/"+c.ID.Hex()
    //sendEmail("user email from request", signupConfirmLink)
}
func handleUpload(w http.ResponseWriter, r *http.Request) error {
    db := context.Get(r, "database").(*mgo.Session)
    var c uploadPayload

    //s3 url will come here
    targetUrl := "https://localhost" //"https://12.343.134"
    if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return nil
    }

    filename:= c.Filename
    bodyBuf := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuf)

    // this step is very important
    fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
    if err != nil {
        fmt.Println("error writing to buffer")
        return err
    }

    // open file handle
    fh, err := os.Open(filename)
    if err != nil {
        fmt.Println("error opening file")
        return err
    }
    defer fh.Close()

    //iocopy
    _, err = io.Copy(fileWriter, fh)
    if err != nil {
        return err
    }

    contentType := bodyWriter.FormDataContentType()
    bodyWriter.Close()

    resp, err := http.Post(targetUrl, contentType, bodyBuf)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    resp_body, err := ioutil.ReadAll(resp.Body)
    payload := &uploadDBPayload{image:resp_body}
    if err := db.DB("QCI").C("uploadimage").Insert(&payload); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        fmt.Fprintln(w, "Image Upload unsuccessful.")
        return errors.New("Image Upload unsuccessful.")
    } else {
        fmt.Fprintln(w, "Image Upload successful.")
    }
    return nil
}

