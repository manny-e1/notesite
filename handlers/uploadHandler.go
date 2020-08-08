package handlers

import (
	"fmt"
	"path/filepath"

	//"path/filepath"
	"time"

	//"fmt"
	"net/http"
	"net/url"
	//"time"
	"strconv"
	//"path/filepath"
	"github.com/ns/form"
	"github.com/ns/models"
)

//
func LoggedIn(r *http.Request) string{
	session, _ := Store.Get(r,"session")
	untyped, ok := session.Values["username"]
	if !ok{
		return ""
	}
	usernamee, ok := untyped.(string)
	if !ok {
		return ""
	}

	return usernamee
}
func CourseName(r *http.Request) string{
	session, _ := Store.Get(r,"session")
	untyped, ok := session.Values["course"]
	if !ok{
		return ""
	}
	course, ok := untyped.(string)
	if !ok {
		return ""
	}

	return course
}

func UserID(r *http.Request) uint{
	session, _ := Store.Get(r,"session")
	untyped, ok := session.Values["id"]
	if !ok{
		return 1
	}
	id, ok := untyped.(uint)
	if !ok {
		return 1
	}

	return id
}



func (uh *UserHandler) UploadedNotesHandler(w http.ResponseWriter, r *http.Request) {


	usr, _ := uh.userService.GetUserByUsername(LoggedIn(r))
	if usr.Role == "Teacher"{
		notes, _ := uh.uploadService.NotesByCourseName(CourseName(r), "")
		note, _ := uh.uploadService.NotesByCourseName(CourseName(r), "approved")
		values := url.Values{}
		values.Add("propic", usr.Image)
		values.Add("name", usr.Name)
		values.Add("username", usr.Username)

		tmplData := struct {
			Values     url.Values
			VErrors    form.ValidationErrors
			Notes []models.Upload
			Note []models.Upload

		}{
			Values:     nil,
			VErrors:    nil,
			Notes: notes,
			Note: note,
		}
		//fmt.Println(tmplData)

		uh .tmpl.ExecuteTemplate(w, "teacherHomepage.html", tmplData)
	}else {
		notes, _ := uh.uploadService.NotesByCourseName(CourseName(r), "approved")
		tmplData := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			Note   []models.Upload
		}{
			Values:  nil,
			VErrors: nil,
			Note:   notes,
		}
		//fmt.Println(tmplData)

		uh.tmpl.ExecuteTemplate(w, "homepage.html", tmplData)
	}
}

func (uh *UserHandler) MyNotesHandler(w http.ResponseWriter, r *http.Request) {

	notes, _ := uh.uploadService.NotesByUploader(LoggedIn(r),"approved")
	usr,_ := uh.userService.GetUserByUsername(LoggedIn(r))
	values := url.Values{}
	values.Add("propic", usr.Image)
	values.Add("name", usr.Name)

	tmplData := struct {
			Values     url.Values
			VErrors    form.ValidationErrors
			Notes []models.Upload

		}{
			Values:     nil,
			VErrors:    nil,
			Notes: notes,
		}
	if usr.Role == "Student"{
		uh .tmpl.ExecuteTemplate(w, "mynotes.html", tmplData)
	}else{
		uh .tmpl.ExecuteTemplate(w, "teachernote.html", tmplData)
	}
}

func (uh *UserHandler) UploadNoteHandler(w http.ResponseWriter, r *http.Request) {
exts := []string{".pdf", ".docx", ".sys", ".txt", ".doc", ".ppt" , ".pptx",".zip"}
fmt.Println("in this")
	usr,_ := uh.userService.GetUserByUsername(LoggedIn(r))
	if r.Method == http.MethodGet {
		values := url.Values{}
		values.Add("name", usr.Name)

		uploadForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
		}{
			Values:  nil,
			VErrors: nil,
		}
		fmt.Println("get here")
		if usr.Role == "Student"{
			uh .tmpl.ExecuteTemplate(w, "uploadnote.html", uploadForm)
		}else{
			uh .tmpl.ExecuteTemplate(w, "teacherupload.html", uploadForm)
		}
	}
fmt.Println("got here")
fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		// Parse the form data
		usr,_ := uh.userService.GetUserByUsername(LoggedIn(r))
		err := r.ParseMultipartForm(1 << 20)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Validate the form contents
		uploadForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		uploadForm.Required("title", "description")
		uploadForm.MinLength("title",5)
		uploadForm.MinLength("description", 10)
		// If there are any errors, redisplay the signup form.
		if !uploadForm.Valid() {
			uh.tmpl.ExecuteTemplate(w, "uploadnote.html", uploadForm)
			return
		}
		fmt.Println("here again")
		mf, fh, err := r.FormFile("file")
		if err != nil {
			uploadForm.VErrors.Add("upload", "File error")
			uh .tmpl.ExecuteTemplate(w, "uploadnote.html", uploadForm)
			fmt.Println(err)
		}

		defer mf.Close()
		fmt.Println("ezi zershalew")
		fmt.Println(form.ExtensionChecker(fh.Filename,exts))
		if form.ExtensionChecker(filepath.Ext(fh.Filename),exts) == false{
			uploadForm.VErrors.Add("upload", "You can only upload the listed File Types!")
			uh .tmpl.ExecuteTemplate(w, "uploadnote.html", uploadForm)
			return
		}
		fmt.Println("upload maregu ga")
		UploadFile(&mf, Trimmer(fh.Filename))
		if usr.Role == "Student"{
			upld := &models.Upload{
				Title: r.FormValue("title"),
				File_Description: r.FormValue("description"),
				File_Path: Trimmer(fh.Filename),
				File_Uploader: LoggedIn(r) ,
				File_Uploader_ID:uint(usr.ID),
				File_Uploaded_To : CourseName(r),
				File_Type: filepath.Ext(fh.Filename),
				File_Uploaded_On: time.Now(),

			}
			_, errs := uh .uploadService.StoreNote(upld)
			if len(errs) > 0 {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}else{
			upld := &models.Upload{
				Title: r.FormValue("title"),
				File_Description: r.FormValue("description"),
				File_Path: Trimmer(fh.Filename),
				File_Uploader: LoggedIn(r) ,
				File_Uploader_ID:uint(usr.ID),
				File_Uploaded_To : CourseName(r),
				File_Type: filepath.Ext(fh.Filename),
				File_Uploaded_On: time.Now(),
				File_Status:"approved",

			}
			_, errs := uh .uploadService.StoreNote(upld)
			if len(errs) > 0 {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}



		http.Redirect(w, r, "/homepage", http.StatusSeeOther)
	}
}

func (uh *UserHandler) DeleteNotesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		_, errs := uh.uploadService.DeleteNote(uint(id))
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
	http.Redirect(w, r, "/mynotes", http.StatusSeeOther)
}
func (uh *UserHandler) ApproveNotesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("approve")
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		 uh.uploadService.ApproveNote(uint(id))

	}
	http.Redirect(w, r, "/homepage", http.StatusSeeOther)
}