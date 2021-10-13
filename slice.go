package goutils

func Int64Unique(reqList []int64) []int64 {
	var result []int64
	for _, v1 := range reqList {
		isExist := false
		for _, v2 := range result {
			if v1 == v2 {
				isExist = true
				break
			}
		}
		if !isExist {
			result = append(result, v1)
		}
	}
	return result
}

func Int64Filter(req []int64) []int64 {
	var result []int64
	for _, v := range req {
		if v > 0 {
			result = append(result, v)
		}
	}
	return result
}

func IntUnique(reqList []int) []int {
	var result []int
	for _, v1 := range reqList {
		isExist := false
		for _, v2 := range result {
			if v1 == v2 {
				isExist = true
				break
			}
		}
		if !isExist {
			result = append(result, v1)
		}
	}
	return result
}

func StringUnique(reqList []string) []string {
	var result []string
	for _, v1 := range reqList {
		if v1 == "" {
			continue
		}
		isExist := false
		for _, v2 := range result {
			if v1 == v2 {
				isExist = true
				break
			}
		}
		if !isExist {
			result = append(result, v1)
		}
	}
	return result
}

func SliceUnique(reqList interface{}) interface{} {
	if v, ok := reqList.([]int64); ok {
		return Int64Unique(v)
	}
	if v, ok := reqList.([]int); ok {
		return IntUnique(v)
	}
	if v, ok := reqList.([]string); ok {
		return StringUnique(v)
	}
	return nil
}
