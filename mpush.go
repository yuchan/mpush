package main

import (
	"encoding/json"
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
	Custom string
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
		customValue := c.DefaultPostForm("custom", "")
		var msg Message
		msg.Status = http.StatusOK
		msg.Badge, _ = strconv.Atoi(badge)
		msg.Body = body
		msg.Sound = sound
		msg.Custom = customValue
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

	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(msg.Custom), &dat); err != nil {
		fmt.Println(err)
	} else {
		for k, v := range dat {
			p.SetCustomValue(k, valueString(v))
		}
	}

	m := apns.NewNotification()
	m.Payload = p
	m.DeviceToken = token
	m.Priority = apns.PriorityImmediate
	m.Identifier = 12345
	m.ID = "user_id:timestamp"

	c.Send(m)
}

func valueString(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	case float64:
		return fmt.Sprint(data.(float64))
	case bool:
		return fmt.Sprint(data.(bool))
	case nil:
		return "null"
	case []interface{}:
		var str []byte
		str = append(str, '[')
		for _, v := range data.([]interface{}) {
			str = append(str, valueString(v)...)
			str = append(str, ' ')
		}
		str = append(str, ']')
		return fmt.Sprint(string(str))
	case map[string]interface{}:
		var str []byte
		str = append(str, '{')
		for k, v := range data.(map[string]interface{}) {
			str = append(str, k...)
			str = append(str, ':')
			str = append(str, valueString(v)...)
			str = append(str, ' ')
		}
		str = append(str, '}')
		return fmt.Sprint(string(str))
	default:
		return ""
	}
}
