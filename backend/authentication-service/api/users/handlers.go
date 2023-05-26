package users

import (
	"authentication-service/database"
	"authentication-service/services/authentication"
	"authentication-service/services/user"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func signUpNewUser(
	dbManager database.GormDbManager,
	userService user.Service,
	authenticationService authentication.Service,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		}

		err := c.ShouldBindWith(&form, binding.JSON)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request form")
			return
		}

		createdUser, createError := userService.Create(c, dbManager, form.Email, form.FirstName, form.LastName, form.Password)
		if createError != nil {
			c.JSON(http.StatusBadRequest, createError.JSON())
			return
		}

		token, tokenError := authenticationService.GenerateToken(c, dbManager, form.Email)
		if tokenError != nil {
			c.JSON(http.StatusBadRequest, tokenError.JSON())
			return
		}

		createdUser.Token = token

		c.JSON(http.StatusOK, createdUser)
	}
}

func login(
	dbManager database.GormDbManager,
	userService user.Service,
	authenticationService authentication.Service,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := c.ShouldBindWith(&form, binding.JSON)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request form")
			return
		}

		user, userError := userService.LoginUser(c, dbManager, form.Email, form.Password)
		if userError != nil {
			c.JSON(http.StatusBadRequest, userError.JSON())
			return
		}

		token, userError := authenticationService.GenerateToken(c, dbManager, form.Email)
		if userError != nil {
			c.JSON(http.StatusBadRequest, userError.JSON())
			return
		}

		user.Token = token

		c.JSON(http.StatusOK, user)
	}
}

func logout(
	dbManager database.GormDbManager,
	authenticationService authentication.Service,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form struct {
			Email string `json:"email"`
		}

		err := c.ShouldBindWith(&form, binding.JSON)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request form")
			return
		}

		token, generateErr := authenticationService.GenerateTokenForUserLogout(c, dbManager, form.Email)
		if generateErr != nil {
			c.JSON(http.StatusBadRequest, generateErr.JSON())
			return
		}

		c.JSON(http.StatusOK, token)
	}
}

func getUserMe(
	dbManager database.GormDbManager,
	userService user.Service,
	authenticationService authentication.Service,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		userEmail, err := authenticationService.FetchUserEmailFromToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.JSON())
			return
		}

		user, err := userService.FindByEmail(c, dbManager, userEmail)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.JSON())
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
