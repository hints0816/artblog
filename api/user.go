package api

import (
	"encoding/json"
	"fmt"
	"hello/middleware"
	"hello/model"
	"hello/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": 1,
		},
	)
}

// GitHub OAuth App Callback
func GithubCallback(c *gin.Context) {
	client_id := utils.ClientId
	client_secret := utils.ClientSecret
	redirect_uri := utils.RedirectUri
	code := c.Query("code")
	state := c.Query("state")
	fmt.Println(state)
	accessResp, accessErr := http.
		PostForm("https://github.com/login/oauth/access_token",
			url.Values{
				"client_id":     {client_id},
				"client_secret": {client_secret},
				"code":          {code},
				"redirect_uri":  {redirect_uri},
			})
	if accessErr != nil {
		// log.Fatal(accessErr)
		return
	}
	defer accessResp.Body.Close()
	readBody, readBodyErr := ioutil.ReadAll(accessResp.Body)
	if readBodyErr != nil {
		// log.Fatal(readBodyErr)
		return
	}
	newUserResp, userErr := http.NewRequest("GET", "https://api.github.com/user", nil)
	if userErr != nil {
		// log.Fatal(userErr)
		return
	}
	githubToken := string(readBody)[13:53]
	newUserResp.Header.Add("Authorization", "Bearer "+githubToken)
	client := &http.Client{}
	userResp, _ := client.Do(newUserResp)
	readBody2, readBodyErr2 := ioutil.ReadAll(userResp.Body)
	if readBodyErr2 != nil {
		// log.Fatal(readBodyErr2)
		return
	}
	userinfo := []byte(string(readBody2))
	var i interface{}
	jsonErr := json.Unmarshal(userinfo, &i)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	m := i.(map[string]interface{})

	github := int(m["id"].(float64))
	githubUser, _ := model.GetUserByGithub(github)
	var profileData model.Profile
	// First Time Using Github
	if githubUser.Github == 0 {
		githubUser.Github = github
		githubUser.Username = strconv.Itoa(int(m["id"].(float64)))
		githubUser.Role = 1
		model.CreateUser(&githubUser)

		profileData.Name = m["login"].(string)
		profileData.ID = int(githubUser.ID)
		profileData.Avatar = m["avatar_url"].(string)
		model.AddProfile(&profileData)
	}

	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: githubUser.Username,
		Id:       githubUser.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "GinBlog",
		},
	}

	token, tokenErr := j.CreateToken(claims)
	if tokenErr != nil {
		fmt.Println(tokenErr)
	}
	c.SetCookie("remember_token", githubToken, 604800, "/",
		utils.Domain, false, true)
	c.SetCookie("token", token, 604800, "/",
		utils.Domain, false, true)
	c.Redirect(302, utils.Domain)
}
