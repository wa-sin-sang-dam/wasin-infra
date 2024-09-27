package main

type AddTargetResponse struct {
	Msg string `json:"msg"`
}

type RemoveTargetResponse struct {
	Msg string `json:"msg"`
}

type ListTargetsResponse struct {
	Msg     string   `json:"msg"`
	Targets []string `json:"targets"`
}
