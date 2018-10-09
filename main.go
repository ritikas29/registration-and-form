package main 
import (
    "fmt"
    "github.com/gorilla/context"
    "gopkg.in/mgo.v2"
    "log"
	"net/http"
	//"github.com/form/handler"
	routes "github.com/form/routes"
	//"github.com/form/model"
	//"mime/multipart"
)
func IsEmpty(data string) bool  {
    if len(data) == 0 {
        return true
    } else {
        return false
    }
}
func main() {

    db, err := mgo.Dial("localhost")
    if err != nil {
        log.Fatal("cannot dial mongo", err)
    }
    defer db.Close() // clean up when we're done

    // Adapt our handle function using withDB
    signup := Adapt(http.HandlerFunc(signuphandle), withDB(db))
    login := Adapt(http.HandlerFunc(loginhandle), withDB(db))
    upload := Adapt(http.HandlerFunc(uploadhandle), withDB(db))


    http.Handle("/signup", context.ClearHandler(signup))
    http.Handle("/login", context.ClearHandler(login))
    http.Handle("/upload", context.ClearHandler(upload))


    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("Server listening at 8080")
    }

}