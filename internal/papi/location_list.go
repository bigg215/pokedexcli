package papi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := Locations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Locations{}, err
		}
		return locationsResp, nil
	}

	dat, err := c.processGetRequest(url)

	if err != nil {
		return Locations{}, nil
	}

	locationsResp := Locations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil

}
