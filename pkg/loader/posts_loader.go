package loader

import (
	"bytes"
	"fmt"
	"github.com/eliasmeireles/go-pdf-generator/pkg/web"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	webProviderHost = getWebProviderHost() // Default value if env is not set
)

type TemplateData struct {
	WebProviderHost string
	PostsData       template.JS // Use template.JS to prevent escaping
}

func fetchPostsData() (string, error) {
	resp, err := http.Get(strings.TrimSuffix(webProviderHost, "/") + "/posts.json")
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func RenderHTMLTemplate() (string, error) {
	tmpl, err := template.New("htmlTemplate").Parse(web.Template)
	if err != nil {
		return "", err
	}

	postsData, err := fetchPostsData()
	if err != nil {
		return "", err
	}

	// Use template.JS to ensure the JSON data is not escaped
	data := TemplateData{
		WebProviderHost: webProviderHost,
		PostsData:       template.JS(postsData),
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	htmlContent := buf.String()
	return htmlContent, nil
}

func getWebProviderHost() string {
	host := os.Getenv("WEB_PROVIDER_HOST")
	if host == "" {
		host = "http://localhost:3000"
	}
	log.Print("Web provider host: " + host)
	return host
}
