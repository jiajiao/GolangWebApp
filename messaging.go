package main

import (
	"log"
	"net/smtp"
	"net/url"
	"fmt"
	"errors"
)

type loginAuth struct{
username, password string
}

func emaillnk(e, u, h string) {
	// Set up authentication information.i
	fmt.Println(e)
	fmt.Println(u)
	fmt.Println(h)

	//using our  outlook365.com register@gocloudcoding.com address requires SSL which requires a paid for service on Herokuto - to be added later.
	auth := LoginAuth("ricpelong@hotmail.co.uk", "2shoes2loose1")
	baseUrl := "http://golangwebapp.herokuapp.com/verify/?"
	params := url.Values{}
	params.Add("token", h)
	params.Add("email", e)
	finalUrl := baseUrl + params.Encode()

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{e}
	msg := []byte("To:" + e + "\r\n" +
		"Subject: Complete App Registration\r\n" +
		"\r\n" +
		"Click on link to complete App registration.\r\n" +
		finalUrl)
	err := smtp.SendMail("smtp.live.com:587", auth, "ricpelong@hotmail.co.uk", to, msg)

	if err != nil {
		log.Fatal(err)
	}
}

//from: github.com/andelf/5118732
//Outlook nolonger supports Auth Plain,the method in smtp. Andelf has already produced a Auth Login version, so it was easier touse that.

func LoginAuth(username,password string) smtp.Auth{
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server*smtp.ServerInfo)(string, []byte, error){
	return"LOGIN", []byte{},nil
}


func (a *loginAuth) Next(fromServer []byte, more bool)([]byte, error){
	if more{
		switch string(fromServer){
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown fromServer")
}


}
return nil, nil

}



