package page

import (
	"go_ssr_template/models"
	"go_ssr_template/session"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LoginPage struct {
	Page
	db             *gorm.DB
	sessionManager *session.Manager
}

type LoginPageData struct {
	PageData
}

func NewLoginPage(db *gorm.DB, sessionManager *session.Manager) *Page {
	deps := &LoginPage{
		db:             db,
		sessionManager: sessionManager,
	}

	return &Page{
		Path:           "/login",
		Template:       "login",
		Deps:           deps,
		GetPageData:    deps.GetPageData,
		GetPostHandler: deps.GetPostHandler(),
	}
}

func (p *LoginPage) GetPageData(c echo.Context) any {
	pageData := LoginPageData{
		PageData: PageData{
			Title:    "Login page",
			UrlError: c.QueryParam("error"),
			Success:  c.QueryParam("success"),
		},
	}

	return pageData
}

func (p *LoginPage) GetPostHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		submitUser := models.User{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}

		if submitUser.Password == "" || submitUser.Email == "" {
			return c.Redirect(http.StatusSeeOther, p.Path+"?error=need username and password")
		}

		dbUser := &models.User{}
		result := p.db.Where(&models.User{Email: submitUser.Email}).Find(dbUser)
		if result.Error != nil {
			logrus.
				WithError(result.Error).
				Error("error getting user from the database")
			return c.Redirect(http.StatusSeeOther, p.Path+"?error=internal server problem")
		}

		if dbUser.Password == "" {
			return c.Redirect(http.StatusSeeOther, p.Path+"?error=wrong user name or password")
		}

		comparison, err := dbUser.ComparePassword(submitUser.Password)
		if err != nil {
			logrus.
				WithError(err).
				Error("error comparing passwords")
			return c.Redirect(http.StatusSeeOther, p.Path+"?error=internal server problem")
		}

		if !comparison {
			return c.Redirect(http.StatusSeeOther, p.Path+"?error=wrong user name or password")
		}

		p.sessionManager.InitSession(dbUser.Email, dbUser.ID, c)
		return c.Redirect(http.StatusSeeOther, p.Path+"?success=logged in")
	}
}
