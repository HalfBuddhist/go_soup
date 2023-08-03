package ts_simplejson

import (
	"encoding/json"
	"fmt"
	"testing"

	sjson "github.com/bitly/go-simplejson"
)

func TestJsonParse(t *testing.T) {
	resp := `{"code": "00",
				   "message": "SUCCESS",
				   "describe": "成功",
				   "resultInfo": { "uniqueNumber": "201808161133401673324075025000035" },
				   "test_slice": [1,2,3],
				   "data": {
					"id": 6,
					"mid": 2,
					"type": "Dir",
					"qid": 5,
					"path": "/defaultClient/zzzz",
					"dirino": 16007630338,
					"name": null,
					"cursize": 0,
					"maxsize": 1293942784,
					"curfiledirs": 0,
					"maxfiledirs": 0,
					"uid_gid_ino": 0,
					"createTime": 1715585897326,
					"alarmRadio": 0.8,
					"alarmFileNumRadio": 0.6,
					"alarmSizeRadio": 0.7,
					"applicationType": null,
					"soft": false,
					"lastModifyDate": 0
				  }
				 }`
	js, errs := sjson.NewJson([]byte(resp))
	if errs != nil {
		return
	}
	discount := js.Get("resultInfo").Get("uniqueNumber")
	fmt.Println(discount)
	strcode, _ := js.Get("code").String()
	fmt.Println(strcode)
	intcode, _ := js.Get("code").Int()
	fmt.Println(intcode)
	path := js.GetPath("resultInfo", "uniqueNumber")
	fmt.Println(path)
	// get empty field
	errType, err := js.Get("error").String()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result: ", errType)
	}
	// get array item, index start trom 0
	a, _ := js.Get("test_slice").GetIndex(1).Int()
	fmt.Println(a)
	// int type
	dataMap := js.Get("data").MustMap()
	qidNumber, ok := dataMap["qid"].(json.Number)
	if ok {
		fmt.Println("is json number")
		qid, _ := qidNumber.Int64()
		fmt.Printf("%d\n", qid)
	} else {
		fmt.Printf("%v\n", qidNumber)
	}
	fmt.Printf("%T", js.Get("resultInfo"))
}
