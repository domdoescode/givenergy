package givenergy

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type GetInverterInfoResponse struct {
	DeviceType      string  `json:"deviceType"`
	ProductModel    string  `json:"productModel"`
	SerialNum       string  `json:"serialNum"`
	Address         int     `json:"address"`
	BatPercent      string  `json:"batPercent"`
	EndUserPostcode string  `json:"endUserPostcode"`
	NominalCapacity string  `json:"nominalCapacity"`
	NominalPower    float64 `json:"nominalPower"`
	EndUser         string  `json:"endUser"`
	Success         bool    `json:"success"`
	CompanyAddress  string  `json:"companyAddress"`
	Alias           string  `json:"alias"`
	DatalogSn       string  `json:"datalogSn"`
	CommissionDate  string  `json:"commissionDate"`
	Company         string  `json:"company"`
	EndUserAddress  string  `json:"endUserAddress"`
	BatteryType     string  `json:"batteryType"`
	CompanyPostcode string  `json:"companyPostcode"`
}

func GetInverterInfo(serialNumber string) (*GetInverterInfoResponse, error) {
	var response GetInverterInfoResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("inverter/getInverterInfo"), nil)
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

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type GetInverterRuntimeResponse struct {
	PGridText  string  `json:"pGridText"`
	PacText    string  `json:"pacText"`
	Ppv2       float64 `json:"ppv2"`
	Ppv2Text   string  `json:"ppv2Text"`
	Ppv1       float64 `json:"ppv1"`
	PLoadText  string  `json:"pLoadText"`
	Pac        float64 `json:"pac"`
	Success    bool    `json:"success"`
	Lost       bool    `json:"lost"`
	StatusText string  `json:"statusText"`
	PLoad      int     `json:"pLoad"`
	PGrid      float64 `json:"pGrid"`
	Ppv1Text   string  `json:"ppv1Text"`
	Status     string  `json:"status"`
}

func GetInverterRuntime(serialNumber string) (*GetInverterRuntimeResponse, error) {
	var response GetInverterRuntimeResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("inverter/getInverterRuntime"), nil)
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

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type GetInverterEnergyResponse struct {
	SerialNumber string `json:"serialNum"`
	EnergyInfo
}

func GetInverterEnergy(serialNumber string, outputType int, dateText string) (*GetInverterEnergyResponse, error) {
	var response GetInverterEnergyResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("inverter/getInverterEnergy"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("serialNum", serialNumber)
	q.Add("type", strconv.Itoa(outputType))
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

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type GetInverterBatCellResponse struct {
	Success       bool `json:"success"`
	Lost          bool `json:"lost"`
	MaxVoltModule int  `json:"maxVoltModule"`
	MinVoltModule int  `json:"minVoltModule"`
	Modules       []struct {
		CellVoltage1         int    `json:"cellVoltage1"`
		CellVoltage2         int    `json:"cellVoltage2"`
		CellVoltage3         int    `json:"cellVoltage3"`
		CellVoltage4         int    `json:"cellVoltage4"`
		CellVoltage5         int    `json:"cellVoltage5"`
		CellVoltage6         int    `json:"cellVoltage6"`
		CellVoltage7         int    `json:"cellVoltage7"`
		CellVoltage8         int    `json:"cellVoltage8"`
		CellVoltage9         int    `json:"cellVoltage9"`
		CellVoltage10        int    `json:"cellVoltage10"`
		CellVoltage11        int    `json:"cellVoltage11"`
		CellVoltage12        int    `json:"cellVoltage12"`
		CellVoltage13        int    `json:"cellVoltage13"`
		CellVoltage14        int    `json:"cellVoltage14"`
		CellVoltage15        int    `json:"cellVoltage15"`
		CellVoltage16        int    `json:"cellVoltage16"`
		CellTempreture1Text  string `json:"cellTempreture1Text"`
		CellTempreture2Text  string `json:"cellTempreture2Text"`
		CellTempreture3Text  string `json:"cellTempreture3Text"`
		CellTempreture4Text  string `json:"cellTempreture4Text"`
		ModuleVoltage        int    `json:"moduleVoltage"`
		DesignCapText        string `json:"designCapText"`
		ModuleSoc            int    `json:"moduleSoc"`
		HasBmsCellModuleInfo bool   `json:"hasBmsCellModuleInfo"`
		ModuleTempretureText string `json:"moduleTempretureText"`
		MaxVoltDiffValue     int    `json:"maxVoltDiffValue"`
		Module               int    `json:"module"`
		Charging             int    `json:"charging"`
		Discharging          int    `json:"discharging"`
		MaxVoltItem          int    `json:"maxVoltItem"`
		MinVoltItem          int    `json:"minVoltItem"`
		FullCapText          string `json:"fullCapText"`
	} `json:"modules"`
}

func GetInverterBatCell(serialNumber string) (*GetInverterBatCellResponse, error) {
	var response GetInverterBatCellResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("pcs/batCell/getBatCellData"), nil)
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

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}
