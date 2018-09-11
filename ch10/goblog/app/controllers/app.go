package controllers

import (
	"goblog/app/models"

	"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
)

type App struct {
	GormController
	CurrentUser *models.User
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) CreateSession(username, password string) revel.Result {
	var user models.User

	// username으로 사용자 조회
	c.Txn.Where(&models.User{Username: username}).First(&user)

	// bcrypt 패키지의 CompareHashAndPassword 함수로 패스워드 비교
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	// 패스워드가 일치하면 세션 생성후 포스트 목록 화면으로 이동
	if err == nil {
		authKey := revel.Sign(user.Username)
		c.Session["authKey"] = authKey
		c.Session["username"] = user.Username
		c.Flash.Success("Welcome, " + user.Name)
		return c.Redirect(Post.Index)
	}

	// 세션 정보를 모두 제거하고 홈으로 이동
	for k := range c.Session {
		delete(c.Session, k)
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(Home.Index)
}

func (c App) DestroySession() revel.Result {
	// clear session
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(Home.Index)
}

func (c *App) setCurrentUser() revel.Result {
	// 뷰에서 currentUser를 사용할 수 있게 RenderArgs에 CurrentUser를 추가
	defer func() {
		if c.CurrentUser != nil {
			c.ViewArgs["currentUser"] = c.CurrentUser
		} else {
			delete(c.ViewArgs, "currentUser")
		}
	}()

	// 세션에서 username과 authKey를 가져옴
	username, ok := c.Session["username"]
	if !ok || username == "" {
		return nil
	}

	authKey, ok := c.Session["authKey"]
	if !ok || authKey == "" {
		return nil
	}

	// revel.Verify 함수로 authKey가 유효한지 확인
	// authKey가 유효하면 username으로 사용자를 조회하고 컨트롤러의 CurrentUser에 저장
	if match := revel.Verify(username, authKey); match {
		var user models.User
		c.Txn.Where(&models.User{Username: username}).First(&user)
		if &user != nil {
			c.CurrentUser = &user
		}
	}
	return nil
}
