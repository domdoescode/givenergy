package givenergy

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type GetPlantDayLineResponse struct {
	Data []struct {
		Pac    string `json:"pac"`
		Month  int    `json:"month"`
		Hour   int    `json:"hour"`
		Year   int    `json:"year"`
		Time   string `json:"time"`
		Day    int    `json:"day"`
		Minute int    `json:"minute"`
		Second int    `json:"second"`
	} `json:"data"`
	Success bool `json:"success"`
}

func GetPlantDayLine(plantID int, dateText string) (*GetPlantDayLineResponse, error) {
	var response GetPlantDayLineResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plantChart/dayLine"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("plantId", strconv.Itoa(plantID))
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

type GetPlantMonthColumnResponse struct {
	Data []struct {
		ExportEnergy         string `json:"exportEnergy"`
		DischargeEnergyToday string `json:"DischargeEnergyToday"`
		PvEnergy             string `json:"pvEnergy"`
		ChargeEnergyToday    string `json:"ChargeEnergyToday"`
		ImportEnergy         string `json:"importEnergy"`
		ConsumptionEnergy    string `json:"consumptionEnergy"`
		InvImportEnergy      string `json:"InvImportEnergy"`
		InvExportEnergy      string `json:"InvExportEnergy"`
		Day                  int    `json:"day"`
	} `json:"data"`
	Success bool `json:"success"`
	DayMax  int  `json:"dayMax"`
}

func GetPlantMonthColumn(plantID int, year, month string) (*GetPlantMonthColumnResponse, error) {
	var response GetPlantMonthColumnResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plantChart/monthColumn"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("plantId", strconv.Itoa(plantID))
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

type GetPlantYearColumnResponse struct {
	Data []struct {
		ExportEnergy         string `json:"exportEnergy"`
		DischargeEnergyToday string `json:"DischargeEnergyToday"`
		PvEnergy             string `json:"pvEnergy"`
		ChargeEnergyToday    string `json:"ChargeEnergyToday"`
		ImportEnergy         string `json:"importEnergy"`
		ConsumptionEnergy    string `json:"consumptionEnergy"`
		InvImportEnergy      string `json:"InvImportEnergy"`
		Month                int    `json:"month"`
		InvExportEnergy      string `json:"InvExportEnergy"`
	} `json:"data"`
	Success bool `json:"success"`
}

func GetPlantYearColumn(plantID int, year string) (*GetPlantYearColumnResponse, error) {
	var response GetPlantYearColumnResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plantChart/yearColumn"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("plantId", strconv.Itoa(plantID))
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

type GetPlantTotalColumnResponse struct {
	Data []struct {
		ExportEnergy         string `json:"exportEnergy"`
		DischargeEnergyToday string `json:"DischargeEnergyToday"`
		PvEnergy             string `json:"pvEnergy"`
		ChargeEnergyToday    string `json:"ChargeEnergyToday"`
		ImportEnergy         string `json:"importEnergy"`
		ConsumptionEnergy    string `json:"consumptionEnergy"`
		InvImportEnergy      string `json:"InvImportEnergy"`
		Year                 int    `json:"year"`
		InvExportEnergy      string `json:"InvExportEnergy"`
	} `json:"data"`
	Success bool `json:"success"`
}

func GetPlantTotalColumn(plantID int) (*GetPlantTotalColumnResponse, error) {
	var response GetPlantTotalColumnResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plantChart/totalColumn"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("plantId", strconv.Itoa(plantID))
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
