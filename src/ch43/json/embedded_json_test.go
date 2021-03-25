package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}

var jsonStr1 = `{
	"base":{
		"name":"tyf",
		"age":29
	},
	"skills":{
		"technologies":["Java","Go","C"]
	}
}`

func TestMyJson(t *testing.T) {
	person := new(Person)
	json.Unmarshal([]byte(jsonStr1), person)
	t.Log(person)
	if marshal, err := json.Marshal(person); err != nil {
		t.Error("出错")
	} else {
		t.Log(string(marshal))
	}
}
