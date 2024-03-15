package main

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"haseeb.khan/email"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

type Template struct {
	TemplateId   int    `gorm:"primarykey" json:"id"`
	TemplateData string `json:"data"`
}

func InitDB() {
	dsn := "host=localhost user=postgres password=root dbname=empdb port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect to DB \ncause: %v", err)
	}
	db.AutoMigrate(&Template{})
	log.Println("connection sucess")
}

func LoadTemplateFromDB(id int64) (Template, error) {
	var template Template
	if err := db.First(&template, id); err.Error != nil {
		log.Printf("Unable to find template with id: %v\ncause: %v", id, err)
		return Template{}, err.Error
	}
	return template, nil
}

type TemplateRequest struct {
	TemplateId int64  `json:"template_id"`
	Receiver   string `json:"receiver"`
	Variables  struct {
		Otp string `json:"otp"`
	} `json:"vars"`
}

func LoadEnvironmentVariables() (string, string, string, string) {
	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Unable to load environment file.")
	}

	from := os.Getenv("FROM")
	password := os.Getenv("PASS")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	return from, password, host, port
}

func SendEmail(ctx iris.Context) {
	var request TemplateRequest
	err = ctx.ReadJSON(&request)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Invalid request body.", "cause": err.Error()})
		return
	}

	var email_body bytes.Buffer

	templ, err := LoadTemplateFromDB(request.TemplateId)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"cause": err})
		return
	}

	html := templ.TemplateData
	t := template.Must((template.New("otp").Parse(html)))
	err = t.Execute(&email_body, request)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject:One Time Password " + "\n" + mime + "\n"
	email := append([]byte(subject), email_body.Bytes()...)

	if err != nil {
		log.Println(err)
		ctx.StopWithJSON(iris.StatusInternalServerError, iris.Map{"message": err})
		return
	}

	from, password, host, port := LoadEnvironmentVariables()

	toList := []string{request.Receiver}

	auth := smtp.PlainAuth("", from, password, host)
	err = smtp.SendMail(host+":"+port, auth, from, toList, email)

	if err != nil {
		log.Println(err.Error())
		ctx.StopWithJSON(iris.StatusBadRequest, iris.Map{"error": err})
		return
	}

	log.Println("Mail sent successfully !")

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "email sent"})
}

func main() {
	//InitDB()
	response, err := email.SendEmail()
	if err != nil {
		return
	}
	log.Println(response)

	//server := iris.Default()
	//
	//server.Post("/sendemail", SendEmail)
	//
	//server.Run(iris.Addr(":8080"))
}
