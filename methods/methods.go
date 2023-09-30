package methods

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

func Map[T comparable, V comparable](slice []T, f func(param T) V) []V {

	newList := []V{}
	for _, v := range slice {
		newList = append(newList, f(v))
	}

	return newList
}

func Filter[T comparable](slice []T, f func(param T) bool) []T {

	newList := []T{}
	for _, v := range slice {
		if f(v) {
			newList = append(newList, v)
		}
	}

	return newList
}

func Reduce[T comparable, V comparable](slice []T, f func(param T) V) V {

	var result V
	for _, v := range slice {
		result = f(v)
	}

	return result
}

func Find[T comparable](slice []T, f func(param T) bool) []T {

	newList := []T{}

	for _, v := range slice {
		if f(v) {
			newList = append(newList, v)
		}
	}

	return newList
}

func FindFirst[T comparable](slice []T, f func(param T) bool) T {

	for _, v := range slice {
		if f(v) {
			return v
		}
	}

	return slice[0]
}

func FindLast[T comparable](slice []T, f func(param T) bool) []T {

	newList := []T{}

	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			newList = append(newList, slice[i])
		}
	}

	return newList
}

func FindIndex[T comparable](slice []T, f func(param T) bool) int {

	for i, v := range slice {
		if f(v) {
			return i
		}
	}

	return -1
}

func FindLastIndex[T comparable](slice []T, f func(param T) bool) int {

	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}

	return -1
}

func Every[T comparable](slice []T, f func(param T) bool) bool {

	for _, v := range slice {
		if !f(v) {
			return false
		}
	}

	return true
}

func Some[T comparable](slice []T, f func(param T) bool) bool {

	for _, v := range slice {
		if f(v) {
			return true
		}
	}

	return false
}

func IndexOf[T comparable](slice []T, f func(param T) bool) int {

	for i, v := range slice {
		if f(v) {
			return i
		}
	}

	return -1
}

func LastIndexOf[T comparable](slice []T, f func(param T) bool) int {

	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}

	return -1
}

func Includes[T comparable](slice []T, f func(param T) bool) bool {

	for _, v := range slice {
		if f(v) {
			return true
		}
	}

	return false
}

func ForEach[T comparable](slice []T, f func(param T)) {

	for _, v := range slice {
		f(v)
	}
}

func ConvertJsonInObjectMongo(val interface{}, convert interface{}) error {
	jsonData, err := bson.MarshalExtJSON(val, false, false)

	if err != nil {
		return err
	}

	err = bson.UnmarshalExtJSON(jsonData, false, convert)

	if err != nil {
		return err
	}

	return nil
}

func ConvertJsonInObject(val interface{}, convert interface{}) error {
	jsonData, err := json.Marshal(val)

	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, convert)

	if err != nil {
		return err
	}

	return nil
}
