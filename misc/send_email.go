package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	t, err := template.ParseFiles("base.html")
	if err != nil {
		log.Println(err)
	}

	i := struct {
		Name string
	}{"TestName"}

	var buf bytes.Buffer
	if err := t.Execute(&buf, i); err != nil {
		log.Println(err)
	}

	body := buf.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "test_smtp@corp.acronis.com")
	m.SetHeader("To", []string{"magomed.nakhaev@acronis.de", "magomed.nakhaev@acronis.ru"}...)
	m.SetHeader("Subject", "Hello")
	m.SetHeader("X-Tenant-UUID", "a713ab12-11e3-4ceb-82ab-f2d0a6314345")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.corp.acronis.com", 25, "", "")
	//d := gomail.NewDialer("10.248.50.56", 25, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("Message sent successfully")
	}
}
