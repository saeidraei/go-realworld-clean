package uc

import (
	"encoding/json"
	"fmt"
	"github.com/saeidraei/go-realworld-clean/domain"
	"time"
)

func (i interactor) UrlGet(id string) (*domain.Url, error) {

	if i.cacheRW != nil {
		val, err := i.cacheRW.Get("url:" + id)
		if err!=nil{
			fmt.Println("error:", err)
		}
		if val != nil {
			var url domain.Url
			err := json.Unmarshal([]byte(val.(string)), &url)
			if err != nil {
				fmt.Println("error:", err)
			}
			return &url, nil
		}
	}
	url, err := i.urlRW.GetByID(id)

	if err != nil {
		return nil, err
	}
	if i.cacheRW != nil {
		b, err := json.Marshal(url)
		err = i.cacheRW.Set("url:"+id, b, 60*time.Second)
		if err != nil {
			fmt.Println("error:", err)
		}
	}
	return url, nil
}
