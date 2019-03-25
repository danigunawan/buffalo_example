package actions

import (
	"github.com/akosgarai/buffalo_example/models"
)

func (as *ActionSuite) Test_Auth_New() {
	res := as.HTML("/signin").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_Auth_Create() {
	u := &models.Administrator{
		Email:           "mark@example.com",
		Pwd:             "password",
		PwdConfirmation: "password",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	res := as.HTML("/signin").Post(u)
	as.Equal(302, res.Code)
	as.Equal("/", res.Location())
}

func (as *ActionSuite) Test_Auth_Create_Redirect() {
	u := &models.Administrator{
		Email:           "mark@example.com",
		Pwd:             "password",
		PwdConfirmation: "password",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	as.Session.Set("redirectURL", "/some/url")

	res := as.HTML("/signin").Post(u)
	as.Equal(302, res.Code)
	as.Equal(res.Location(), "/some/url")
}

func (as *ActionSuite) Test_Auth_Create_UnknownUser() {
	u := &models.Administrator{
		Email:    "mark@example.com",
		Password: "password",
	}
	res := as.HTML("/signin").Post(u)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}

func (as *ActionSuite) Test_Auth_Create_BadPassword() {
	u := &models.Administrator{
		Email:           "mark@example.com",
		Pwd:             "password",
		PwdConfirmation: "password",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	u.Password = "bad"
	res := as.HTML("/signin").Post(u)
	as.Equal(422, res.Code)
	as.Contains(res.Body.String(), "invalid email/password")
}
