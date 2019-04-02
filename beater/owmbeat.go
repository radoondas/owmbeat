package beater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/radoondas/owmbeat/config"
)

// Owmbeat configuration.
type Owmbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

const (
	owmApiUrl    = "http://api.openweathermap.org"
	owmApiBoxUri = "data/2.5/box/city"

	selector = "owmbeat"
)

type OwmResponseData struct {
	Cod          int       `json:"cod"`
	Calctime     float32   `json:"calctime"`
	Cnt          int       `json:"cnt"`
	Measurements []Measure `json:"list"`
}

type Measure struct {
	CityId    int          `json:"id"`
	Timestamp int          `json:"dt"`
	Name      string       `json:"name"`
	Coord     Coordination `json:"coord"`
	Main      Main         `json:"main"`
	Wind      Wind         `json:"wind"`
	Rain      Rain         `json:"rain"`
	Snow      Snow         `json:"snow"`
	Clouds    Clouds       `json:"clouds"`
	Weather   []Weather    `json:"weather"`
}

type Coordination struct {
	Lat float32 `json:"Lat"`
	Lon float32 `json:"Lon"`
}

type Main struct {
	Temp       float32 `json:"temp"`
	Temp_min   float32 `json:"temp_min"`
	Temp_max   float32 `json:"temp_max"`
	Pressure   float32 `json:"pressure"`
	Humidity   int     `json:"humidity"`
	Sea_level  float32 `json:"sea_level"`
	Grnd_level float32 `json:"grnd_level"`
}

type Wind struct {
	Speed float32 `config:"speed"`
	Deg   float32 `config:"deg"`
}

type Clouds struct {
	Today int `json:"today"`
}

type Rain struct {
	OneHour    float32 `json:"1h"`
	ThreeHours float32 `json:"3h"`
}

type Snow struct {
	OneHour    float32 `json:"1h"`
	ThreeHours float32 `json:"3h"`
}

type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

// http://api.openweathermap.org/data/2.5/box/city?bbox=16,47,23,50,10&appid=9ae766cc7027b713ac80c6458ec911b5
// bbox bounding box [lon-left,lat-bottom,lon-right,lat-top,zoom]

// New creates an instance of owmbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	bt := &Owmbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts owmbeat.
func (bt *Owmbeat) Run(b *beat.Beat) error {
	logp.NewLogger(selector).Info("owmbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	for _, r := range bt.config.Regions {
		go func(r config.Region) {
			ticker := time.NewTicker(bt.config.Period)
			defer ticker.Stop()

			for {
				logp.NewLogger(selector).Info("Region: ", r.Name)
				err := bt.GetOWM(r)
				if err != nil {
					logp.NewLogger(selector).Error("Error while getting OWM data: %v", err)
				}

				select {
				case <-bt.done:
					goto GotoFinish
				case <-ticker.C:
				}
			}
		GotoFinish:
		}(r)
	}

	<-bt.done

	return nil
}

// Stop stops owmbeat.
func (bt *Owmbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func (bt *Owmbeat) GetOWM(region config.Region) error {

	var ParsedUrl *url.URL
	client := &http.Client{}

	ParsedUrl, err := url.Parse(owmApiUrl)
	if err != nil {
		logp.NewLogger(selector).Error("Unable to parse URL string")
		panic(err)
	}

	ParsedUrl.Path += owmApiBoxUri

	parameters := url.Values{}
	parameters.Add("appid", bt.config.AppId)

	bbox := strconv.Itoa(
		region.LonLeft) + "," +
		strconv.Itoa(region.LatBottom) + "," +
		strconv.Itoa(region.LonRight) + "," +
		strconv.Itoa(region.LatTop) + "," +
		strconv.Itoa(region.Zoom)
	parameters.Add("bbox", bbox)

	ParsedUrl.RawQuery = parameters.Encode()

	logp.NewLogger(selector).Debug("Requesting OWM data: ", ParsedUrl.String())

	req, err := http.NewRequest("GET", ParsedUrl.String(), nil)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		logp.NewLogger(selector).Debug("Status code: ", res.StatusCode)
		logp.NewLogger(selector).Debug("Status code: ", res.Body)
		return fmt.Errorf("HTTP %v", res)
	}

	body, err := ioutil.ReadAll(res.Body)
	logp.NewLogger(selector).Debug(body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	owmdata := OwmResponseData{}
	err = json.Unmarshal([]byte(body), &owmdata)
	if err != nil {
		fmt.Printf("error: %v", err)
		//panic(err)
		return err
	}

	logp.NewLogger(selector).Debug("Unmarshal-ed Owm data: ", owmdata)

	transformedData := bt.TransformOwmData(owmdata, region.Name, region.Description)

	ts := time.Now()
	for _, d := range transformedData {

		event := beat.Event{
			Timestamp: ts,
			Fields: common.MapStr{
				"type":           "owmbeat",
				"openweathermap": d,
			},
		}

		bt.client.Publish(event)
		logp.NewLogger(selector).Debug("Event: ", event)
	}

	return nil
}

func (bt *Owmbeat) TransformOwmData(data OwmResponseData, regionName string, regionDesc string) []common.MapStr {

	measurements := []common.MapStr{}

	for _, m := range data.Measurements {
		measure := common.MapStr{
			"region": common.MapStr{
				"name":        regionName,
				"description": regionDesc,
			},
			"city": common.MapStr{
				"id":   m.CityId,
				"name": m.Name,
			},
			"timestamp": m.Timestamp * 1000,
			"location": common.MapStr{
				"lat": m.Coord.Lat,
				"lon": m.Coord.Lon,
			},
			"main": common.MapStr{
				"temp":      m.Main.Temp,
				"tempMin":   m.Main.Temp_min,
				"tempMax":   m.Main.Temp_max,
				"pressure":  m.Main.Pressure,
				"humidity":  m.Main.Humidity,
				"seaLevel":  m.Main.Sea_level,
				"grndLevel": m.Main.Grnd_level,
			},
			"wind": common.MapStr{
				"speed": m.Wind.Speed,
				"deg":   m.Wind.Deg,
			},
			"rain": common.MapStr{
				"1h": m.Rain.OneHour,
				"3h": m.Rain.ThreeHours,
			},
			"snow": common.MapStr{
				"1h": m.Snow.OneHour,
				"3h": m.Snow.ThreeHours,
			},
			"clouds": common.MapStr{
				"today": m.Clouds.Today,
			},
			// We do not use Weather[] section
		}

		measurements = append(measurements, measure)
	}

	return measurements
}
