package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"ms-order/model"
	"net/http"
)

func TrackPackage(awb string, courier string) (model.TrackingInfo, error) {
	url := fmt.Sprintf("https://api.binderbyte.com/v1/track?api_key=%s&courier=%s&awb=%s",
		"2979fc0c9a5992172dade3a5176868b3d24cec2db60b1f52ae535b6c7eab9618", courier, awb)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return model.TrackingInfo{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	// send HTTP req w/default client
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.TrackingInfo{}, err
	}

	defer res.Body.Close()

	// read body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return model.TrackingInfo{}, err
	}

	var respBody model.TrackingInfo

	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return model.TrackingInfo{}, err
	}

	return respBody, nil
}
