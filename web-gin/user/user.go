package user

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	User
	Msg string `json:"message"`
}

func (u *User) say(msg string) *Message {
	ret := Message{
		*u, u.Name + " said, " + msg,
	}
	return &ret
}

func UserSay(c *gin.Context) {
	action := c.Query("action")
	if action != "say" {
		c.JSON(http.StatusExpectationFailed, "Not Support")
		return
	}

	name := c.Query("name")
	user := User{
		rand.Int63(), name,
	}
	msg := c.Query("msg")
	message := user.say(msg)

	c.JSON(http.StatusOK, message)
}
