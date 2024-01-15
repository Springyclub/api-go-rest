package models

type Personality struct {
	Name    string `json:"name"`
	History string `json:"history"`
	Id      int    `json:"id"`
}

var Personalities []Personality
