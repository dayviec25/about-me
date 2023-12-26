package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

type HomePage struct {
	Title   string
	Message string
}

// TemplateRegistry is a custom HTML template renderer
type TemplateRegistry struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func createMapFromSlice(slice []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range slice {
		m[v] = true
	}
	return m
}

func isInMap(m map[string]bool, key string) bool {
	return m[key]
}

func main() {
	e := echo.New()
	e.Static("/static", "dist")

	// Initialize the template engine
	templates, err := template.ParseGlob("templates/pages/*.tmpl")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	allowedPages := []string{"index", "home", "upload", "pricing"}
	lookupMap := createMapFromSlice(allowedPages)
	e.GET("/:page", func(c echo.Context) error {
		page := c.Param("page")
		if isInMap(lookupMap, page) {
			return c.Render(http.StatusOK, page+".tmpl", nil)
		}
		return c.NoContent(http.StatusNotFound)
	})

	e.GET("/*", func(c echo.Context) error {
		// Extract page name from URL
		page := strings.TrimPrefix(c.Request().URL.Path, "/pages/")
		page = strings.TrimSuffix(page, ".tmpl")

		// Check if the template exists
		if tmpl := templates.Lookup(page + ".tmpl"); tmpl != nil {
			return c.Render(http.StatusOK, page+".tmpl", nil)
		}
		return c.NoContent(http.StatusNotFound)
	})
	// Read these from environment variables or a secure source
	spaceName := "nyc3"
	spaceRegion := "us-east-1"
	accessKeyID := os.Getenv("DO_STORAGE_ACCESS_KEY")
	secretAccessKey := os.Getenv("DO_STORAGE_SECRET_ACCESS_KEY")

	log.Println("Using DigitalOcean Spaces bucket " + spaceName + " in " + spaceRegion + " region.")
	log.Println("Access Key ID: " + accessKeyID)
	log.Println("Secret Access Key: " + secretAccessKey)

	// AWS session, configured for DigitalOcean Spaces
	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String(spaceRegion),
		Endpoint:         aws.String("https://" + spaceName + ".digitaloceanspaces.com"),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	}))
	uploader := s3manager.NewUploader(sess)

	// Define bucketName (replace with your bucket name)
	bucketName := "upload-ai"

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// Log the error
		log.Printf("Error occurred: %v", err)

		// Handle different types of errors
		if he, ok := err.(*echo.HTTPError); ok {
			// For Echo's HTTP errors
			c.String(he.Code, he.Message.(string))
		} else {
			// For other errors
			c.String(http.StatusInternalServerError, err.Error())
		}
	}

	// Routes
	e.POST("/upload_endpoint", func(c echo.Context) error {
		log.Println("Received request at /upload_endpoint")
		file, err := c.FormFile("file")
		if err != nil {
			log.Println("error getting file")
			return err
		}
		src, err := file.Open()
		if err != nil {
			log.Println("error opening file")
			return err
		}
		defer src.Close()

		// Upload to S3
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(file.Filename),
			Body:   src,
		})
		if err != nil {
			log.Println("error uploading file")
			return err
		}
		log.Println("File uploaded successfully")
		return c.String(http.StatusOK, "File "+file.Filename+" uploaded successfully.")
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
