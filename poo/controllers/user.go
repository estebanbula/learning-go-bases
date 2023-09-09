package controllers

import (
	"fmt"
	"learning-go/poo/models"
	"time"
)

func NewUser() {
	user := new(models.User)
	user.AddUser(1, "Esteban", time.Now(), true)
	fmt.Println(user)
}
