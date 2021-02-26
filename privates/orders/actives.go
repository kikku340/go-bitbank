package orders

import (
	"bytes"
	"net/url"

	e "github.com/go-numb/go-bitbank/errors"

	"encoding/json"
	"net/http"
	"path"

	"github.com/go-numb/go-bitbank/privates/auth"
)

func (p *Request) ActiveOrders(b *ActiveOrderBody) (ActiveOrder, error) {
	u, err := url.ParseRequestURI(BASEURL)
	if err != nil {
		return ActiveOrder{}, err
	}
	u.Path = path.Join(VERSION, PATH, "active_orders")

	j, err := json.Marshal(b)
	if err != nil {
		return ActiveOrder{}, err
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(j))
	if err != nil {
		return ActiveOrder{}, err
	}

	auth.MakeHeader(p.Token, p.Secret, nil, req)

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return ActiveOrder{}, err
	}
	defer res.Body.Close()

	var resp ActiveOrderResponse
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return ActiveOrder{}, e.Handler(resp.Data.Code, err)
	}

	return resp.Data, nil
}
