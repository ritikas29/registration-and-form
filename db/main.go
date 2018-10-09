package db
import (
    "github.com/gorilla/context"
    "gopkg.in/mgo.v2"
	"net/http"
	//"github.com/form/models"
	"github.com/form/routes"
    
)
//type Adapter func(http.Handler) http.Handler

func withDB(db *mgo.Session) Adapter {

    // return the Adapter
    return func(h http.Handler) http.Handler {

        // the adapter (when called) should return a new handler
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

            // copy the database session
            dbsession := db.Copy()
            defer dbsession.Close() // clean up

            // save it in the mux context
            context.Set(r, "database", dbsession)

            // pass execution to the original handler
            h.ServeHTTP(w, r)

        })
    }
}