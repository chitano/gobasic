package main

import (
	"fmt"
	"regexp"
	"time"
)

type userID uint32
type Email string
type AddressId uint32
type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Address struct {
	Id           AddressId
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	PostalCode   string
	Country      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type User struct {
	Id      userID
	Name    string
	Email   Email
	Address Address
	Phone   string
	Gender  Gender
}

func (g Gender) IsValid() bool {
	switch g {
	case Male, Female:
		return true
	}
	return false
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func (e Email) IsValid() bool {
	return emailRegex.MatchString(string(e))
}

func main() {
	now := time.Now()

	users := []User{
		{
			Id:    1,
			Name:  "Alice Johnson",
			Email: "alice.johnsonexample.com",
			Address: Address{
				Id:           101,
				AddressLine1: "123 Maple Street",
				AddressLine2: "Apt 4B",
				City:         "Seattle",
				State:        "WA",
				PostalCode:   "98101",
				Country:      "USA",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			Phone:  "+1-206-555-0101",
			Gender: Female,
		},
		{
			Id:    2,
			Name:  "Bob Smith",
			Email: "bob.smith@example.com",
			Address: Address{
				Id:           102,
				AddressLine1: "456 Oak Avenue",
				AddressLine2: "Suite 12",
				City:         "Austin",
				State:        "TX",
				PostalCode:   "73301",
				Country:      "USA",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			Phone:  "+1-512-555-0102",
			Gender: Male,
		},
		{
			Id:    3,
			Name:  "Carol Lee",
			Email: "carol.lee@example.com",
			Address: Address{
				Id:           103,
				AddressLine1: "789 Pine Road",
				AddressLine2: "",
				City:         "Denver",
				State:        "CO",
				PostalCode:   "80014",
				Country:      "USA",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			Phone:  "+1-303-555-0103",
			Gender: Female,
		},
		{
			Id:    4,
			Name:  "David Brown",
			Email: "david.brown@example.com",
			Address: Address{
				Id:           104,
				AddressLine1: "321 Cedar Lane",
				AddressLine2: "Floor 2",
				City:         "Chicago",
				State:        "IL",
				PostalCode:   "60601",
				Country:      "USA",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			Phone:  "+1-312-555-0104",
			Gender: Male,
		},
		{
			Id:    5,
			Name:  "Eva Green",
			Email: "eva.green@example.com",
			Address: Address{
				Id:           105,
				AddressLine1: "654 Birch Boulevard",
				AddressLine2: "Unit 9",
				City:         "Miami",
				State:        "FL",
				PostalCode:   "33101",
				Country:      "USA",
				CreatedAt:    now,
				UpdatedAt:    now,
			},
			Phone:  "+1-305-555-0105",
			Gender: Female,
		},
	}

	for _, u := range users {
		fmt.Println(u.Email.IsValid())
		fmt.Printf("User: %-15s | Email: %-30s | Valid: %v\n", u.Name, u.Email, u.Email.IsValid())
	}

}
