package main

import (
    "html/template"
    "net/http"
    "os"
    "io"
    "log"
    
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

// Template structure for Echo
type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    e := echo.New()
    e.Static("/static", "dist")

    // Initialize the template engine
    tmpl, err := template.ParseFiles("index.tmpl")
    if err != nil {
        panic(err)
    }
    e.Renderer = &Template{
        templates: tmpl,
    }

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
    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "pages/index.html", HomePage{Title: "Home Page", Message: "Hello World"})
    })

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
