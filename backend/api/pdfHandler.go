package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Pdf struct {
	Name string `json:"name"`
	DownloadUrl string `json:"download_url"`
}

const baseGithubApiUrl string = "https://api.github.com"

func pdfHandler(w http.ResponseWriter, req *http.Request) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}


	githubToken := os.Getenv("GITHUB_TOKEN")

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")

	pdfsUrl := fmt.Sprintf("%s/repos/kyotosplit/books/contents/pdfs?ref=main", baseGithubApiUrl)

	req, err = http.NewRequest(http.MethodGet, pdfsUrl, nil)
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

	json.NewEncoder(w).Encode(pdfs)
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}

	const port string = ":8000"

	http.HandleFunc("/", pdfHandler)
	fmt.Printf("Listening on http://localhost%s", port)

	err = http.ListenAndServe(port, nil)	

	if err != nil {
		log.Fatalf("Error ocurred trying to start the server: %s", err)
	}
}
