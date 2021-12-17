package api

import (
	"encoding/json"
	"fmt"
	"hello/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

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
	client_id := "6c85823c1a121313be52"
	client_secret := "82a8bd3beb1d432f0d5903aa5e47923be4d4a6c8"
	code := c.Query("code")
	state := c.Query("state")
	fmt.Println(state)
	resp, err := http.PostForm("https://github.com/login/oauth/access_token", url.Values{"client_id": {client_id}, "client_secret": {client_secret}, "code": {code}, "redirect_uri": {"http://47.119.167.128:3000/api/blog/github"}})
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println(err)
	comma := string(b)
	fmt.Println(comma[13:53])
	req, err1 := http.NewRequest("GET", "https://api.github.com/user", nil)
	fmt.Println(err1)
	fmt.Println(req)
	req.Header.Add("Authorization", "Bearer "+comma[13:53])
	client := &http.Client{}
	resp2, _ := client.Do(req)
	b2, err3 := ioutil.ReadAll(resp2.Body)
	userinfo := []byte(string(b2))
	fmt.Println(resp2)
	var i interface{}
	err4 := json.Unmarshal(userinfo, &i)
	if err4 != nil {
		fmt.Println(err4)
	}
	m := i.(map[string]interface{})
	login := m["login"].(string)
	fmt.Println(err3)
	c.SetCookie("key_cookie", login, 600, "/",
		"http://localhost:8080", false, true)
	c.Redirect(302, "http://localhost:8080")
}
