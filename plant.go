package givenergy

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type GetPlantListResponse struct {
	Total   int         `json:"total"`
	Success bool        `json:"success"`
	Rows    []PlantInfo `json:"rows"`
}

func GetPlantList() (*GetPlantListResponse, error) {
	var response GetPlantListResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plant/getPlantList"), nil)
	if err != nil {
		return &response, err
	}

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

type GetPlantInfoResponse struct {
	PlantInfo
}

func GetPlantInfo(plantID int) (*GetPlantInfoResponse, error) {
	var response GetPlantInfoResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plant/getPlantInfo"), nil)
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

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type GetPlantSummaryResponse struct {
	ExportEnergyTodayText           string  `json:"exportEnergyTodayText"`
	ImportEnergyTodayText           string  `json:"importEnergyTodayText"`
	ImportEnergyToday               string  `json:"importEnergyToday"`
	Soc                             int     `json:"soc"`
	ExportEnergyTotal               string  `json:"exportEnergyTotal"`
	ImportEnergyTotal               string  `json:"importEnergyTotal"`
	ExportEnergyToday               string  `json:"exportEnergyToday"`
	TotalGenerationText             string  `json:"totalGenerationText"`
	BatPower                        string  `json:"batPower"`
	HasMeterInfo                    bool    `json:"hasMeterInfo"`
	CurrentText                     string  `json:"currentText"`
	ConsumptionEnergyTotalText      string  `json:"consumptionEnergyTotalText"`
	BatteryCapacity                 float64 `json:"batteryCapacity"`
	LoadPowerkW                     string  `json:"loadPowerkW"`
	ConsumptionEnergyToday          string  `json:"consumptionEnergyToday"`
	GridPower                       string  `json:"gridPower"`
	EnergyTodayText                 string  `json:"energyTodayText"`
	ActiveInverterCount             string  `json:"activeInverterCount"`
	AverageHomeSelfConsumptionToday int     `json:"averageHomeSelfConsumptionToday"`
	ExportEnergyTotalText           string  `json:"exportEnergyTotalText"`
	ImportEnergyTotalText           string  `json:"importEnergyTotalText"`
	PvPower                         string  `json:"pvPower"`
	GridPowerkW                     string  `json:"gridPowerkW"`
	Success                         bool    `json:"success"`
	TotalGeneration                 string  `json:"totalGeneration"`
	AverageHomeSelfConsumptionTotal string  `json:"averageHomeSelfConsumptionTotal"`
	BttTotalText                    string  `json:"bttTotalText"`
	LoadPower                       int     `json:"loadPower"`
	DischargeEnergyToday            string  `json:"dischargeEnergyToday"`
	TotalCo2Saved                   string  `json:"totalCo2Saved"`
	BtMonth                         string  `json:"btMonth"`
	VoltageText                     string  `json:"voltageText"`
	BttTotal                        string  `json:"bttTotal"`
	SelfConsumptionTotalText        string  `json:"selfConsumptionTotalText"`
	ChargeEnergyToday               string  `json:"chargeEnergyToday"`
	SelfConsumptionTodayText        string  `json:"selfConsumptionTodayText"`
	PInv                            int     `json:"pInv"`
	PvPowerText                     string  `json:"pvPowerText"`
	SelfConsumptionToday            string  `json:"selfConsumptionToday"`
	SavingsToday                    string  `json:"savingsToday"`
	SavingsTotal                    string  `json:"savingsTotal"`
	SelfConsumptionTotal            string  `json:"selfConsumptionTotal"`
	ConsumptionEnergyTotal          string  `json:"consumptionEnergyTotal"`
	GridRelianceToday               int     `json:"gridRelianceToday"`
	GridEnergyOutToday              string  `json:"gridEnergyOutToday"`
	GridRelianceTotal               int     `json:"gridRelianceTotal"`
	ConsumptionEnergyTodayText      string  `json:"consumptionEnergyTodayText"`
	EnergyToday                     string  `json:"energyToday"`
	GridEnergyOutTotal              string  `json:"gridEnergyOutTotal"`
	BtToday                         string  `json:"btToday"`
	FrequencyText                   string  `json:"frequencyText"`
	FeedInTariffToday               string  `json:"feedInTariffToday"`
	FeedInTariffTotal               string  `json:"feedInTariffTotal"`
	BatteryCapacityText             string  `json:"batteryCapacityText"`
	LoadPowerText                   string  `json:"loadPowerText"`
	MonthCo2Saved                   string  `json:"monthCo2Saved"`
	TodayCo2Saved                   string  `json:"todayCo2Saved"`
	DailyAlarmCount                 int     `json:"dailyAlarmCount"`
	PvPowerkW                       string  `json:"pvPowerkW"`
}

func GetPlantSummary() (*GetPlantSummaryResponse, error) {
	var response GetPlantSummaryResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plant/getPlantSummary"), nil)
	if err != nil {
		return &response, err
	}

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

type GetPlantRuntimeResponse struct {
	NormalInverterCount  int    `json:"normalInverterCount"`
	EnergyTotal          string `json:"energyTotal"`
	EnergyTodayText      string `json:"energyTodayText"`
	EnergyToday          string `json:"energyToday"`
	EnergyTotalText      string `json:"energyTotalText"`
	PlantID              int    `json:"plantId"`
	UnknownInverterCount int    `json:"unknownInverterCount"`
	ErrorInverterCount   int    `json:"errorInverterCount"`
	LostDatalogCount     int    `json:"lostDatalogCount"`
	NominalPower         int    `json:"nominalPower"`
	NormalDatalogCount   int    `json:"normalDatalogCount"`
	CurrentPower         string `json:"currentPower"`
	Load                 string `json:"load"`
	EnergyTree           string `json:"energyTree"`
	Success              bool   `json:"success"`
	WaitingInverterCount int    `json:"waitingInverterCount"`
	CurrentPowerText     string `json:"currentPowerText"`
	LostInverterCount    int    `json:"lostInverterCount"`
	EnergyCo2            string `json:"energyCo2"`
}

func GetPlantRuntime(plantID int) (*GetPlantRuntimeResponse, error) {
	var response GetPlantRuntimeResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plant/getPlantRuntime"), nil)
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

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type GetPlantDevicesResponse struct {
	Datalogs []struct {
		DeviceType  string `json:"deviceType"`
		SerialNum   string `json:"serialNum"`
		HasInverter bool   `json:"hasInverter"`
		Lost        bool   `json:"lost"`
		DatalogType string `json:"datalogType"`
		Inverters   []struct {
			DeviceType       string `json:"deviceType"`
			ProductModel     string `json:"productModel"`
			SerialNum        string `json:"serialNum"`
			Address          int    `json:"address"`
			BatPercent       string `json:"batPercent"`
			EnergyTodayText  string `json:"energyTodayText"`
			EnergyTotalText  string `json:"energyTotalText"`
			CurrentPowerText string `json:"currentPowerText"`
			NominalCapacity  int    `json:"nominalCapacity"`
			Alias            string `json:"alias"`
			BatteryType      string `json:"batteryType"`
			Status           string `json:"status"`
		} `json:"inverters"`
		Alias string `json:"alias"`
	} `json:"datalogs"`
	Success bool `json:"success"`
	PlantID int  `json:"plantId"`
}

func GetPlantDevices(plantID int) (*GetPlantDevicesResponse, error) {
	var response GetPlantDevicesResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plant/getPlantDevices"), nil)
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

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type GetPlantEnergyResponse struct {
	PlantID int `json:"plantId"`
	EnergyInfo
}

func GetPlantEnergy(plantID, outputType int, dateText string) (*GetPlantEnergyResponse, error) {
	var response GetPlantEnergyResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plant/getPlantEnergy"), nil)
	if err != nil {
		return &response, err
	}

	q := req.URL.Query()
	q.Add("plantId", strconv.Itoa(plantID))
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

type GetPlantEventsResponse struct {
	Total   int    `json:"total"`
	Success bool   `json:"success"`
	PlantID string `json:"plantId"`
	Rows    []struct {
		CustomEventTypeText string `json:"customEventTypeText"`
		SerialNum           string `json:"serialNum"`
		RecordID            int    `json:"recordId"`
		CustomEventType     string `json:"customEventType"`
		EndUser             string `json:"endUser"`
		EventText           string `json:"eventText"`
		Renormal            bool   `json:"renormal"`
		Alias               string `json:"alias"`
		DatalogSn           string `json:"datalogSn"`
		Company             string `json:"company"`
		StartTime           string `json:"startTime"`
		RenormalTime        string `json:"renormalTime"`
		EndTime             string `json:"endTime"`
		Event               string `json:"event"`
		ClearedTime         string `json:"clearedTime"`
		Status              string `json:"status"`
	} `json:"rows"`
}

func GetPlantEvents(plantID int) (*GetPlantEventsResponse, error) {
	var response GetPlantEventsResponse

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", getURL("plant/event/getPlantEvents"), nil)
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

	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}

type PlantInfo struct {
	Country          string `json:"country"`
	EnergyTotal      string `json:"energyTotal"`
	EnergyTodayText  string `json:"energyTodayText"`
	Timezone         int    `json:"timezone"`
	TimezoneText     string `json:"timezoneText"`
	EnergyToday      string `json:"energyToday"`
	EnergyTotalText  string `json:"energyTotalText"`
	PlantID          int    `json:"plantId"`
	NominalPower     int    `json:"nominalPower"`
	CurrentPower     string `json:"currentPower"`
	Name             string `json:"name"`
	CurrentPowerText string `json:"currentPowerText"`
	CountryText      string `json:"countryText"`
	CreateDate       string `json:"createDate"`
}

type EnergyInfo struct {
	Success    bool   `json:"success"`
	EnergyText string `json:"energyText"`
	Energy     string `json:"energy"`
}
