package Handler

//
// removeRepByMap
// @Description:通过字典来去重
// @param slc
// @return []interface{}
//
func removeRepByMap(slc []interface{}) []interface{} {
	result := []interface{}{}
	tempMap := map[interface{}]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

//
// removeRepByLoop
// @Description: 通过循环去重
// @param slc
// @return []interface{}
//
func removeRepByLoop(slc []interface{}) []interface{} {
	result := []interface{}{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

//
// RemoveRep
// @Description: 切片去重
// @param slc
// @return []interface{}
//
func RemoveRep(slc []interface{}) []interface{} {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return removeRepByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return removeRepByMap(slc)
	}
}
