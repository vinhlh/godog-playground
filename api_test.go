package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cucumber/godog"
)

var restaurantCode string
var actualResponse *Response

type Category struct {
	Name string `json:"name"`
}

type Menu struct {
	Name string `json:"name"`
	Categories []Category `json:"menu_categories"`
}

type ResponseData struct {
	Menus []Menu `json:"menus"`
}

type Response struct {
	Data ResponseData `json:"data"`
}

func aRestaurant(vendorCode string) error {
	restaurantCode = vendorCode
	return nil
}

func iVisitRetaurantDetailPage() error {
	endpoint := fmt.Sprintf("https://sg.fd-api.com/api/v5/vendors/%s?include=menus&language_id=1&dynamic_pricing=0&opening_type=pickup", restaurantCode)
	resp, err := http.Get(endpoint)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var r Response
	if err = json.Unmarshal(body, &r); err != nil {
		return err
	}
	actualResponse = &r
	return nil
}

func thereShouldBeMenu(expected string) error {
	if len(actualResponse.Data.Menus) == 0 {
		return fmt.Errorf("no menu found")
	}

	actual := actualResponse.Data.Menus[0].Name
	if actual != expected {
		return fmt.Errorf("expected %s but actual %s", expected, actual)
	}

	return nil
}

func thereShouldBeCategory(expected string) error {
	if len(actualResponse.Data.Menus) == 0 {
		return fmt.Errorf("no menu found")
	}

	categories := actualResponse.Data.Menus[0].Categories
	actualCategoryByValue := map[string]bool{}
	for _, c := range categories {
		actualCategoryByValue[c.Name] = true
	}

	expectedCategories := strings.Split(expected, ", ")
	for _, c := range expectedCategories {
		if ok, _ := actualCategoryByValue[c]; !ok {
			return fmt.Errorf("expected %s but not found", c)
		}
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		restaurantCode = ""
		actualResponse = nil
	})

	ctx.Step(`^a restaurant "([^"]*)"$`, aRestaurant)
	ctx.Step(`^I visit retaurant detail page$`, iVisitRetaurantDetailPage)
	ctx.Step(`^there should be "([^"]*)" menu$`, thereShouldBeMenu)
	ctx.Step(`^there should be "([^"]*)" category$`, thereShouldBeCategory)
}
