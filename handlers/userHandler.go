package handlers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/ns/form"
	"github.com/ns/models"
	"github.com/ns/upload"
	"github.com/ns/user"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var Store =  sessions.NewCookieStore([]byte("secret"))

// UserHandler handler handles user related requests
type UserHandler struct {
	tmpl           *template.Template
	userService    user.UserService
	uploadService  upload.UploadService
	LoggedInUser   *models.User

}


func NewUserHandler(t *template.Template,
    usrServ user.UserService,
    upldServ upload.UploadService,
	) *UserHandler{
	return &UserHandler{tmpl: t,
	    userService: usrServ,
	    uploadService : upldServ,
	    }
}

func Authenticated(handler http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		session,_ := Store.Get(r,"session")
		username , ok := session.Values["username"]

		if !ok || len(username.(string)) == 0 {
			http.Redirect(w,r,"/login",302)
			return
		}
		handler.ServeHTTP(w,r)
	}
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		loginForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
		}{
			Values:  nil,
			VErrors: nil,
		}
		uh.tmpl.ExecuteTemplate(w, "index.html", loginForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		loginForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		usr, errs := uh.userService.GetUserByUsername(r.FormValue("username"))
		if len(errs) > 0 {
			loginForm.VErrors.Add("generic", "Your username or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "index.html", loginForm)
			return
		}
		if usr.Password != r.FormValue("pass"){
			loginForm.VErrors.Add("generic", "Your username or password is wrong")
			uh.tmpl.ExecuteTemplate(w, "index.html", loginForm)
			return
		}


		fmt.Println("about to redirect")
fmt.Println("________")
		session, _ := Store.Get(r,"session")
		session.Values["username"] = r.FormValue("username")
		session.Values["course"] = usr.Course
		session.Values["role"] = usr.Role
		session.Values["id"] = usr.ID
		session.Save(r,w)

			http.Redirect(w, r, "/homepage", http.StatusSeeOther)

	}
}


func (uh *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r,"session")
	session.Values["username"] = ""
	session.Values["course"] = ""
	session.Values["role"] = ""
	session.Options.MaxAge = -1
	session.Save(r,w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (uh *UserHandler) Nav(w http.ResponseWriter, r *http.Request) {
	usr, errrs := uh.userService.GetUserByUsername(LoggedIn(r))
	if len(errrs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	values := url.Values{}
	values.Add("name", usr.Name)
	form := struct {
		Values   url.Values
		VErrors  form.ValidationErrors
	}{
		Values:   values,
		VErrors:  form.ValidationErrors{},
	}
	if usr.Role == "Student"{
		uh.tmpl.ExecuteTemplate(w, "navbar.html", form)
		return
	}else{
		uh.tmpl.ExecuteTemplate(w, "teacherNav.html", form)
		return
	}

}

func (uh *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("username"))
	fmt.Println(r.FormValue("name"))
	fmt.Println(r.FormValue("password"))
	fmt.Println(r.FormValue("confirmpassword"))
	fmt.Println(r.FormValue("gender"))
	fmt.Println(r.FormValue("role"))
	fmt.Println(r.FormValue("course"))
	if r.Method == http.MethodGet {
		signUpForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
		}{
			Values:  nil,
			VErrors: nil,
		}
		uh.tmpl.ExecuteTemplate(w, "signup.html", signUpForm)
		return
	}

	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Validate the form contents
		//updateForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
		//updateForm.Required("name", "username","email", "password", "confrimpassword")
		//updateForm.MatchesPattern("email", form.EmailRX)
		//updateForm.MinLength("password", 8)
		//updateForm.PasswordMatches("password", "confrimpassword")
		//// If there are any errors, redisplay the signup form.
		//if !updateForm.Valid() {
		//	uh.tmpl.ExecuteTemplate(w, "signup.html", updateForm)
		//	return
		//}
		//
		//check := uh.userService.UsernameExists(r.FormValue("username"))
		//if check {
		//	updateForm.VErrors.Add("username", "Username Already Exists")
		//	uh.tmpl.ExecuteTemplate(w, "signup.html", updateForm)
		//	return
		//}
		//hashedPassword, err := password.Hash(r.FormValue("password"))
		//if err != nil {
		//	updateForm.VErrors.Add("password", "Password Could not be stored")
		//	uh.tmpl.ExecuteTemplate(w, "signup.html", updateForm)
		//	return
		//}
	//fmt.Println(hashedPassword)
		user := &models.User{
			Name: r.FormValue("name"),
			Username : r.FormValue("username"),
			Email:    r.FormValue("email"),
			Password: string(r.FormValue("password")),
			Gender:    r.FormValue("Gender"),
			Role: r.FormValue("role"),
			Course : r.FormValue("course"),
			Joindate : time.Now(),
		}
		fmt.Println(user)
		_, errs := uh.userService.RegisterUser(user)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (uh *UserHandler) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
    //exts := []string{".png", ".jpg", ".jpeg", ".gif"}

	usr, errrs := uh.userService.GetUserByUsername(LoggedIn(r))
	if len(errrs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

			if r.Method == http.MethodGet {

				usr,_:=uh.userService.GetUserByUsername(LoggedIn(r))

				values := url.Values{}
				values.Add("propic", usr.Image)
				values.Add("name", usr.Name)
				updateform := struct {
					Values   url.Values
					VErrors  form.ValidationErrors
				}{
					Values:   values,
					VErrors:  form.ValidationErrors{},
				}
				if usr.Role == "Teacher" {
					uh.tmpl.ExecuteTemplate(w, "teachereditprofile.html", updateform)
					return
				}else{
					uh.tmpl.ExecuteTemplate(w, "editprofile.html", updateform)
					return
				}
			}
			if r.Method == http.MethodPost {
				fmt.Println("its gets here")
				//errr := r.ParseMultipartForm(1 << 20)
				//if errr != nil {
				//	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				//	return
				//}
				r.ParseForm()
				fmt.Println("here to")
				updateForm := form.Input{Values: r.PostForm, VErrors: form.ValidationErrors{}}
				updateForm.MatchesPattern("email", form.EmailRX)
				updateForm.MinLength("new_password", 8)
				updateForm.MinLength("username", 8)
				updateForm.PasswordMatches("new_password", "confirm_new_password")
				//if !updateForm.Valid() {
				//	uh.tmpl.ExecuteTemplate(w, "editprofile.html", updateForm)
				//	return
				//}


				//if !form.ExtensionChecker(filepath.Ext(fh.Filename),exts){
				//  updateForm.VErrors.Add("upload", "You can't up'!")
				//  uh.tmpl.ExecuteTemplate(w, "editprofile.html", updateForm)
				// return
				//}



				fmt.Println(r.FormValue("name"))
				if r.FormValue("name") != ""{
					usr.Name = r.FormValue("name")
				}

				if r.FormValue("email") != ""{
					usr.Email = r.FormValue("email")
				}
				if r.FormValue("currentpassword") != "" && r.FormValue("newpassword") != "" &&  r.FormValue("confirmnewpassword") != ""{
					usr.Password = r.FormValue("newpassword")
				}
				if r.FormValue("bio") != ""{
					usr.About = r.FormValue("bio")
				}
				if r.FormValue("image") != " "{
					mf, fh, err := r.FormFile("image")
					if err != nil {
						updateForm.VErrors.Add("profilepic", "File error")
						uh.tmpl.ExecuteTemplate(w, "editprofile.html", updateForm)
						return
					}
						defer mf.Close()
						if fh.Filename != ""{
							usr.Image = Trimmer(fh.Filename)
						}
						UploadFile(&mf, Trimmer(fh.Filename))

				}




				_,errs := uh.userService.EditUser(usr)
				if len(errs) > 0 {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "/homepage", http.StatusSeeOther)
				return
			}
		}



func (uh *UserHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {

	usr, errrs := uh.userService.GetUserByUsername(LoggedIn(r))
		if len(errrs) > 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	values := url.Values{}
	values.Add("propic", usr.Image)
	values.Add("name", usr.Name)
	form := struct {
		Values   url.Values
		VErrors  form.ValidationErrors
		User  *models.User
	}{
		Values:   values,
		VErrors:  form.ValidationErrors{},
		User : usr,
	}

		uh.tmpl.ExecuteTemplate(w, "idnl.html", form)
		return


}

func (uh *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Println("getuserhandler")
	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		user, errs := uh.userService.GetUser(uint(id))
		values := url.Values{}
		values.Add("propic", user.Image)
		values.Add("name", user.Name)
		form := struct {
			Values   url.Values
			VErrors  form.ValidationErrors
			User  *models.User
		}{
			Values:   values,
			VErrors:  form.ValidationErrors{},
			User : user,
		}
		fmt.Println(user)
		if len(errs) > 0 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		uh.tmpl.ExecuteTemplate(w, "idnl.html", form)
	}
}