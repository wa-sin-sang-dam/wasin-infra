package main

type AddTargetRequest struct {
	JobName string `json:"job_name"`
	IP      string `json:"ip"`
}

type ListTargetRequest struct{}

type RemoveTargetRequest struct {
	JobName string `json:"job_name"`
	IP      string `json:"ip"`
}
