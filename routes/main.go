package  routes 
import (
   "errors"
    "fmt"
    "net/http"
    handler "github.com/form/handler"
)
func loginhandle(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        http.Error(w, "Not supported", http.StatusMethodNotAllowed)
    default:
        handleLogin(w, r)
    }
}


func signuphandle(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        http.Error(w, "Not supported", http.StatusMethodNotAllowed)
    default:
        fmt.Println("SIGNUP DEFAULT CASE")
        handleSignup(w, r)
    }
}
func uploadhandle(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        http.Error(w, "Not supported", http.StatusMethodNotAllowed)
    default:
        handleUpload(w, r)
    }
}