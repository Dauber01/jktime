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
		"skills":["java", "go", "c"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Fatal(err)
	}
}
