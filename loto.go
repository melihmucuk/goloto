package goloto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Cekilis) GetDates() ([]Date, error) {
	url := fmt.Sprintf("%s/listCekilisleriTarihleri.php?tur=%s", BaseURL, c.Tur)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var dates []Date
	jsonErr := json.Unmarshal(body, &dates)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return dates, nil
}

func (c *Cekilis) GetResults(date Date) (Result, error) {
	url := fmt.Sprintf("%s/cekilisler/%s/%s.json", BaseURL, c.Tur, date.Tarih)
	resp, err := http.Get(url)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result{}, err
	}
	var results Result
	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		return Result{}, jsonErr
	}
	return results, nil
}
