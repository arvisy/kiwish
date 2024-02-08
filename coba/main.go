package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	RajaOngkir struct {
		Results []struct {
			CityID   string `json:"city_id"`
			CityName string `json:"city_name"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func main() {
	// data, err := os.ReadFile("./data.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var response Response
	// err = json.Unmarshal(data, &response)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// container := map[string]string{}
	// for _, val := range response.RajaOngkir.Results {
	// 	container[val.CityName] = val.CityID
	// }

	// data, err = json.Marshal(container)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// os.WriteFile("./result.json", data, 0777)

	url := "https://api.rajaongkir.com/starter/cost"

	payload := strings.NewReader("origin=501&destination=114&weight=1700&courier=jne")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("key", "fc45837355230c7ae97b1f68656bcd66")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
