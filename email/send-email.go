package email

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type EmailRequest struct {
	From struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	ReplyTo    string   `json:"reply_to"`
	Subject    string   `json:"subject"`
	TemplateId int      `json:"template_id"`
	Tags       []string `json:"tags"`
	Content    []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"content"`
	Attachments []struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	} `json:"attachments"`
	Personalizations []struct {
		Attributes struct {
			NAME string `json:"NAME"`
			OTP  string `json:"OTP"`
		} `json:"attributes"`
		To []struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"to"`
		Cc []struct {
			Email string `json:"email"`
		} `json:"cc"`
		Bcc []struct {
			Email string `json:"email"`
		} `json:"bcc"`
		TokenTo     string `json:"token_to"`
		TokenCc     string `json:"token_cc"`
		TokenBcc    string `json:"token_bcc"`
		Attachments []struct {
			Name    string `json:"name"`
			Content string `json:"content"`
		} `json:"attachments"`
	} `json:"personalizations"`
	Settings struct {
		OpenTrack        bool   `json:"open_track"`
		ClickTrack       bool   `json:"click_track"`
		UnsubscribeTrack bool   `json:"unsubscribe_track"`
		IpPool           string `json:"ip_pool"`
	} `json:"settings"`
	Bcc []struct {
		Email string `json:"email"`
	} `json:"bcc"`
	Schedule int `json:"schedule"`
}

type EmailResponse struct {
	Data struct {
		MessageId string `json:"message_id"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func LoadFromEnv() (string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env file")
	}
	apiUrl := os.Getenv("API_URL")
	apiKey := os.Getenv("API_KEY")

	return apiUrl, apiKey
}

func GetHTML() (string, error) {
	html, err := os.ReadFile("index.html")
	if err != nil {
		log.Println("Unable to load HTML file")
		return "", err
	}
	return string(html), nil
}

func SendEmail() (string, error) {
	var emailResponse EmailResponse
	apiUrl, apiKey := LoadFromEnv()

	// html, err := GetHTML()
	// if err != nil {
	// 	return "", err
	// }

	request := EmailRequest{
		From: struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}{
			Email: "noreply@heytorus.com",
			Name:  "Torus",
		},
		ReplyTo:    "",
		Subject:    "Testing Email Sending through Netcore API",
		TemplateId: 48616,
		Tags: []string{
			"test",
		},
		Content: []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		}{
			{
				Type:  "html",
				Value: "<p></p>",
			},
		},
		Attachments: []struct {
			Name    string `json:"name"`
			Content string `json:"content"`
		}{
			{
				Name:    "",
				Content: "",
			},
		},
		Personalizations: []struct {
			Attributes struct {
				NAME string `json:"NAME"`
				OTP  string `json:"OTP"`
			} `json:"attributes"`
			To []struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"to"`
			Cc []struct {
				Email string `json:"email"`
			} `json:"cc"`
			Bcc []struct {
				Email string `json:"email"`
			} `json:"bcc"`
			TokenTo     string `json:"token_to"`
			TokenCc     string `json:"token_cc"`
			TokenBcc    string `json:"token_bcc"`
			Attachments []struct {
				Name    string `json:"name"`
				Content string `json:"content"`
			} `json:"attachments"`
		}{
			{
				Attributes: struct {
					NAME string `json:"NAME"`
					OTP  string `json:"OTP"`
				}{
					NAME: "Kapil Kuthe",
					OTP:  "131313",
				},
				To: []struct {
					Email string `json:"email"`
					Name  string `json:"name"`
				}{
					{
						Email: "haseeb.khan@heytorus.com",
						Name:  "Haseeb Khan",
					},
				},
				Cc: []struct {
					Email string `json:"email"`
				}{},
				Bcc: []struct {
					Email string `json:"email"`
				}{},
				TokenTo:  "",
				TokenCc:  "",
				TokenBcc: "",
				Attachments: []struct {
					Name    string `json:"name"`
					Content string `json:"content"`
				}{},
				//Attachments: nil,
			},
		},
		Settings: struct {
			OpenTrack        bool   `json:"open_track"`
			ClickTrack       bool   `json:"click_track"`
			UnsubscribeTrack bool   `json:"unsubscribe_track"`
			IpPool           string `json:"ip_pool"`
		}{OpenTrack: true, ClickTrack: true, UnsubscribeTrack: false, IpPool: "shared"},
		Bcc: []struct {
			Email string `json:"email"`
		}{{Email: ""}},
		Schedule: 0,
	}

	requestBody, err := json.Marshal(request)
	if err != nil {
		log.Println("Unable to convert request to JSON.\ncause: " + err.Error())
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("Unable to create request.\ncause: " + err.Error())
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", apiKey)

	response, err := client.Do(req)
	if err != nil {
		log.Println("Unable to call api.\ncause: " + err.Error())
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading response body.\ncause: " + err.Error())
		return "", err
	}

	if err := json.Unmarshal(body, &emailResponse); err != nil {
		log.Println(err)
		return "", err
	}

	return string(body), nil
}
