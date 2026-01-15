package API

import (
	"encoding/base64"
	"encoding/json"
	"github.com/superg3m/stoic-go/Core/Utility"
	"net/http"
	"time"

	"github.com/superg3m/stoic-go/Core/Router"
	"github.com/superg3m/stoic-go/inc/LoginKey"
	"github.com/superg3m/stoic-go/inc/User"
)

func checkUserAuth(request *Router.StoicRequest, response *Router.StoicResponse) {
	cookie, err := request.Cookie("go_garden_auth_token")
	if err != nil {
		response.AddError("Unauthorized!")
		http.SetCookie(response, &http.Cookie{
			Name:     "go_garden_auth_token",
			Value:    "SECRET_VALUE",
			Expires:  time.Now().Add(0 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
		})
		return
	}

	decoded, err3 := base64.StdEncoding.DecodeString(cookie.Value)
	if err3 != nil {
		response.AddError(err3.Error())
		return
	}

	var cookieData User.CookieData
	if err2 := json.Unmarshal(decoded, &cookieData); err2 != nil {
		response.AddError(err2.Error())
		return
	}

	response.SetData(cookieData.ID)
}

func loginUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")

	user, errors := User.FromEmail(email)
	if user == nil {
		response.AddErrors(errors, "Email or password don't match!")
		return
	}

	loginKey, errors1 := LoginKey.FromUserID_Provider(user.ID, LoginKey.PASSWORD)
	if loginKey == nil {
		response.AddErrors(errors1, "Email or password don't match!")
		return
	}

	hashedPassword := Utility.Sha256HashString(password)
	if loginKey.Key != hashedPassword {
		response.AddError("Email or password don't match!")
		return
	}

	_, err := request.Cookie("go_garden_auth_token")
	if err == nil {
		response.SetData("Auth Token Found!")
		return
	} else {
		jsonData, _ := json.Marshal(User.CookieData{ID: user.ID})
		encoded := base64.StdEncoding.EncodeToString(jsonData)
		http.SetCookie(response, &http.Cookie{
			Name:     "go_garden_auth_token",
			Value:    encoded,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
			Path:     "/",
		})

		response.SetData("Login Successful!")
	}
}

func logoutUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	http.SetCookie(response, &http.Cookie{
		Name:     "go_garden_auth_token",
		Value:    "SECRET_VALUE",
		Expires:  time.Now().Add(0 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
	})
}

func createUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	email := request.GetStringParam("email")
	password := request.GetStringParam("password")

	user := User.New()
	user.Email = email
	create := user.Create()
	if create.IsBad() {
		response.AddErrors(create.GetErrors(), "Failed to create user")

		return
	}

	loginKey := LoginKey.New()
	loginKey.UserID = user.ID
	loginKey.Key = password
	loginKey.Provider = LoginKey.PASSWORD
	loginKey.Key = Utility.Sha256HashString(loginKey.Key)
	create = loginKey.Create()
	if create.IsBad() {
		response.AddErrors(create.GetErrors(), "Failed to create login key")
		user.Delete()

		return
	}

	// response.SetData(user) // This should work but right now it doesn't
	response.SetData("User Created Successfully!")
}

func getUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	id := request.GetIntParam("id")
	user, errors := User.FromID(id)
	if errors != nil {
		response.AddErrors(errors, "Failed to get user from ID")
		return
	}

	response.SetData(user)
}

func updateUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	id := request.GetIntParam("id")
	email := request.GetStringParam("email")
	oldPassword := request.GetStringParam("oldPassword")
	newPassword := request.GetStringParam("newPassword")

	user, errors := User.FromID(id)
	if errors != nil {
		response.AddErrors(errors, "Failed to get user from Email")
		return
	}

	// user.LastLogin = Utility.NewTime(time.Now())

	if email != "" {
		user.Email = email
	}
	update := user.Update()
	if update.IsBad() {
		response.AddErrors(update.GetErrors(), "Failed to update user")
		return
	}

	loginKey, errors2 := LoginKey.FromUserID_Provider(user.ID, LoginKey.PASSWORD)
	if errors2 != nil {
		response.AddErrors(errors2, "Failed to get LoginKey from UserID and Provider")
		return
	}

	if loginKey.Key != Utility.Sha256HashString(oldPassword) {
		response.AddError("Old password don't match!")
		return
	}

	loginKey.Key = Utility.Sha256HashString(newPassword)
	update = loginKey.Update()
	if update.IsBad() {
		response.AddErrors(update.GetErrors(), "Failed to update login key")
		return
	}

	response.SetData("User Updated Successfully!")
}

func deleteUser(request *Router.StoicRequest, response *Router.StoicResponse) {
	id := request.GetIntParam("id")

	user, errors := User.FromID(id)
	if errors != nil {
		response.AddErrors(errors, "Failed to get user from ID")
		return
	}

	del := user.Delete()
	if del.IsBad() {
		response.AddErrors(del.GetErrors(), "Failed to delete user")
		return
	}

	response.SetData("User Deleted Successfully!")
}

func init() {
	Router.RegisterApiEndpoint("User/Login", loginUser, "POST",
		Router.MiddlewareValidParams("email", "password"),
	)

	Router.RegisterApiEndpoint("User/Logout", logoutUser, "POST")
	Router.RegisterApiEndpoint("User/Authorized", checkUserAuth, "POST")

	Router.RegisterApiEndpoint("User", createUser, "POST",
		Router.MiddlewareValidParams("email", "password"),
	)
	Router.RegisterApiEndpoint("User", getUser, "GET",
		Router.MiddlewareValidParams("id"),
	)
	Router.RegisterApiEndpoint("User", updateUser, "PATCH",
		Router.MiddlewareValidParams("id", "email", "oldPassword", "newPassword"),
	)
	Router.RegisterApiEndpoint("User", deleteUser, "DELETE",
		Router.MiddlewareValidParams("id"),
	)
}
