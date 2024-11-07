package irispkg

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// Book example.
type Book struct {
	Title string `json:"title"`
}

type BookController struct {
	/* dependencies */
}

// GET: http://localhost:8080/books
func (c *BookController) Get() []Book {
	return []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}
}

// POST: http://localhost:8080/books
func (c *BookController) Post(b Book) int {
	println("Received Book: " + b.Title)

	return iris.StatusCreated
}

// Get /books/:id
func (c *BookController) GetBy(ctx iris.Context, id int64) int {
	println("url path:", ctx.Path())
	println("url:", ctx.Request().URL.String())
	println("id:", id)
	return iris.StatusOK
}

// GET /books/:id/title/:name
func (c *BookController) GetByTitleBy(ctx iris.Context, id int64, name string) int {
	println("2url path:", ctx.Path())
	println("2url:", ctx.Request().URL.String())
	println("2id:", id)
	println("2name:", name)
	return iris.StatusOK
}

// GET /books/:id/title
func (c *BookController) GetByTitle(ctx iris.Context, id int64) int {
	println("3url path:", ctx.Path())
	println("3url:", ctx.Request().URL.String())
	println("3id:", id)
	return iris.StatusOK
}

func StartIrisApp() {
	app := iris.New()
	booksAPI := app.Party("/books")
	m := mvc.New(booksAPI)
	m.Handle(new(BookController))
	app.Listen(":8080")
}
