package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type jsonTagGetter struct {
	cache map[string][]string
}

func (g *jsonTagGetter) Get(e any) ([]string, error) {
	if g.cache == nil {
		g.cache = make(map[string][]string)
	}

	// TODO: add check for type
	key := reflect.TypeOf(e).Name()

	if tags, ok := g.cache[key]; ok {
		fmt.Println("cache hit")
		return tags, nil
	}

	tags, err := GetTagsJSON(e)
	if err != nil {
		return nil, err
	}

	g.cache[key] = tags

	return tags, nil
}

func GetTagsJSON(e any) ([]string, error) {
	t := reflect.TypeOf(e)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t == nil {
		return nil, errors.New("nil input")
	}

	if k := t.Kind(); k != reflect.Struct {
		return nil, fmt.Errorf("invalid kind of input: %v", k)
	}

	var allFieldNames []string

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		tagValue := f.Tag.Get("json")
		if tagValue == "" {
			continue
		}

		fieldName := strings.Split(tagValue, ",")[0]

		allFieldNames = append(allFieldNames, fieldName)
	}

	return allFieldNames, nil
}
