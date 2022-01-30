package arso

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ARSO struct {
	Forecast24H struct {
		Features []struct {
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
			Type       string `json:"type"`
			Properties struct {
				ID       string `json:"id"`
				Country  string `json:"country"`
				ParentID string `json:"parentId"`
				Days     []struct {
					Date     string `json:"date"`
					Timeline []struct {
						Interval                      string    `json:"interval"`
						RhShortText                   string    `json:"rh_shortText"`
						CloudsShortText               string    `json:"clouds_shortText"`
						DdShortText                   string    `json:"dd_shortText"`
						Tnsyn                         string    `json:"tnsyn"`
						CloudsIconWwsynIcon           string    `json:"clouds_icon_wwsyn_icon"`
						Valid                         time.Time `json:"valid"`
						Rh                            string    `json:"rh"`
						DdffIcon                      string    `json:"ddff_icon"`
						CloudsShortTextWwsynShortText string    `json:"clouds_shortText_wwsyn_shortText"`
						PaShortText                   string    `json:"pa_shortText"`
						Txsyn                         string    `json:"txsyn"`
						Msl                           string    `json:"msl"`
						Tp24HAcc                      string    `json:"tp_24h_acc"`
						WwsynShortText                string    `json:"wwsyn_shortText"`
						WwsynDecodeText               string    `json:"wwsyn_decodeText"`
						FfVal                         string    `json:"ff_val"`
						FfShortText                   string    `json:"ff_shortText"`
						FfmaxVal                      string    `json:"ffmax_val"`
						WwsynIcon                     string    `json:"wwsyn_icon"`
						CloudBaseShortText            string    `json:"cloudBase_shortText"`
						Time                          string    `json:"time"`
					} `json:"timeline"`
					UTCoffset string    `json:"UTCoffset"`
					Sunset    time.Time `json:"sunset"`
					Sunrise   time.Time `json:"sunrise"`
				} `json:"days"`
				Title string `json:"title"`
			} `json:"properties"`
		} `json:"features"`
		Language    string    `json:"language"`
		IconFormat  string    `json:"icon_format"`
		IconBaseURL string    `json:"icon_base_url"`
		DataType    string    `json:"dataType"`
		TsUpdated   time.Time `json:"tsUpdated"`
		TsValid     string    `json:"tsValid"`
		Type        string    `json:"type"`
	} `json:"forecast24h"`

	Observation struct {
		Features []struct {
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
			Type       string `json:"type"`
			Properties struct {
				Distance string `json:"distance"`
				Title    string `json:"title"`
				Country  string `json:"country"`
				Altitude string `json:"altitude"`
				Days     []struct {
					Date     string `json:"date"`
					Timeline []struct {
						Interval                      string    `json:"interval"`
						CloudsShortText               string    `json:"clouds_shortText"`
						DdShortText                   string    `json:"dd_shortText"`
						FfShortText                   string    `json:"ff_shortText"`
						FfmaxVal                      string    `json:"ffmax_val"`
						RhShortText                   string    `json:"rh_shortText"`
						DdffIcon                      string    `json:"ddff_icon"`
						WwsynIcon                     string    `json:"wwsyn_icon"`
						CloudsIconWwsynIcon           string    `json:"clouds_icon_wwsyn_icon"`
						WwsynShortText                string    `json:"wwsyn_shortText"`
						PaShortText                   string    `json:"pa_shortText"`
						Msl                           string    `json:"msl"`
						T                             string    `json:"t"`
						Time                          string    `json:"time"`
						Rh                            string    `json:"rh"`
						Valid                         time.Time `json:"valid"`
						FfVal                         string    `json:"ff_val"`
						CloudsShortTextWwsynShortText string    `json:"clouds_shortText_wwsyn_shortText"`
					} `json:"timeline"`
					UTCoffset string    `json:"UTCoffset"`
					Sunset    time.Time `json:"sunset"`
					Sunrise   time.Time `json:"sunrise"`
				} `json:"days"`
				ZoomLevel      string `json:"zoomLevel"`
				LObservHistory string `json:"lObservHistory"`
				ParentID       string `json:"parentId"`
				Type           string `json:"type"`
				ID             string `json:"id"`
			} `json:"properties"`
		} `json:"features"`
		Language    string    `json:"language"`
		IconFormat  string    `json:"icon_format"`
		IconBaseURL string    `json:"icon_base_url"`
		DataType    string    `json:"dataType"`
		TsUpdated   time.Time `json:"tsUpdated"`
		TsValid     string    `json:"tsValid"`
		Params      struct {
			PaShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"pa_shortText"`
			CloudsShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"clouds_shortText"`
			DdShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"dd_shortText"`
			WwsynLongText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"wwsyn_longText"`
			FfShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"ff_shortText"`
			FfmaxVal struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"ffmax_val"`
			RhShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"rh_shortText"`
			WwsynIcon struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"wwsyn_icon"`
			CloudsIconWwsynIcon struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"clouds_icon_wwsyn_icon"`
			Msl struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"msl"`
			T struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"t"`
			Rh struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"rh"`
			DdffIcon struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"ddff_icon"`
			FfVal struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"ff_val"`
		} `json:"params"`
		Type string `json:"type"`
	} `json:"observation"`
	Forecast1H struct {
		Features []struct {
			Geometry struct {
				Type        string    `json:"type"`
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
			Type       string `json:"type"`
			Properties struct {
				ID       string `json:"id"`
				Country  string `json:"country"`
				ParentID string `json:"parentId"`
				Days     []struct {
					Date     string `json:"date"`
					Timeline []struct {
						FfmaxVal                      string    `json:"ffmax_val"`
						RhShortText                   string    `json:"rh_shortText"`
						CloudsShortText               string    `json:"clouds_shortText"`
						DdShortText                   string    `json:"dd_shortText"`
						CloudsIconWwsynIcon           string    `json:"clouds_icon_wwsyn_icon"`
						Valid                         time.Time `json:"valid"`
						Rh                            string    `json:"rh"`
						DdffIcon                      string    `json:"ddff_icon"`
						T                             string    `json:"t"`
						PaShortText                   string    `json:"pa_shortText"`
						Msl                           string    `json:"msl"`
						SnAcc                         string    `json:"sn_acc"`
						WwsynShortText                string    `json:"wwsyn_shortText"`
						WwsynDecodeText               string    `json:"wwsyn_decodeText"`
						FfVal                         string    `json:"ff_val"`
						FfShortText                   string    `json:"ff_shortText"`
						Interval                      string    `json:"interval"`
						WwsynIcon                     string    `json:"wwsyn_icon"`
						CloudBaseShortText            string    `json:"cloudBase_shortText"`
						Time                          string    `json:"time"`
						CloudsShortTextWwsynShortText string    `json:"clouds_shortText_wwsyn_shortText"`
						TpAcc                         string    `json:"tp_acc"`
					} `json:"timeline"`
					UTCoffset string    `json:"UTCoffset"`
					Sunset    time.Time `json:"sunset"`
					Sunrise   time.Time `json:"sunrise"`
				} `json:"days"`
				Title string `json:"title"`
			} `json:"properties"`
		} `json:"features"`
		Language    string    `json:"language"`
		IconFormat  string    `json:"icon_format"`
		IconBaseURL string    `json:"icon_base_url"`
		DataType    string    `json:"dataType"`
		TsUpdated   time.Time `json:"tsUpdated"`
		TsValid     string    `json:"tsValid"`
		Params      struct {
			PaShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"pa_shortText"`
			CloudsShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"clouds_shortText"`
			DdShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"dd_shortText"`
			WwsynDecodeText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"wwsyn_decodeText"`
			FfShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"ff_shortText"`
			FfmaxVal struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"ffmax_val"`
			RhShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"rh_shortText"`
			WwsynIcon struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"wwsyn_icon"`
			CloudsIconWwsynIcon struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"clouds_icon_wwsyn_icon"`
			Rh struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"rh"`
			WwsynShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"wwsyn_shortText"`
			Msl struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"msl"`
			T struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"t"`
			SnAcc struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"sn_acc"`
			TpAcc struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"tp_acc"`
			DdffIcon struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"ddff_icon"`
			FfVal struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"ff_val"`
			CloudBaseShortText struct {
				Name string `json:"name"`
				Unit string `json:"unit"`
				Desc string `json:"desc"`
			} `json:"cloudBase_shortText"`
		} `json:"params"`
		Type string `json:"type"`
	} `json:"forecast1h"`
}

func GetArso(kraj string) ARSO {
	url := fmt.Sprintf("http://www.vreme.si/api/1.0/location/?lang=sl&location=%s&format=json", kraj)

	spaceClient := http.Client{
		Timeout: time.Second * 12, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		panic(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		panic(readErr)
	}

	arso := ARSO{}
	jsonErr := json.Unmarshal(body, &arso)
	if jsonErr != nil {
		panic(jsonErr)
	}
	return arso
}
