package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
	"fmt"
)

type SmsArgs struct {
	Number, Content string
}

type EmailArgs struct {
	To, Subject, Content string
}

type Response struct {
	Result string
}

type SmsService struct{}
type EmailService struct{}

func (t *SmsService) SendSMS(r *http.Request, args *SmsArgs, result *Response) error {
	*result = Response{Result: fmt.Sprintf("Sms sent to %s", args.Number)}
	return nil
}

func (t *EmailService) SendEmail(r *http.Request, args *EmailArgs, result *Response) error {
	*result = Response{Result: fmt.Sprintf("Email sent to %s", args.To)}
	return nil
}


func main() {
	s := rpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	sms := new(SmsService)
	email := new(EmailService)

	s.RegisterService(sms, "sms")
	s.RegisterService(email, "email")

	r := mux.NewRouter()
	r.Handle("/delivery", s)
	http.ListenAndServe(":1337", r)
}