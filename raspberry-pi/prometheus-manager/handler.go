package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os/exec"
)

const (
	JOBNAME_ROUTERS            = "routers"      // 모니터링 라우터 관리용 Job Name
	PATH_PIPE                  = "/pipe/docker" // 호스트에 명령어를 전송할 파이프
	COMMAND_RESTART_PROMETHEUS = "docker compose restart prometheus"
)

type APIHandler interface {
	AddTarget(rw http.ResponseWriter, req *http.Request)
	RemoveTarget(rw http.ResponseWriter, req *http.Request)
	ListTargets(rw http.ResponseWriter, req *http.Request)
}

type apiHandler struct {
	config Config
}

func NewAPIHander(config Config) APIHandler {
	return &apiHandler{
		config: config,
	}
}

func (ah *apiHandler) AddTarget(rw http.ResponseWriter, req *http.Request) {
	slog.Info("API Received: Add Target.")

	var body AddTargetRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		msg := fmt.Sprintf("error while decoding json: %s", err.Error())
		slog.Error(msg)
		http.Error(rw, msg, http.StatusBadRequest)
		return
	}

	err = ah.config.AddTarget(JOBNAME_ROUTERS, body.IP)
	if err != nil {
		msg := fmt.Sprintf("error while adding target to config: %s", err.Error())
		slog.Error(msg)
		http.Error(rw, msg, http.StatusInternalServerError)
		return
	}

	err = ah.config.Save()
	if err != nil {
		msg := fmt.Sprintf("error while saving config: %s", err.Error())
		slog.Error(msg)
		http.Error(rw, msg, http.StatusInternalServerError)
		return
	}

	rawCmd := fmt.Sprintf("echo \"%s\" > %s", COMMAND_RESTART_PROMETHEUS, PATH_PIPE)
	cmd := exec.Command("ash", "-c", rawCmd)
	_, err = cmd.Output()
	if err != nil {
		msg := fmt.Sprintf("error while executing command: %s", err.Error())
		slog.Error(msg)
		http.Error(rw, msg, http.StatusInternalServerError)
		return
	}

	slog.Info("successfully added")
	response := AddTargetResponse{Msg: "Successfully Added."}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}

func (ah *apiHandler) RemoveTarget(rw http.ResponseWriter, req *http.Request) {
	slog.Info("API Received: Remove Target.")

	var body RemoveTargetRequest
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		msg := fmt.Sprintf("error while decoding request JSON body: %s", err.Error())
		slog.Error(msg)
		http.Error(rw, msg, http.StatusBadRequest)
		return
	}

	err = ah.config.RemoveTarget(JOBNAME_ROUTERS, body.IP)
	if err != nil {
		msg := fmt.Sprintf("error while removing target from config: %s", err.Error())
		slog.Error(msg)
		http.Error(rw, msg, http.StatusInternalServerError)
		return
	}

	err = ah.config.Save()
	if err != nil {
		msg := fmt.Sprintf("error while saving config: %s", err.Error())
		slog.Error(msg)
		http.Error(rw, msg, http.StatusInternalServerError)
		return
	}

	rawCmd := fmt.Sprintf("echo \"%s\" > %s", COMMAND_RESTART_PROMETHEUS, PATH_PIPE)
	cmd := exec.Command("ash", "-c", rawCmd)
	_, err = cmd.Output()
	if err != nil {
		msg := fmt.Sprintf("error while executing command: %s", err.Error())
		slog.Error(msg)
		http.Error(rw, msg, http.StatusInternalServerError)
		return
	}

	slog.Info("successfully removed the target")
	response := AddTargetResponse{Msg: "Successfully Removed."}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}

func (ah *apiHandler) ListTargets(rw http.ResponseWriter, req *http.Request) {
	slog.Info("API Received: List Targets.")

	targets, err := ah.config.ListTargets(JOBNAME_ROUTERS)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	response := ListTargetsResponse{
		Msg:     "Successfully fetched the targets.",
		Targets: targets,
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}
