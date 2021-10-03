package df

import (
	"fmt"
	"io/ioutil"

	b64 "encoding/base64"

	"github.com/nbanitama/df-supoort/api"
	"github.com/nbanitama/df-supoort/auth"
)

func Export(projectID, serviceAccountPath, destFile string) (err error) {
	fmt.Printf("Starting exporting the agent %v\n", projectID)
	baseURL := "https://dialogflow.googleapis.com/v2/"
	URL := baseURL + "projects/" + projectID + "/agent:export"

	token := auth.GetAccessTokenFromServiceAccount(serviceAccountPath)

	result, err := api.DFAgentExport(URL, token)
	if err != nil {
		fmt.Printf("get err %v\n", err.Error())
		return
	}
	fmt.Println(result.Name, result.Done)

	uDec, err := b64.StdEncoding.DecodeString(result.Response.AgentContent)
	if err != nil {
		fmt.Printf("error decoding base64 %v\n", err)
		return
	}
	err = ioutil.WriteFile(destFile, uDec, 0)
	if err != nil {
		fmt.Printf("error writing file %v\n", err)
	} else {
		fmt.Println("DONE")
	}
	return
}
