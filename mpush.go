package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/timehop/apns"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Message struct {
	Status int
	Badge  int
	Body   string
	Sound  string
}

func main() {
	r := gin.Default()

	r.GET("/help", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      http.StatusOK,
			"page":        "help",
			"description": "https://github.com/yuchan/mpush",
		})
	})

	r.POST("/push", func(c *gin.Context) {
		token := c.PostForm("token")
		body := c.PostForm("body")
		badge := c.DefaultPostForm("badge", "0")
		sound := c.DefaultPostForm("sound", "default.aiff")
		var msg Message
		msg.Status = http.StatusOK
		msg.Badge, _ = strconv.Atoi(badge)
		msg.Body = body
		msg.Sound = sound
		tokens := strings.Split(token, ",")
		for i, tok := range tokens {
			go func(i int, tok string) {
				SendPush(tok, "cert.pem", "key.pem", &msg)
			}(i, tok)
		}
		c.JSON(http.StatusOK, msg)
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

// Send PushNotification to device associated with token.
func SendPush(token string, cert string, key string, msg *Message) {
	c, err := apns.NewClientWithFiles(apns.SandboxGateway, cert, key)

	if err != nil {
		log.Fatal("could not create new client", err.Error())
	}

	go func() {
		for f := range c.FailedNotifs {
			fmt.Println("Notif", f.Notif.ID, "failed with", f.Err.Error())
		}
	}()

	p := apns.NewPayload()
	p.APS.Alert.Body = msg.Body
	badge := msg.Badge
	p.APS.Badge = &badge
	p.APS.Sound = msg.Sound
	p.APS.ContentAvailable = 1

	p.SetCustomValue("link", "zombo://dot/com")
	p.SetCustomValue("game", map[string]int{"score": 234})

	m := apns.NewNotification()
	m.Payload = p
	m.DeviceToken = token
	m.Priority = apns.PriorityImmediate
	m.Identifier = 12345
	m.ID = "user_id:timestamp"

	c.Send(m)
}
