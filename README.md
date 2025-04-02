# PDF Generation Service

This service generates PDFs from HTML content served by a web provider. It uses a headless Chrome instance to render the
HTML and convert it to a PDF.

---

## API Endpoint

### Generate PDF from URL

**Endpoint**: `GET` -> [/api/pdf-generator/v1/pdf](http://localhost:8080/api/pdf-generator/v1/pdf)

**Description**: Generates a PDF from the HTML content at the specified URL.

**Query Parameters**:

- `url` (required): The URL of the HTML content to convert to PDF.
- `class` (optional): A list of element classes to wait for before generating the PDF. The service ensures these
  elements are visible before proceeding.
- `id` (optional): A list of element IDs to wait for before generating the PDF. The service ensures these elements are
  visible before proceeding.
- `fileName` (optional): The desired name of the generated PDF file. If not provided, an uuid will be used.
- `appendText` (optional): Append the provided text to the PDF. Nice to test with parallel execution.

**Response**:

- A PDF file is returned as a binary stream with the `Content-Disposition` header set for download.

---

## Example `curl` Command

To generate a PDF from a URL:

```bash
mkdir .out
curl -o ./.out/output.pdf "http://localhost:8080/api/pdf-generator/v1/pdf?url=https://go.dev/doc"
```

You can also pass the query param url as a base64 just provided a new query param `base64=true`

```bash
mkdir .out
curl -o ./.out/output.pdf "http://localhost:8080/api/pdf-generator/v1/pdf?url=aHR0cHM6Ly9nby5kZXYvZG9j&base64=true"
```

To generate a PDF and wait for specific elements to be visible:

```bash
mkdir .out
curl -o ./.out/output.pdf "http://localhost:8080/api/pdf-generator/v1/pdf?url=https://go.dev/doc/tutorial/getting-started&id=prerequisites&id=nav"
```

---

## Docker Compose Setup

The service is designed to run in a Docker container alongside a web provider service (e.g., Nginx) that serves the HTML
content.

### Services

1. **`web-provider-app`**:
    - Serves static HTML files.
    - Exposes port `3000` on the host.

2. **`pdf-generator-app`**:
    - Runs the PDF generation service.
    - Exposes port `8080` for the API and port `9223` for Chrome's remote debugging.

### Start the Services

1. Build and start the services:
   ```bash
   docker-compose up --build
   ```

2. Access the services:
    - Web provider: `http://localhost:3000`
    - PDF generator API: `http://localhost:8080/api/pdf-generator/v1/pdf`

---


