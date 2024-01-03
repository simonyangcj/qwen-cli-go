package qwen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qwen-cli/cmd/qwen/app/config"
	"qwen-cli/pkg/api"
	"qwen-cli/pkg/util"
	"time"
)

func SendToQwen(config *config.ApiConfig, input *api.Input, parameters *api.ApiParameter, model string, enableSSE bool, timeoutSecond int) (*api.Response, error) {
	headers := handleHeader(enableSSE, config)
	request := &api.Request{
		Input:      input,
		Parameters: parameters,
		Model:      model,
	}

	dataSend, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	//fmt.Println(string(dataSend), headers)
	result, code, err := util.CommonRequest(config.ApiURL, "POST", "", dataSend, headers, true, true, time.Duration(timeoutSecond)*time.Second)
	if err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf("request failed, code: %d, result: %s", code, result)
	}

	response := &api.Response{}

	err = json.Unmarshal(result, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func handleHeader(enableSSE bool, config *config.ApiConfig) map[string]string {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", config.ApiKey),
	}
	if enableSSE {
		// in fact one of following header is enough
		// but i think it won't hurt to set both
		headers["Accept"] = "text/event-stream"
		headers["X-DashScope-SSE"] = "enable"
	}
	return headers
}
