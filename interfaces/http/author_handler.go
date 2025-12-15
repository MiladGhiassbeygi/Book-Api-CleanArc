package http

import (
	"net/http"
	"strconv"

	"book-api-cleanarc/internal/command"
	"book-api-cleanarc/internal/query"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	Command *command.AuthorCommandService
	Query   *query.AuthorQueryService
}

func NewAuthorHandler(cmd *command.AuthorCommandService, qry *query.AuthorQueryService) *AuthorHandler {
	return &AuthorHandler{Command: cmd, Query: qry}
}

func (h *AuthorHandler) RegisterRoutes(r *gin.Engine) {
	authors := r.Group("/authors")
	{
		authors.GET("", h.GetAll)
		authors.GET("/:id", h.GetByID)
		authors.POST("", h.Create)
		authors.DELETE("/:id", h.Delete)
	}
}

// @Summary Get all authors
// @Description Get list of all authors
// @Tags authors
// @Produce json
// @Success 200 {array} domain.Author
// @Failure 500 {object} map[string]string
// @Router /authors [get]
func (h *AuthorHandler) GetAll(c *gin.Context) {
	authors, err := h.Query.GetAllAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, authors)
}

// @Summary Get author by ID
// @Description Get a single author by ID
// @Tags authors
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} domain.Author
// @Failure 404 {object} map[string]string
// @Router /authors/{id} [get]
func (h *AuthorHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)

	author, err := h.Query.GetAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}

// @Summary Create author
// @Description Create a new author
// @Tags authors
// @Accept json
// @Produce json
// @Param author body domain.Author true "Author"
// @Success 201 {object} domain.Author
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /authors [post]
func (h *AuthorHandler) Create(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author, err := h.Command.CreateAuthor(input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, author)
}

// @Summary Delete author
// @Description Delete an author by ID
// @Tags authors
// @Param id path int true "Author ID"
// @Success 200 {object} domain.Author
// @Failure 404 {object} map[string]string
// @Router /authors/{id} [delete]
func (h *AuthorHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 64)

	author, err := h.Command.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, author)
}
