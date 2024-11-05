package main

import (
	"go-api/utils"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	magic := r.Group("/magic")
	{
		magic.GET("/sum", magicSumHandler)
		magic.POST("/pow", magicPowHandler)
		magic.GET("/odd", magicOddHandler)
		magic.POST("/grade", magicGradeHandler)
		magic.GET("/name", magicNameHandler)
		magic.POST("/tria", magicTriaHandler)
	}

	account := r.Group("/account")
	{
		account.GET("/create", createAccountHandler)
		account.GET("/read", readAccountHandler)
		account.GET("/update", updateAccountHandler)
		account.GET("/delete", deleteAccountHandler)
		account.GET("/list", listAccountHandler)
		account.GET("/yourusername", usernameHandler)
		account.POST("/login", authLoginHandler)
	}

	r.Run()

}

func magicSumHandler(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number"})
		return
	}
	result := utils.MagicSum(n)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func magicPowHandler(c *gin.Context) {
	var data struct {
		Num int `json:"n"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	result := utils.MagicPow(data.Num)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func magicOddHandler(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number"})
		return
	}
	result := utils.MagicOdd(n)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func magicGradeHandler(c *gin.Context) {
	var data struct {
		Num int `json:"n"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	result := utils.MagicGrade(data.Num)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func magicNameHandler(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number"})
		return
	}
	result := utils.MagicName(n)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func magicTriaHandler(c *gin.Context) {
	var data struct {
		Num int `json:"n"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	result := utils.MagicTria(data.Num)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func createAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func readAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{},
	})
}

func updateAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func deleteAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func listAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": []interface{}{},
	})
}

func usernameHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Halo, Nama saya Rizky",
	})
}

// AuthLogin Handler
func authLoginHandler(c *gin.Context) {
	var data struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	isValidUsername := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(data.Username)
	// Ubah regex untuk password untuk menerima huruf dan angka
	isValidPassword := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(data.Password)

	if isValidUsername && isValidPassword {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
	}
}
