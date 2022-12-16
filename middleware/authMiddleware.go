package middleware

import (
	"errors"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	email, password, ok := c.Request.BasicAuth()

	if !ok {
		c.AbortWithError(http.StatusBadRequest, errors.New("authentication not found"))
		return
	}

	sqlStatement := `
	SELECT * FROM users
	WHERE email = $1 AND password = $2;`
	res, err := database.DbConnection.Exec(sqlStatement, email, password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if count == 0 {
		c.AbortWithError(http.StatusUnauthorized, errors.New("email atau password salah"))
		return
	}

	c.Next()
}
