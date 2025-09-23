package fncGlobal

import (
	mdlGlobal "back/global/model"
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Handle Login function
func FncGlobalAllusrLoginx(c *gin.Context) {

	// Variable login
	var usript mdlGlobal.MdlGlobalAllusrParams
	var usrdbs mdlGlobal.MdlGlobalAllusrDtbase

	// Bind JSON body input to var usript
	if err := c.BindJSON(&usript); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("allusr_usrlst")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find user by username in database
	err := tablex.FindOne(contxt, bson.M{"usrnme": usript.Usrnme}).Decode(&usrdbs)
	if err != nil {
		c.JSON(401, gin.H{"error": "usrnme"})
		return
	}

	// Compare provided password with stored password hash
	err = bcrypt.CompareHashAndPassword([]byte(usrdbs.Psswrd), []byte(usript.Psswrd))
	if err != nil {
		c.JSON(401, gin.H{"error": "psswrd"})
		return
	}

	// Generate JWT Token
	claimp := &mdlGlobal.MdlGlobalAllusrInputs{
		Usrnme: usript.Usrnme,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		},
	}

	// Translate JWT Token to strings
	tknraw := jwt.NewWithClaims(jwt.SigningMethodHS256, claimp)
	tknstr, err := tknraw.SignedString(jwtkey)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	// Konversi array mdlAllusrLoginxParams
	tknobj := &mdlGlobal.MdlGlobalAllusrTokens{
		Stfnme: usrdbs.Stfnme,
		Usrnme: usrdbs.Usrnme,
		Access: usrdbs.Access,
		Keywrd: usrdbs.Keywrd,
	}
	tknfnl, err := json.Marshal(tknobj)
	if err != nil {
		c.JSON(500, gin.H{"error": "Unable to set cookie"})
		return
	}

	// Set the JWT token in the cookie
	c.SetCookie("tokenx", tknstr, 10800, "/", Ipadrs, false, true)
	c.SetCookie("nowusr", string(tknfnl), 10800, "/", Ipadrs, false, true)
	c.JSON(200, "Login Successfull")
}

// Handle Logout function
func FncGlobalAllusrLogout(c *gin.Context) {

	// Delete Cookie
	c.SetCookie("tokenx", "", -1, "/", Ipadrs, false, true)
	c.SetCookie("nowusr", "", -1, "/", Ipadrs, false, true)
	c.JSON(200, "Logout")
}

// Handle Logout function
func FncGlobalAllusrTokenx(c *gin.Context) {

	// Get cookie
	cookie := c.GetHeader("Authorization")
	if cookie == "" {
		c.String(401, "Authorization header missing")
		return
	}

	// Convert JWT to claims
	tokenx, err := jwt.ParseWithClaims(cookie, &mdlGlobal.MdlGlobalAllusrInputs{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtkey, nil
		})

	// Final Result
	if err != nil {
		c.JSON(500, "Loggin First")
	} else if _, ok := tokenx.Claims.(*mdlGlobal.MdlGlobalAllusrInputs); ok {
		c.JSON(200, "Access Accepted")
	} else {
		c.JSON(500, "Loggin First")
	}
}

func FncGlobalAllusrApplst(c *gin.Context) {

	// Select database and collection
	tablex := Client.Database(Dbases).Collection("allusr_applst")
	contxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get route data
	datarw, err := tablex.Find(contxt, bson.M{})
	if err != nil {
		panic("fail")
	}
	defer datarw.Close(contxt)

	// Append to slice
	var slices = []mdlGlobal.MdlGlobalAllusrApplst{}
	for datarw.Next(contxt) {
		var object mdlGlobal.MdlGlobalAllusrApplst
		if err := datarw.Decode(&object); err == nil {
			slices = append(slices, object)
		}
	}

	// Send token to frontend
	c.JSON(200, slices)
}
