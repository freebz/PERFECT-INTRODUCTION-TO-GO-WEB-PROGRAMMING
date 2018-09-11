package controllers

import (
	"goblog/app/models"
	"goblog/app/routes"
	
	"github.com/revel/revel"
)

// App을 임베디드 필드로 지정
type Post struct {
	App
}

func (c Post) CheckUser() revel.Result {
	// Index와 Show는 권한을 확인하지 않음.
	switch c.MethodName {
	case "Index", "Show":
		return nil
	}

	// CurrentUser 정보가 없으면 로그인 페이지로 이동
	if c.CurrentUser == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Login)
	}

	// CurrentUser가 관리자가 아니면 로그인 페이지로 이동
	if c.CurrentUser.Role != "admin" {
		c.Response.Status = 401 // Unauthorized
		c.Flash.Error("You are not admin")
		return c.Redirect(App.Login)
	}
	return nil
}

// Order 메서드와 Find 메서드로 전체 포스트 조회
func (c Post) Index() revel.Result {
	var posts []models.Post
	c.Txn.Order("created_at desc").Find(&posts)
	return c.Render(posts)
}

// First 메서드로 id에 해당하는 포스트 조회
func (c Post) Show(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	c.Txn.Where(&models.Comment{PostId: id}).Find(&post.Comments)
	return c.Render(post)
}

// Save 메서드로 포스트 수정
func (c Post) Update(id int, title, body string) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	post.Title = title
	post.Body = body

	c.Txn.Save(&post)
	
	// 뷰에 Flash 메시지 전달
	c.Flash.Success("포스트 수정 완료")

	// 포스트 상세 보기 화면으로 이동
	return c.Redirect(routes.Post.Show(id))
}

// Create 메서드로 포스트 생성
func (c Post) Create(title, body string) revel.Result {
	post := models.Post{Title: title, Body: body}
	c.Txn.Create(&post)
	
	// 뷰에 Flash 메시지 전달
	c.Flash.Success("포스트 작성 완료")

	// 포스트 목록 화면으로 이동
	return c.Redirect(routes.Post.Index())
}

// Delete 메서드로 포스트 삭제
func (c Post) Destroy(id int) revel.Result {
	c.Txn.Where("post_id = ?", id).Delete(&models.Comment{})
	c.Txn.Where("id = ?", id).Delete(&models.Post{})
	
	// 뷰에 Flash 메시지 전달
	c.Flash.Success("포스트 삭제 완료")

	// 포스트 목록 화면으로 이동
	return c.Redirect(routes.Post.Index())
}

func (c Post) New() revel.Result {
	post := models.Post{}
	return c.Render(post)
}

func (c Post) Edit(id int) revel.Result {
	var post models.Post
	c.Txn.First(&post, id)
	return c.Render(post)
}
