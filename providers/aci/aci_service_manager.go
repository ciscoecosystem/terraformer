package aci

import (
	"errors"
	"log"
	"strconv"

	container "github.com/Jeffail/gabs"
)

type ACIServiceManager struct {
	MOURL  string
	client *ACIClient
}

func NewServiceManager(moURL string, client *ACIClient) *ACIServiceManager {

	sm := &ACIServiceManager{
		MOURL:  moURL,
		client: client,
	}
	return sm
}

// CheckForErrors parses the response and checks of there is an error attribute in the response
func CheckForErrors(cont *container.Container, method string, skipLoggingPayload bool) error {
	number, err := strconv.Atoi(G(cont, "totalCount"))
	if err != nil {
		if !skipLoggingPayload {
			log.Printf("[DEBUG] Exit from errors %v", cont)
		} else {
			log.Printf("[DEBUG] Exit from errors %s", err.Error())
		}
		return err
	}
	imdata := cont.S("imdata").Index(0)
	if number > 0 {

		if imdata.Exists("error") {

			if stripQuotes(imdata.Path("error.attributes.code").String()) == "103" {
				if !skipLoggingPayload {
					log.Printf("[DEBUG] Exit from errors %v", cont)
				}
				return nil
			} else {
				if stripQuotes(imdata.Path("error.attributes.text").String()) == "" && stripQuotes(imdata.Path("error.attributes.code").String()) == "403" {
					if !skipLoggingPayload {
						log.Printf("[DEBUG] Exit from errors %v", cont)
					}
					return errors.New("Unable to authenticate. Please check your credentials")
				}
				if !skipLoggingPayload {
					log.Printf("[DEBUG] Exit from errors %v", cont)
				}

				return errors.New(stripQuotes(imdata.Path("error.attributes.text").String()))
			}
		}

	}

	if imdata.String() == "{}" && method == "GET" {
		if !skipLoggingPayload {
			log.Printf("[DEBUG] Exit from errors %v", cont)
		}

		return errors.New("Error retriving Object: Object may not exists")
	}
	if !skipLoggingPayload {
		log.Printf("[DEBUG] Exit from errors %v", cont)
	}
	return nil
}

func (sm *ACIServiceManager) GetViaURL(url string) (*container.Container, error) {
	req, err := sm.client.MakeRestRequest("GET", url, nil, true)

	if err != nil {
		return nil, err
	}

	obj, _, err := sm.client.Do(req)
	// if !sm.client.skipLoggingPayload {
	// 	log.Printf("Getvia url %+v", obj)
	// }
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, errors.New("Empty response body")
	}
	return obj, CheckForErrors(obj, "GET", sm.client.skipLoggingPayload)

}
