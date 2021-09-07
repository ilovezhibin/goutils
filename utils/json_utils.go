package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type JSONObject struct {
	jsonMap map[string]interface{}
}

func (object *JSONObject) setJsonMap(jsomMap map[string]interface{}) {
	object.jsonMap = jsomMap
}

func (object *JSONObject) GetJSONObject(key string) (JSONObject, error) {
	value := object.jsonMap[key]
	jsonObject := JSONObject{}
	switch value.(type) {
	case map[string]interface{}:
		temp := value.(map[string]interface{})
		jsonObject.jsonMap = temp
		return jsonObject, nil
	}
	return jsonObject, fmt.Errorf("类型不对")
}

func (object *JSONObject) GetFloat64Array(key string) ([]float64, error) {
	value := object.jsonMap[key]
	switch value.(type) {
	case []interface{}:
		temp := value.([]interface{})
		fmt.Println("cap=" + strconv.Itoa(len(temp)))
		array := make([]float64, len(temp))
		for index, v := range temp {
			switch v.(type) {
			case float32:
				array[index] = float64(v.(float32))
			case float64:
				array[index] = v.(float64)
			case uint8, uint16, uint32, uint64, int8, int16, int32, int64, int, uint:
				data, _ := json.Marshal(value)
				num, _ := strconv.ParseFloat(string(data), 64)
				array[index] = num

			default:
				return nil, fmt.Errorf("子类型不对")
			}
		}
		return array, nil
	default:
		return nil, fmt.Errorf("父类型不对")
	}
}

func (object *JSONObject) GetInt64Array(key string) ([]int64, error) {
	value := object.jsonMap[key]
	switch value.(type) {
	case []interface{}:
		temp := value.([]interface{})
		fmt.Println("cap=" + strconv.Itoa(len(temp)))
		array := make([]int64, len(temp))
		for index, v := range temp {
			switch v.(type) {
			case float32:
				array[index] = int64(v.(float32))
			case float64:
				array[index] = int64(v.(float64))
			case uint8, uint16, uint32, uint64, int8, int16, int32, int64, int, uint:
				data, _ := json.Marshal(value)
				num, _ := strconv.ParseInt(string(data), 10, 64)
				array[index] = num
			default:
				return nil, fmt.Errorf("子类型不对")
			}
		}
		return array, nil
	default:
		return nil, fmt.Errorf("父类型不对")
	}
}

func (object *JSONObject) GetInt64(key string) (num int64, err error) {
	str := object.GetString(key)
	num, err = strconv.ParseInt(str, 10, 64)
	return
}

func (object *JSONObject) GetInt(key string) (num int, err error) {
	str := object.GetString(key)
	num, err = strconv.Atoi(str)
	return
}

func (object *JSONObject) GetString(key string) string {
	value := object.jsonMap[key]
	if value == nil {
		return ""
	}
	//case uint8,uint16,uint32,uint64,int8,int16,int32,int64,int,uint,[]interface{},map[string]interface{}:
	switch value.(type) {
	case string:
		return value.(string)
	default:
		data, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		return string(data)
	}
}

func ParseJSONObject(str string) (*JSONObject, error) {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &jsonMap)
	if err != nil {
		return nil, err
	}
	jsonObject := new(JSONObject)
	jsonObject.setJsonMap(jsonMap)
	return jsonObject, nil
}
