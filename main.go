package main

import (
	"strconv"

	"github.com/gofiber/fiber"
)

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos = []*Todo{
	{ID: 1, Name: "Walk the dog", Completed: false},
	{ID: 2, Name: "Walk the cat", Completed: false},
	{ID: 3, Name: "Walk the snake", Completed: false},
	{ID: 4, Name: "Walk the big cat", Completed: false},
}

var posts = []*Post{
	{ID: 1, Title: "Post 1", Description: "Post 1 Description"},
	{ID: 2, Title: "Post 2", Description: "Post 2 Description"},
	{ID: 3, Title: "Post 3", Description: "Post 3 Description"},
	{ID: 4, Title: "Post 4", Description: "Post 4 Description"},
}

func GetTodos(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(todos)
}
func GetPosts(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(posts)
}

// create

func CreateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name string `json:"name"`
	}

	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}
	todo := &Todo{
		ID:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}
	todos = append(todos, todo)
	ctx.Status(fiber.StatusCreated).JSON(todo)
}

func CreatePost(ctx *fiber.Ctx) {
	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	var body request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}

	post := &Post{
		ID:          len(posts) + 1,
		Title:       body.Title,
		Description: body.Description,
	}
	posts = append(posts, post)
	ctx.Status(fiber.StatusCreated).JSON(post)
}

func GetTodo(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}
	for _, todo := range todos {
		if todo.ID == id {
			ctx.Status(fiber.StatusOK).JSON(todo)
			return
		}
	}
	ctx.Status(fiber.StatusNotFound)
}

func GetPost(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}
	for _, post := range posts {
		if post.ID == id {
			ctx.Status(fiber.StatusOK).JSON(post)
			return
		}
	}
	ctx.Status(fiber.StatusNotFound)
}
func DeleteTodo(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[0:i], todos[i+1:]...)
			ctx.Status(fiber.StatusNoContent)
			return
		}
	}
	ctx.Status(fiber.StatusNotFound)
}
func DeletePost(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}
	for i, post := range posts {
		if post.ID == id {
			posts = append(posts[0:i], posts[i+1:]...)
			ctx.Status(fiber.StatusNoContent)
			return
		}
	}
	ctx.Status(fiber.StatusNotFound)
}

func UpdateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}
	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse bdy",
		})
		return
	}
	var todo *Todo
	for _, t := range todos {
		if t.ID == id {
			todo = t
			break
		}
	}
	if todo.ID == 0 {
		ctx.Status(fiber.StatusNotFound)
	}
	if body.Name != nil {
		todo.Name = *body.Name
	}
	if body.Completed != nil {
		todo.Completed = *body.Completed
	}
	ctx.Status(fiber.StatusOK).JSON(todo)
}
func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("hello world")
	})

	SetupRoutes(app)

	err := app.Listen("4000")
	if err != nil {
		panic(err)
	}
}

func SetupRoutes(app *fiber.App) {
	todoRoutes := app.Group("/todos")
	todoRoutes.Get("/", GetTodos)
	todoRoutes.Get("/:id", GetTodo)
	todoRoutes.Patch("/:id", UpdateTodo)
	todoRoutes.Post("/", CreateTodo)
	postRoutes := app.Group("/posts")
	postRoutes.Get("/", GetPosts)
	postRoutes.Get("/:id", GetPost)
	postRoutes.Post("/", CreatePost)

}
