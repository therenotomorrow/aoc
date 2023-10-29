package y2015

import "encoding/json"

func sumNumsSafe(v interface{}) float64 {
	sum := 0.0

	switch val := v.(type) {
	case float64:
		sum = val
	case []interface{}:
		for _, obj := range val {
			sum += sumNumsSafe(obj)
		}
	case map[string]interface{}:
		for _, obj := range val {
			sum += sumNumsSafe(obj)
		}
	}

	return sum
}

func Day12Part1(jsonStr string) int {
	var obj interface{}

	_ = json.Unmarshal([]byte(jsonStr), &obj)

	return int(sumNumsSafe(obj))
}

func sumNumsUnsafe(v interface{}) float64 {
	sum := 0.0

	switch val := v.(type) {
	case float64:
		sum = val
	case []interface{}:
		for _, obj := range val {
			sum += sumNumsUnsafe(obj)
		}
	case map[string]interface{}:
		mapSum := 0.0

		for _, obj := range val {
			trueObj, ok := obj.(string)

			if ok && trueObj == "red" {
				mapSum = 0.0
				break
			}

			mapSum += sumNumsUnsafe(obj)
		}

		sum += mapSum
	}

	return sum
}

func Day12Part2(jsonStr string) int {
	var obj interface{}

	_ = json.Unmarshal([]byte(jsonStr), &obj)

	return int(sumNumsUnsafe(obj))
}
