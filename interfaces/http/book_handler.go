package http

import (
	"net/http"
	"strconv"

	"book-api-cleanarc/internal/command"
	"book-api-cleanarc/internal/query"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	Command *command.BookCommandService
	Query   *query.BookQueryService
}

func NewBookHandler(cmd *command.BookCommandService, qry *query.BookQueryService) *BookHandler {
	return &BookHandler{Command: cmd, Query: qry}
}
func (h *BookHandler) RegisterRoutes(r *gin.Engine) {
	books := r.Group("/books")
	{
		books.GET("", h.GetAll)
		books.GET("/:id", h.GetByID)
		books.POST("", h.Create)
	}

}

// @Summary Get all books
// @Description Get list of all books
// @Tags books
// @Produce json
// @Success 200 {array} domain.Book
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (h *BookHandler) GetAll(c *gin.Context) {
	books, err := h.Query.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary Get book by ID
// @Description Get a single book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} domain.Book
// @Failure 404 {object} map[string]string
// @Router /books/{id} [get]
func (h *BookHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)

	book, err := h.Query.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// @Summary Get books by author
// @Description Get all books for a specific author
// @Tags books
// @Produce json
// @Param authorID path int true "Author ID"
// @Success 200 {array} domain.Book
// @Failure 500 {object} map[string]string
// @Router /authors/{authorID}/books [get]
func (h *BookHandler) GetByAuthor(c *gin.Context) {
	authorParam := c.Param("authorID")
	authorID, _ := strconv.ParseUint(authorParam, 10, 64)

	books, err := h.Query.GetBooksByAuthor(uint(authorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Summary Create book
// @Description Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body domain.Book true "Book"
// @Success 201 {object} domain.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (h *BookHandler) Create(c *gin.Context) {
	var input struct {
		Title    string `json:"title"`
		AuthorID uint   `json:"author_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.Command.CreateBook(input.Title, input.AuthorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}
