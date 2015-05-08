package main

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/jacobhe/nova-vote/auth"
	"github.com/macaron-contrib/binding"
	"github.com/macaron-contrib/session"
	"log"
	"net/http"
)

func main() {
	m := macaron.Classic()
	m.Use(session.Sessioner())
	m.Use(macaron.Renderer())

	authorize := auth.Authorize()

	m.Get("/", indexHandler)
	m.Post("/login", binding.Bind(auth.User{}), auth.Authenticate)
	m.Group("/vote", func() {
		m.Post("", binding.Bind(VoteForm{}), voteFormHandler)
	}, authorize)

	m.Run()

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

type Company struct {
	Id   int
	Name string
}

type VoteForm struct {
	Id int `form:"Id" binding:"Required"`
}

func indexHandler(ctx *macaron.Context) {
	var companies = make([]Company, 3)
	var c1 = Company{Id: 1, Name: "c1"}
	var c2 = Company{Id: 2, Name: "c2"}
	var c3 = Company{Id: 3, Name: "c3"}
	companies[0] = c1
	companies[1] = c2
	companies[2] = c3
	ctx.Data["companies"] = companies
	ctx.HTML(200, "index") // 200 为响应码
}

func voteFormHandler(vote VoteForm) string {
	return fmt.Sprintf("Id: %d", vote.Id)
}
