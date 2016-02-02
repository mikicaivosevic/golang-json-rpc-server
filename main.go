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
	rpc := rpc.NewServer()

	rpc.RegisterCodec(json.NewCodec(), "application/json")
	rpc.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	sms := new(SmsService)
	email := new(EmailService)

	rpc.RegisterService(sms, "sms")
	rpc.RegisterService(email, "email")

	router := mux.NewRouter()
	router.Handle("/delivery", rpc)
	http.ListenAndServe(":1337", router)
}