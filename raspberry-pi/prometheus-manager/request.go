package main

type Empty struct{}

type AddTargetRequest struct {
	IP string `json:"ip"`
}

type ListTargetRequest Empty

type RemoveTargetRequest struct {
	IP string `json:"ip"`
}
