package maps

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

/* Slice 格式化 */
func iterSlice(arr interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(arr)
	ret := make(map[string]interface{})
	
	if v.Kind() != reflect.Slice {
		return ret, errors.New("not match Slice")
	}
	l := v.Len()
	for i := 0; i < l; i++ {
		/* 将结果转换成str格式符合输出 */
		strI := strconv.Itoa(i)
		
		ret[strI] = v.Index(i).Interface()
	}
	
	return ret, nil
}

// IterMap map 格式化
func IterMap(arr map[string]interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(arr)
	ret := make(map[string]interface{})
	
	if v.Kind() != reflect.Map {
		return ret, errors.New("not match Map")
	}
	return arr, nil
}

// flatMap 平铺展开
func flatMap(arr map[string]interface{}) (map[string]interface{}, error) {
	
	v := reflect.ValueOf(arr)
	
	if v.Kind() != reflect.Map {
		return arr, errors.New("not match Map")
	}
	
	reMaps := make(map[string]interface{})
	
	for keys, values := range arr {
		/* 判断里面value的类型 */
		vValues := reflect.ValueOf(values)
		
		switch vValues.Kind() {
		case reflect.Slice:
			ret, _ := flatSlice(values)
			
			for miniK, miniV := range ret {
				fixStr := keys + "." + miniK
				reMaps[fixStr] = miniV
			}
		
		case reflect.Map:
			rec, _ := values.(map[string]interface{})
			ret, _ := flatMap(rec)
			
			for miniK, miniV := range ret {
				fixStr := keys + "." + miniK
				reMaps[fixStr] = miniV
			}
		default:
			reMaps[keys] = values
			
		}
		
	}
	return reMaps, nil
	
}

/* Slice 平铺展开*/
func flatSlice(arr interface{}) (map[string]interface{}, error) {
	
	reMaps := make(map[string]interface{})
	retRes, _ := iterSlice(arr)
	
	for keys, values := range retRes {
		v := reflect.ValueOf(values)
		switch v.Kind() {
		case reflect.Slice:
			ret, _ := flatSlice(values)
			
			for miniK, miniV := range ret {
				fixStr := keys + "." + miniK
				reMaps[fixStr] = miniV
			}
		case reflect.Map:
			rec, _ := values.(map[string]interface{})
			ret, _ := flatMap(rec)
			
			for miniK, miniV := range ret {
				fixStr := keys + "." + miniK
				reMaps[fixStr] = miniV
			}
		default:
			reMaps[keys] = values
			
		}
	}
	return reMaps, nil
}

// TryToGet 将多维json 扁平化输出为一维
func TryToGet(arr interface{}) (map[string]interface{}, error) {
	/* 获取 输入 的属性 */
	v := reflect.ValueOf(arr)
	
	switch v.Kind() {
	case reflect.Slice:
		ret, err := flatSlice(arr)
		
		return ret, err
	case reflect.Map:
		/* 强制替换成 map */
		rec, _ := arr.(map[string]interface{})
		
		ret, err := flatMap(rec)
		
		return ret, err
	
	default:
		ret := make(map[string]interface{})
		/* 既不属于 Slice 也不属于 map */
		return ret, errors.New("not map or Slice")
		
	}
}

// KeyValuePairs 将一维的json 输出成str
func KeyValuePairs(m map[string]interface{}) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		_, err := fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
		if err != nil {
			return ""
		}
	}
	return b.String()
}

// ToSlice 将切片结果输出
func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
