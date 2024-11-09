package papi

import "encoding/json"

func (c *Client) GetSpecies(url string) (Species, error) {
	if val, ok := c.cache.Get(url); ok {
		speciesResp := Species{}
		err := json.Unmarshal(val, &speciesResp)
		if err != nil {
			return Species{}, err
		}
		return speciesResp, nil
	}

	dat, err := c.processGetRequest(url)

	if err != nil {
		return Species{}, nil
	}

	speciesResp := Species{}
	err = json.Unmarshal(dat, &speciesResp)
	if err != nil {
		return Species{}, err
	}
	c.cache.Add(url, dat)

	return speciesResp, nil
}
