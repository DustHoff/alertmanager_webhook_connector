package znoony

import (
	"OTRSAlertmanagerHook/logging"
	"OTRSAlertmanagerHook/ticketsystem"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Znoony struct {
	httpClient http.Client
	config     *ticketsystem.Config
}

type TicketResponse struct {
	Error Error `json:"Error"`
}

type Error struct {
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
}
type TicketRequest struct {
	UserLogin string        `json:"UserLogin"`
	Password  string        `json:"Password"`
	Article   TicketArticle `json:"Article"`
	Ticket    Ticket        `json:"Ticket"`
}
type Ticket struct {
	Title        string `json:"Title"`
	State        string `json:"State"`
	Queue        string `json:"Queue"`
	Priority     string `json:"Priority"`
	CustomerUser string `json:"CustomerUser"`
}
type TicketArticle struct {
	ArticleSend bool   `json:"ArticleSend"`
	ContentType string `json:"ContentType"`
	Subject     string `json:"Subject"`
	Body        string `json:"Body"`
}

func NewZnoonyTicket(queue string, subject string, body string) TicketRequest {
	return TicketRequest{
		Ticket: Ticket{
			Queue:        queue,
			State:        "open",
			Title:        subject,
			Priority:     "5 very high",
			CustomerUser: "dho",
		},
		Article: TicketArticle{
			ContentType: "text/plain; charset=utf8",
			ArticleSend: false,
			Subject:     subject,
			Body:        body,
		},
	}
}

func NewZnoonyClient(config *ticketsystem.Config) Znoony {
	insecureSkipVerify, err := strconv.ParseBool(config.Properties["InsecureSkipVerify"])
	if err != nil {
		logging.Error(err)
		insecureSkipVerify = false
	}
	return Znoony{
		httpClient: http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecureSkipVerify,
				},
			},
		},
		config: config,
	}
}

func (z Znoony) CreateTicket(ticket TicketRequest) (TicketResponse, error) {
	ticket.UserLogin = z.config.Properties["username"]
	ticket.Password = z.config.Properties["password"]
	ticket.Ticket.CustomerUser = z.config.Properties["customer"]

	requestBuffer := bytes.NewBufferString("")
	ticketEncoder := json.NewEncoder(requestBuffer)

	err := ticketEncoder.Encode(ticket)
	if err != nil {
		return TicketResponse{}, err
	}

	resp, err := z.httpClient.Post(z.config.URL+"/Ticket", "application/json", requestBuffer)

	if err != nil {
		return TicketResponse{}, err
	}

	if resp.StatusCode >= 400 {
		return TicketResponse{}, errors.New("Unhealthy Response Code (" + fmt.Sprint(resp.StatusCode) + ")")
	}
	var ticketResponse TicketResponse
	responseDecoder := json.NewDecoder(resp.Body)
	responseDecoder.Decode(&ticketResponse)

	if ticketResponse.Error.ErrorCode != "" {
		return ticketResponse, errors.New(ticketResponse.Error.ErrorMessage)
	}

	return ticketResponse, nil
}
