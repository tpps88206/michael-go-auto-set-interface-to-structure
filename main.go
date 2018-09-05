package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	var (
		person     Person
		parameters []string
	)

	parameters = append(parameters, "Name", "Age", "Email")
	data := []byte(`{"Name":"Michael", "Age":25, "Email":example@gmail.com}`)

	personSource := reflect.ValueOf(&person)

	err := mapValue(parameters, data, &personSource)
	if err != nil {
		panic(err)
	}

	fmt.Println(personSource)
}

// init source data structure and check valid
func initSource(source *reflect.Value) (reflect.Value, error) {
	var sourceElement reflect.Value

	sourceElement = source.Elem()
	if !sourceElement.IsValid() {
		return sourceElement, errors.New("ERROR : The source structure is not valid")
	}

	if sourceElement.Kind() != reflect.Struct {
		return sourceElement, errors.New("ERROR : The response structure return Invalid")
	}

	return sourceElement, nil
}

// set value to source from data in parameter
func setValue(parameter string, data interface{}, sourceElement *reflect.Value) error {
	dataValue := reflect.ValueOf(data)
	if !dataValue.IsValid() {
		return errors.New("ERROR : The data of  '" + parameter + "' from database is not valid")
	}

	elementByParameter := sourceElement.FieldByName(parameter)
	if !elementByParameter.IsValid() {
		return errors.New("ERROR : Cannot match '" + parameter + "' from response structure")
	}

	if !elementByParameter.CanSet() {
		return errors.New("ERROR :  '" + parameter + "' cannot be changed. Maybe it is addressable and was not obtained by the use of unexported struct fields.")
	}

	elementByParameter.Set(dataValue)
	return nil
}

// assign value to source from data in the specific parameter
func assignValue(parameter string, data interface{}, source *reflect.Value) error {
	sourceElement, err := initSource(source)
	if err != nil {
		return err
	}

	return setValue(parameter, data, &sourceElement)
}

// map all value to source from data in every parameter which from parameters array
func mapValue(parameters []string, data interface{}, source *reflect.Value) error {
	sourceElement, err := initSource(source)
	if err != nil {
		return err
	}

	for _, parameter := range parameters {
		value, boolean := function.GetProp(data, parameter)
		if !boolean {
			return errors.New("ERROR : Cannot match '" + parameter + "' from database result")
		}

		err = setValue(parameter, value, &sourceElement)
		if err != nil {
			return err
		}
	}

	return nil
}
