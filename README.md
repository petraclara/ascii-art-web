ASCII Web Server (Go) - Documentation
=====================================

Overview
--------
This Go application is a simple web server that allows users to input text and generate ASCII art using different banners. It demonstrates:
- HTTP GET and POST handling with net/http
- HTML template rendering with html/template
- Form parsing and input validation
- Error handling for 404, 400, and 500 HTTP statuses
- Line-by-line ASCII art generation

The server runs locally and serves HTML pages for input and output.

Installation
------------
1. Clone the repository or download the files.
2. Make sure Go is installed (version 1.18+ recommended).
3. Place your HTML templates in the `templates/` directory:
   - `index.html` → main page with input form
   - `badrequest.html` → shown for invalid URLs
4. Implement or provide the `AsciiArt(line, banner)` function for generating ASCII art from text lines.
5. Run the server:

   go run main.go

The server will start on: http://localhost:8080

Endpoints
---------
1. GET `/`
   - Serves the home page where users can input text and choose a banner.
   - Only allows GET requests. Returns 400 Bad Request if the method is not GET.
   - Returns `badrequest.html` if URL path is not exactly `/`.
   - Loads `index.html` template and displays the form.

2. POST `/ascii-art`
   - Processes user-submitted text to generate ASCII art.
   - Only allows POST requests; returns 400 Bad Request otherwise.
   - Parses form data. Returns 400 Bad Request if parsing fails.
   - Expects:
     - `text` → user input text
     - `banner` → ASCII banner type
   - Returns 400 Bad Request if `text` or `banner` is missing.
   - Processes the text line by line with `AsciiArt(line, banner)`.
   - Converts spaces to `&nbsp;` and newlines to `<br>` for HTML display.
   - Loads `index.html` and passes processed ASCII art to the template.

Error Handling
--------------
- 404 Not Found: Served when URL path is invalid.
- 400 Bad Request: Triggered for invalid methods, missing inputs, or failed form parsing.
- 500 Internal Server Error: Catches template parsing failures or ASCII generation errors.

Templates
---------
1. `index.html`
   - Should include:
     - A form that POSTs to `/ascii-art`
     - Input field for `text`
     - Dropdown or input for `banner`
     - Section to display ASCII art using `{{.Result}}`

2. `badrequest.html`
   - Simple page informing the user that the requested URL or request method is invalid.

ASCII Art Function
------------------
Function signature:

AsciiArt(input string, banner string) (string, error)

- Takes a single line of text and a banner type.
- Returns the ASCII-art representation of that line.
- Returns an error if generation fails.
- Server iterates over all lines in the user input and concatenates the results.

Code Structure
--------------
main.go
templates/
  index.html
  badrequest.html

- main() → starts the server and sets up routes
- homeHandler → serves `/`
- asciiArtHandler → handles `/ascii-art`
- AsciiArt() → generates ASCII representation
- Helper logic in handlers for:
  - Validating method
  - Parsing forms
  - Handling line breaks
  - Error responses

Usage Tips
----------
- Always use POST when submitting text for ASCII conversion.
- Ensure your `templates` folder exists and templates are correctly named.
- For multi-line text input, the handler replaces `\r\n` and `\r` with `\n` for consistent processing.
- Optionally, expand the server to store ASCII art in sessions.

Example Flow
------------
1. Navigate to `http://localhost:8080/`
2. Enter text: `Hello Go`
3. Select banner: `standard`
4. Submit the form → POST `/ascii-art`
5. Server returns the home page with ASCII art rendered in the result section.