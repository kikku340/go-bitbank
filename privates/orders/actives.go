package orders

import (
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
	u.Path = path.Join(VERSION, PATH, "cancel_order")

	j, err := json.Marshal(b)
	if err != nil {
		return ActiveOrder{}, err
	}

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return ActiveOrder{}, err
	}

	auth.MakeHeader(p.Token, p.Secret, m, req)

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		return ActiveOrder{}, err
	}
	defer res.Body.Close()

	var resp Response
	json.NewDecoder(res.Body).Decode(&resp)
	if resp.Success != 1 {
		return ActiveOrder{}, e.Handler(resp.Data.Code, err)
	}

	return resp.Data, nil
}
