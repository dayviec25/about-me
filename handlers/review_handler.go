package handlers

import (
	"about-me/data"
	"about-me/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// RegisterReviewHandlers sets up the routes for handling review navigation
func RegisterReviewHandlers(e *echo.Echo) {
	log.Println("Registering review handlers")
	e.GET("/review/:index/prev", GetReview("prev"))
	e.GET("/review/:index/next", GetReview("next"))
	e.GET("/review/initial", GetInitialReview)
}

func GetReview(direction string) echo.HandlerFunc {
	return func(c echo.Context) error {
		aboutMeData, err := data.GetAboutMeData()
		if err != nil {
			log.Printf("Error getting About Me data: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		index, _ := strconv.Atoi(c.Param("index"))
		reviewCount := len(aboutMeData.Reviews)
		if reviewCount == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "No reviews available")
		}

		// Ensure index is within the valid range
		index = (index + reviewCount) % reviewCount

		reviewData := model.ReviewNavigation{
			Review:        &aboutMeData.Reviews[index],
			PrevID:        (index - 1 + reviewCount) % reviewCount,
			NextID:        (index + 1) % reviewCount,
			IsFirstReview: index == 0,
			IsLastReview:  index == reviewCount-1,
		}

		log.Printf("Serving review %d", index)
		return c.Render(http.StatusOK, "review-partial.tmpl", reviewData)
	}
}

func GetInitialReview(c echo.Context) error {
	log.Println("Serving initial review")
	aboutMeData, err := data.GetAboutMeData()
	if err != nil {
		log.Printf("Error getting About Me data: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	reviewCount := len(aboutMeData.Reviews)
	if reviewCount == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "No reviews available")
	}

	reviewData := model.ReviewNavigation{
		Review:        &aboutMeData.Reviews[0],
		PrevID:        reviewCount - 1,
		NextID:        1,
		IsFirstReview: true,
		IsLastReview:  reviewCount == 1,
	}
	return c.Render(http.StatusOK, "review-partial.tmpl", reviewData)
}
