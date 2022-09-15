package savedata

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"fyneCode.go/entity"
)

func User_insert_json(username string, email string, contactNumber string) string {
	var Massage []string
	validationErr := false

	if !strings.Contains(email, "@") || email[len(email)-4:] != ".com" {
		validationErr = true
		emailError := "Email Address Not Valid !"
		Massage = append(Massage, emailError)
	}
	if len(contactNumber) != 10 {
		validationErr = true
		contactError := "Contact number not valid!"
		Massage = append(Massage, contactError)
	}
	if validationErr == true {
		return strings.TrimRight(strings.Join(Massage, " "+","), ",")

	}

	OldUser := []entity.User{}
	NewUser := entity.User{}
	input, err := ioutil.ReadFile("userData.json")
	json.Unmarshal(input, &OldUser)
	if err != nil {
		log.Println("error")
	}
	NewUser.UserName = username
	NewUser.PhoneNumber = contactNumber
	NewUser.Email = email
	OldUser = append(OldUser, NewUser)

	userData, err := json.Marshal(OldUser)

	ioutil.WriteFile("userData.json", userData, 0644)

	if err != nil {
		log.Println(err)
		Massage = append(Massage, "Error saving data!")
		return strings.TrimRight(strings.Join(Massage, ","), ",")
	}

	Massage = append(Massage, "All data saved Successfully.")
	return strings.TrimRight(strings.Join(Massage, ","), ",")
}
