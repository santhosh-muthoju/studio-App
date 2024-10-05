package models

import "time"

type Class struct {
	ID        string
	ClassName string
	StartDate time.Time
	EndDate   time.Time
	Capacity  int
}

type ClassRequest struct {
	ClassName string `json:"className"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Capacity  int    `json:"capacity"`
}

var ClassesList []Class

func AddClass(class Class) {
	ClassesList = append(ClassesList, class)
}

func FindClassByID(id string) (Class, bool) {
	for _, class := range ClassesList {
		if class.ID == id {
			return class, true
		}
	}
	return Class{}, false
}
