// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewVitamin struct {
	VitaminType string `json:"VitaminType"`
	Benefits    string `json:"Benefits"`
}

type UpdatedVitamin struct {
	VitaminID   string `json:"VitaminId"`
	VitaminType string `json:"VitaminType"`
	Benefits    string `json:"Benefits"`
}

type Vitamin struct {
	VitaminID   string `json:"VitaminId"`
	VitaminType string `json:"VitaminType"`
	Benefits    string `json:"Benefits"`
}
