package service

import (
	"Rent-And-Buy-App/internal/dtos/postDtos"
	"Rent-And-Buy-App/internal/dtos/productDtos"
	"Rent-And-Buy-App/internal/entity"
	"Rent-And-Buy-App/internal/repository"
	"Rent-And-Buy-App/pkg/Converter"
	"errors"
)

type PostService struct {
	repo        *repository.PostRepository
	prodService *ProductService
	userService *UserService
}

func NewPostService(repo *repository.PostRepository, prodService *ProductService, userService *UserService) *PostService {
	return &PostService{repo: repo, prodService: prodService, userService: userService}
}

func (ps *PostService) GetAll() ([]*postDtos.PostDto, error) {
	posts, err := ps.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var PostDtos []*postDtos.PostDto

	for _, post := range posts {
		productId := Converter.UintToString(post.ProductId)
		productDto, err := ps.prodService.GetById(productId)
		if err != nil {
			return nil, err
		}

		postDto, err := ps.getPostDto(post, productDto)
		if err != nil {
			return nil, err
		}

		PostDtos = append(PostDtos, postDto)
	}

	return PostDtos, err
}

func (ps *PostService) GetById(id string) (*postDtos.PostDto, error) {
	Id := Converter.StringToUint(id)
	post, err := ps.repo.GetById(Id)
	if err != nil {
		return nil, err
	}

	productId := Converter.UintToString(post.ProductId)
	production, err := ps.prodService.GetById(productId)
	if err != nil {
		return nil, err
	}

	postDto, err := ps.getPostDto(post, production)
	if err != nil {
		return nil, err
	}
	return postDto, err
}

func (ps *PostService) GetEntityById(id string) (*entity.Post, error) {
	Id := Converter.StringToUint(id)
	post, err := ps.repo.GetById(Id)
	if err != nil {
		return nil, err
	}

	return post, err
}

// TODO isSeller check will be done in middleware in next patch
func (ps *PostService) Create(post *entity.Post, createProductDto *productDtos.CreateProductDto) (*postDtos.PostDto, error) {
	sellerId := Converter.UintToString(post.SellerId)
	user, err := ps.userService.GetByID(sellerId)
	if err != nil {
		return nil, err
	}
	// IsSeller False
	if !*user.IsSeller {
		return nil, errors.New("Only seller can create posts")
	}

	if post.Price < 0 {
		return nil, errors.New("You cannot set the price below zero")
	}

	newProductDto, err := ps.prodService.Create(createProductDto)
	if err != nil {
		return nil, err
	}

	post.ProductId = newProductDto.ID
	err = ps.repo.Create(post)

	postDto, err := ps.getPostDto(post, newProductDto)
	if err != nil {
		return nil, err
	}

	return postDto, nil
}

func (ps *PostService) Update(post *entity.Post, updateDto *postDtos.UpdatePostDto) (*postDtos.PostDto, error) {

	sellerId := Converter.UintToString(post.SellerId)
	user, err := ps.userService.GetByID(sellerId)
	if err != nil {
		return nil, err
	}

	// IsSeller False
	if !*user.IsSeller {
		return nil, errors.New("Only seller can update posts")
	} else if *user.IsSeller && user.ID != post.SellerId {
		return nil, errors.New("This user is not the seller of this post")
	}

	ProductId := Converter.UintToString(post.ProductId)
	product, err := ps.prodService.GetEntityById(ProductId)
	if err != nil {
		return nil, err
	}

	newProduct, err := ps.prodService.Update(product, &updateDto.Product)
	if err != nil {
		return nil, err
	}

	post.SellerId = updateDto.SellerId
	post.Price = updateDto.Price

	err = ps.repo.Update(post)
	if err != nil {
		return nil, err
	}

	postDto, err := ps.getPostDto(post, newProduct)
	return postDto, err
}

func (ps *PostService) Delete(post *entity.Post) error {
	sellerId := Converter.UintToString(post.SellerId)
	product, err := ps.prodService.GetEntityById(sellerId)
	if product == nil {
		return errors.New("Product not found")
	}

	err = ps.prodService.Delete(product)
	err = ps.repo.Delete(post)
	return err
}

func (ps *PostService) getPostDto(post *entity.Post, productDto *productDtos.ProductDto) (*postDtos.PostDto, error) {
	if productDto == nil {
		return nil, errors.New("You must to provide a product dto")
	}

	//This part might be totally useless. Might delete for optimization
	productId := Converter.UintToString(post.ProductId)
	if isExist, err := ps.prodService.GetById(productId); err != nil || isExist == nil {
		return nil, errors.New("Product not found")
	}

	postDto := &postDtos.PostDto{
		ID:       post.Model.ID,
		SellerId: post.SellerId,
		Product:  *productDto,
		Price:    post.Price,
	}

	return postDto, nil
}
