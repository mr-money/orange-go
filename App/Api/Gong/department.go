package Gong

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
)

//
// FormatDep
// @Description: 常高新部门格式化
// @param c
//
func FormatDep(c *gin.Context) {
	depJson := c.PostForm("data")

	depMap := map[string]map[string]string{}

	err := json.Unmarshal([]byte(depJson), &depMap)
	if err != nil {
		c.JSON(501, gin.H{"err": err})
	}

	//获取用户map
	nameMap := depMap["name"]
	delete(depMap, "name")

	//排序
	var depKeySort []string
	for depKey, _ := range depMap {
		//获取部门key有序索引
		depKey = depKey[3:]
		//赋值切片
		depKeySort = append(depKeySort, depKey)
	}
	//降序
	sort.Slice(depKeySort, func(i, j int) bool {
		return depKeySort[i] > depKeySort[j]
	})

	//顶级公司关键词定义
	keywords := []string{
		"公司",
		"伙)",
		"营层",
		"伙）",
	}

	//定义返回
	resList := []string{}

	//循环人员
nameLoop:
	for nameKey, name := range nameMap {
		//fmt.Println(name, "-------")
		//循环降序部门
		nameRes := name + "/"
		for _, sortKey := range depKeySort {
			//筛选空部门
			if depMap["dep"+sortKey][nameKey] == "" {
				continue
			}

			nameRes += depMap["dep"+sortKey][nameKey] + "/"

			//判断顶级公司
			for _, keyword := range keywords {
				if strings.Contains(nameRes, keyword) {
					fmt.Println(nameRes)
					resList = append(resList, nameRes)
					continue nameLoop
				}
			}
		}

		resList = append(resList, nameRes)
	}

	c.JSON(200, gin.H{"res": resList})
}
