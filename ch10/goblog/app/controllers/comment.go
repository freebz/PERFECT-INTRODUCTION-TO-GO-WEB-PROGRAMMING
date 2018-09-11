package controllers

import (
	"goblog/app/routes"
	"time"

	"github.com/revel/modules/db/app"
	"github.com/revel/revel"
)

type Comment struct {
	*revel.Controller
	db.Transactional
}

func (c Comment) Create(postId int, body, commenter string) revel.Result {
	// 댓글 저장
	_, err := c.Txn.Exec("insert into comments(body, commenter, post_id, created_at, updated_at) values(?,?,?,?,?)", body, commenter, postId, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}

	// 뷰에 Flash 메시지 전달
	c.Flash.Success("댓글 작성 완료")

	// 포스트 상세 보기 화면으로 이동
	return c.Redirect(routes.Post.Show(postId))
}

func (c Comment) Destroy(postId, id int) revel.Result {
	// 댓글 삭제
	if _, err := c.Txn.Exec("delete from comments where id=?", id); err != nil {
		panic(err)
	}

	// 뷰에 Flash 메시지 전달
	c.Flash.Success("댓글 삭제 완료")

	// 포스트 상세 보기 화면으로 이동
	return c.Redirect(routes.Post.Show(postId))
}
