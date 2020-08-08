package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/ns/database"
	"github.com/ns/handlers"
	upldRepo "github.com/ns/upload/repository"
	upldServ "github.com/ns/upload/service"

	userrepo "github.com/ns/user/repository"
	userserv "github.com/ns/user/services"
)
var tmpl *template.Template


func main() {


	tmpl := template.Must(template.ParseGlob("*.html"))
	db:=database.Conn()
	defer db.Close()



	userRepo := userrepo.NewUserRepositoryImpl(db)
	userServ := userserv.NewUserServiceImpl(userRepo)




	upldrepo := upldRepo.NewUploadRepositoryImpl(db)
	upldserv := upldServ.NewUploadSerivceImpl(upldrepo)

	uh := handlers.NewUserHandler(tmpl, userServ,upldserv)

	fs := http.FileServer(http.Dir("asset"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))
	http.HandleFunc("/", uh.Login)
	http.HandleFunc("/logout", uh.Logout)
	http.HandleFunc("/signup", uh.Signup)
	http.HandleFunc("/homepage",handlers.Authenticated(uh.UploadedNotesHandler))
	http.HandleFunc("/upload",handlers.Authenticated(uh.UploadNoteHandler))
	http.HandleFunc("/editprofile",handlers.Authenticated(uh.UpdateProfileHandler))
	http.HandleFunc("/profilepage",handlers.Authenticated(uh.ProfileHandler))
	http.HandleFunc("/mynotes",handlers.Authenticated(uh.MyNotesHandler))
	http.HandleFunc("/notes/delete",handlers.Authenticated(uh.DeleteNotesHandler))
	http.HandleFunc("/notes/approve",handlers.Authenticated(uh.ApproveNotesHandler))
	http.HandleFunc("/profile",handlers.Authenticated(uh.GetUserHandler))
	http.HandleFunc("/a",index)
	http.ListenAndServe(":8080", nil)

}
func index(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content_Type","text/html; charset=utf-8")
	io.WriteString(w,`<img src="asset/files/payoneer.jpg">`)
}

