package model

type Message struct{
	Command     string    `json:"cmd"`
	Source    	string    `json:"src"`
	Destination	string    `json:"dst"`
	Message     string    `json:"msg"`
}