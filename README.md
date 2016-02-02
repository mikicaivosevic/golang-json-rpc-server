##Golang JSON-RPC Server example

####SMS Args
```go
type SmsArgs struct {
	Number, Content string
}
```

####Email Args
```go
type EmailArgs struct {
	To, Subject, Content string
}
```

####Response structure for email and SMS service
```go
type Response struct {
	Result string
}
```

####Simple example of SMS service without logic
```go
type SmsService struct{}

func (t *SmsService) SendSMS(r *http.Request, args *SmsArgs, result *Response) error {
	*result = Response{Result: fmt.Sprintf("Sms sent to %s", args.Number)}
	return nil
}
```

####Simple example of Email service without logic
```go
type EmailService struct{}

func (t *EmailService) SendEmail(r *http.Request, args *EmailArgs, result *Response) error {
	*result = Response{Result: fmt.Sprintf("Email sent to %s", args.To)}
	return nil
}
```


####Example how to register services and start rpc server
```go
func main() {
	rpcServer := rpc.NewServer()

	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	sms := new(SmsService)
	email := new(EmailService)

	rpcServer.RegisterService(sms, "sms")
	rpcServer.RegisterService(email, "email")

	router := mux.NewRouter()
	router.Handle("/delivery", rpc)
	http.ListenAndServe(":1337", router)
}
```


####Example how to use JSON-RPC server written in Go with JSON-RPC client written in python

```python
import json
import requests

def rpc_call(url, method, args):
    headers = {'content-type': 'application/json'}
    payload = {
        "method": method,
        "params": [args],
        "jsonrpc": "2.0",
        "id": 1,
    }
    response = requests.post(url, data=json.dumps(payload), headers=headers).json()
    return response['result']

url = 'http://localhost:1337/delivery'

emailArgs = {'To': 'demo@example.com','Subject': 'Hello', 'Content': 'Hi!!!'}
smsArgs = {'Number': '381641234567', 'Content': 'Sms!!!'}
print rpc_call(url, 'email.SendEmail', emailArgs)
print rpc_call(url, 'sms.SendSMS', smsArgs)

```