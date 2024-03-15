
# Go Emailer with HTML Template

This Go application demonstrates how to send HTML emails dynamically populated with data from a database using Go templates.
## Features

- Connects to a Postgres database to fetch template data.
- Generates an HTML email template using Go's template package.
- Sends the email using SMTP.
- Loaded Email and DB Credentials from Environment File.


## Setup
1. Install Go (if not already installed) from [golang.org](golang.org).
2. Clone this repository:
```http
  git clone https://github.com/HSB086/GO-REPO.git
```
3. Install required dependencies:
```http
  go mod tidy
```
4. Configure the database connection in main.go and email settings in .env file.
5. Run the applicaion:
```http
  go run main.go
```

## Configuration
Make sure to update the following configuration parameters:

- main.go:  `username`, `password`, `host`, `port`, and `dbname`.
- .env file: `FROM`, `PASS`, `HOST` and `PORT`
