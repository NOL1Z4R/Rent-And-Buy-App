package handlers

import (
	"Rent-And-Buy-App/internal/dtos/postDtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/service"
	"Rent-And-Buy-App/pkg/Response"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (ph *PostHandler) GetPostAll(c *gin.Context) {
	posts, err := ph.service.GetAll()
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.OK(c, posts)
}

func (ph *PostHandler) GetPostByID(c *gin.Context) {
	id := c.Param("id")
	post, err := ph.service.GetById(id)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}

	Response.OK(c, post)
}

func (ph *PostHandler) CreatePost(c *gin.Context) {
	var createPostDto *postDtos.CreatePostDto
	if err := c.ShouldBindJSON(&createPostDto); err != nil {
		Response.Error(c, err.Error())
		return
	}

	createProductDto := createPostDto.Product

	post := entity.Post{
		SellerId: createPostDto.SellerId,
		Price:    createPostDto.Price,
	}

	postDto, err := ph.service.Create(&post, &createProductDto)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}

	Response.OK(c, postDto)
}

func (ph *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	post, err := ph.service.GetEntityById(id)
	if post == nil || err != nil {
		Response.Error(c, "This post does not exist")
		return
	}

	var updatePostDto *postDtos.UpdatePostDto
	if err := c.ShouldBindJSON(&updatePostDto); err != nil {
		Response.Error(c, err.Error())
		return
	}

	postDto, err := ph.service.Update(post, updatePostDto)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}

	Response.OK(c, postDto)
}

func (ph *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	post, err := ph.service.GetEntityById(id)
	if err != nil {
		Response.Error(c, err.Error())
		return
	}
	if err := ph.service.Delete(post); err != nil {
		Response.Error(c, err.Error())
		return
	}
	Response.OK(c, nil)
}
