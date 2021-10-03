package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	clientReq = &http.Client{Timeout: 10 * time.Second}
)

type ExportResponse struct {
	Name     string          `json:"name"`
	Response ExportAgentData `json:"response"`
	Done     bool            `json:"done"`
}

type ExportAgentData struct {
	AgentContent string `json:"agentContent"`
}

func DFAgentExport(URL, token string) (res ExportResponse, err error) {
	req, err := http.NewRequest("POST", URL, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := clientReq.Do(req)
	if err != nil {
		err = fmt.Errorf("error request %v", err)
		return
	}
	defer resp.Body.Close()
	var bodyBytes []byte
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("error read body %v", err)
		return
	}

	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		err = fmt.Errorf("error unmarshal response (%s) got : %v", string(bodyBytes), err)
	}

	return
}
