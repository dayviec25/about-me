package handlers

import (
	"log"
	"net/http"

	"about-me/blog"

	"github.com/labstack/echo/v4"
)

// RegisterBlogHandlers wires up the blog index and individual post pages.
func RegisterBlogHandlers(e *echo.Echo) {
	log.Println("Registering blog handlers")
	e.GET("/blog", ListPosts)
	e.GET("/blog/:slug", ShowPost)
}

// ListPosts renders the blog index with every post, newest first.
func ListPosts(c echo.Context) error {
	posts, err := blog.LoadPosts()
	if err != nil {
		log.Printf("Error loading blog posts: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.Render(http.StatusOK, "blog-list.tmpl", posts)
}

// ShowPost renders a single post, or 404s if the slug is unknown.
func ShowPost(c echo.Context) error {
	post, err := blog.GetPost(c.Param("slug"))
	if err != nil {
		log.Printf("Error loading blog post: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if post == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}
	return c.Render(http.StatusOK, "blog-post.tmpl", post)
}
