package model

// Worker is a specialist
type Worker struct {
	ID       int    `json:"id" bson:"ID"`
	Name     string `json:"name" bson:"Name"`
	Surname  string `json:"surname" bson:"Surname"`
	Age      int    `json:"age" bson:"Age"`
	Image    string `json:"image" bson:"Image"`
	Sex      int    `json:"sex" bson:"Sex"`
	About    string `json:"about" bson:"About"`
	JobTitle string `json:"jobTitle" bson:"JobTitle"`
	Phone    string `json:"phone" bson:"Phone"`
	Email    string `json:"email" bson:"Email"`
	Level    int    `json:"level" bson:"Level"` //level is a skill
	Salary   int    `json:"salary" bson:"Salary"`

	// Generated vars
	Price int `json:"price" bson:"Price"` //price = level*priceDelta

	// Local vars
	WCID int `json:"-" bson:"WCID"` //working company id

	// Changable vars
	Jobs []*JobHistory `json:"jobs" bson:"Jobs"`
}

// JobHistory is one work
type JobHistory struct {
	StartYear int    `json:"startYear" bson:"StartYear"`
	EndYear   int    `json:"endYear" bson:"EndYear"`
	Company   string `json:"company" bson:"Company"`
	Location  string `json:"location" bson:"Location"`
	Comment   string `json:"comment" bson:"Comment"`
	Title     string `json:"title" bson:"Title"`
}

// HireRegist is request data while hireing
type HireRegist struct {
	ID int `json:"id"`
}
