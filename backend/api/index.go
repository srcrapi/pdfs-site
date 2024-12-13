package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Pdf struct {
	Name string `json:"name"`
	DownloadUrl string `json:"download_url"`
}

const baseGithubApiUrl string = "https://api.github.com"

func pdfHandler(c *gin.Context) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}


	githubToken := os.Getenv("GITHUB_TOKEN")

	pdfsUrl := fmt.Sprintf("%s/repos/kyotosplit/books/contents/pdfs?ref=main", baseGithubApiUrl)

	req, err := http.NewRequest(http.MethodGet, pdfsUrl, nil)
	if err != nil {
		log.Fatalf("client: failed to create a new request: %s", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", githubToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("client: failed to send request: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("client: could not read response body: %s", err)
	}

	var pdfs []Pdf

	err = json.Unmarshal(body, &pdfs)
	if err != nil {
		log.Fatalf("server: failed to parse json data: %s", err)
	}

	// json.NewEncoder(w).Encode(pdfs)
	c.JSON(http.StatusOK, pdfs)
}

/*
func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}

	const port string = ":8000"

	http.HandleFunc("/", Handler)
	fmt.Printf("Listening on http://localhost%s", port)

	err = http.ListenAndServe(port, nil)	

	if err != nil {
		log.Fatalf("Error ocurred trying to start the server: %s", err)
	}
}
*/

func CORSMiddlesware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func Handler(c *gin.Context) {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	app.Use(CORSMiddlesware())
	app.GET("/api/pdfs", pdfHandler)

	app.ServeHTTP(c.Writer, c.Request)
}
