package oas

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewCaseGroup(name string) *group {
	return &group{
		name: name,
	}
}

type group struct {
	name string
	list []*caseItem
}

func (g *group) It(desc string, result string, v interface{}) {
	g.list = append(g.list, &caseItem{
		desc:   desc,
		result: result,
		value:  v,
	})
}

func (g *group) Run(t *testing.T) {
	for i := range g.list {
		item := g.list[i]
		data, errForMarshal := json.Marshal(item.value)
		assert.Nil(t, errForMarshal)
		assert.Equal(t,
			item.result,
			string(data),
			fmt.Sprintf("[%s] %s, marshal failed, results: %s", g.name, item.desc, string(data)),
		)

		expectRv := reflect.Indirect(reflect.ValueOf(item.value))
		value := reflect.New(expectRv.Type()).Interface()
		errForUnmarshal := json.Unmarshal([]byte(item.result), value)
		assert.Nil(t, errForUnmarshal)
		if errForUnmarshal != nil {
			t.Logf("%s", errForUnmarshal.Error())
		}
		value = reflect.Indirect(reflect.ValueOf(value)).Interface()
		assert.Equal(t,
			expectRv.Interface(),
			value,
			fmt.Sprintf("[%s] %s, unmarshal failed)", g.name, item.desc),
		)
	}
}

type caseItem struct {
	value  interface{}
	result string
	desc   string
}
