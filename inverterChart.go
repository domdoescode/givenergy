package givenergy

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type InverterAttribute string

var (
	PAC         InverterAttribute = "pac"
	PPV         InverterAttribute = "ppv"
	VPV1        InverterAttribute = "vpv1"
	IPV1        InverterAttribute = "ipv1"
	VPV2        InverterAttribute = "vpv2"
	IPV2        InverterAttribute = "ipv2"
	PACR        InverterAttribute = "pacr"
	PACS        InverterAttribute = "pacs"
	PACT        InverterAttribute = "pact"
	Temperature InverterAttribute = "temperature"
)

type GetInverterDayLineResponse struct {
	AvgValue     float64 `json:"avgValue"`
	MinValueText string  `json:"minValueText"`
	YAxis        string  `json:"yAxis"`
	XAxis        string  `json:"xAxis"`
	Data         []struct {
		Month  int     `json:"month"`
		Hour   int     `json:"hour"`
		Year   int     `json:"year"`
		Time   string  `json:"time"`
		Day    int     `json:"day"`
		Value  float64 `json:"value"`
		Minute int     `json:"minute"`
		Second int     `json:"second"`
	} `json:"data"`
	Success      bool   `json:"success"`
	AvgValueText string `json:"avgValueText"`
	MaxValueText string `json:"maxValueText"`
}

func GetInverterDayLine(serialNumber string, attribute InverterAttribute, dateText string) (*GetInverterDayLineResponse, error) {
	var response GetInverterDayLineResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("invChart/dayLine"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("serialNum", serialNumber)
	q.Add("attr", string(attribute))
	q.Add("dateText", dateText)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return &response, err
	}

	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &response, err
	}

	log.Println(string(respByte))

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type GetInverterMonthColumnResponse struct {
	Data []struct {
		ExportEnergy      float64 `json:"exportEnergy"`
		ImportEnergy      float64 `json:"importEnergy"`
		ConsumptionEnergy float64 `json:"consumptionEnergy"`
		Day               int     `json:"day"`
		Energy            float64 `json:"energy"`
	} `json:"data"`
	Success bool `json:"success"`
	DayMax  int  `json:"dayMax"`
}

func GetInverterMonthColumn(serialNumber string, year, month string) (*GetInverterMonthColumnResponse, error) {
	var response GetInverterMonthColumnResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("invChart/monthColumn"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("serialNum", serialNumber)
	q.Add("year", year)
	q.Add("month", month)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return &response, err
	}

	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &response, err
	}

	log.Println(string(respByte))

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type GetInverterYearColumnResponse struct {
	Data []struct {
		ExportEnergy      float64 `json:"exportEnergy"`
		ImportEnergy      float64 `json:"importEnergy"`
		ConsumptionEnergy float64 `json:"consumptionEnergy"`
		Month             int     `json:"month"`
		Energy            float64 `json:"energy"`
	} `json:"data"`
	Year    int  `json:"year"`
	Success bool `json:"success"`
}

func GetInverterYearColumn(serialNumber string, year string) (*GetInverterYearColumnResponse, error) {
	var response GetInverterYearColumnResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("invChart/yearColumn"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("serialNum", serialNumber)
	q.Add("year", year)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return &response, err
	}

	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &response, err
	}

	log.Println(string(respByte))

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type GetInverterTotalColumnResponse struct {
	Data []struct {
		ImportEnergy      float64 `json:"importEnergy"`
		ConsumptionEnergy float64 `json:"consumptionEnergy"`
		Year              int     `json:"year"`
		ExportEnergy      float64 `json:"ExportEnergy"`
		Energy            float64 `json:"energy"`
	} `json:"data"`
	Success bool `json:"success"`
}

func GetInverterTotalColumn(serialNumber string) (*GetInverterTotalColumnResponse, error) {
	var response GetInverterTotalColumnResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("invChart/totalColumn"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("serialNum", serialNumber)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return &response, err
	}

	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &response, err
	}

	log.Println(string(respByte))

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}
