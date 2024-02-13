package ini

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Ini(fileName string) (res map[string]map[string]interface{}, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	res = make(map[string]map[string]interface{})
	currentSection := ""
	lineslice := bufio.NewScanner(f)
	for lineslice.Scan() {
		if lineslice.Err() != nil {
			err = lineslice.Err()
			return nil, err
		}
		line := lineslice.Text()
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue // 空行
		}

		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue // 注释行
		}

		index := strings.IndexAny(line, ";#") // 查找分号和井号的最小索引
		if index != -1 {
			line = line[:index] // 截取注释之前的内容
		}
//完成脏数据剔除

		//获取节
		if strings.HasPrefix(line, "[") {
			if !strings.HasSuffix(line, "]") {
				err := fmt.Errorf("[%s] syntax error", line)
				return nil, err
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				continue
			}
//获取节名
			res[sectionName] = make(map[string]interface{})
			currentSection = sectionName
		} else {
//其余为=分隔的k:v
			if !strings.Contains(line, "=") || strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
				err = fmt.Errorf("[%s] is syntax error", line)
				return nil, err
			}
			didx := strings.Index(line, "=")
			key := strings.TrimSpace(line[0:didx])
			value := strings.TrimSpace(line[didx+1:])
//获取k:v
			res[currentSection][key] = value
		}
	}

	return res, err
}


type ParentMap struct {
	ChildMap map[string]interface{}
}

func Getres(res map[string]map[string]interface{}) []ParentMap {
	parentMaps := []ParentMap{}

	for _, childMap := range res {
		parentMap := ParentMap{
			ChildMap: childMap,
		}
		parentMaps = append(parentMaps, parentMap)
	}

	return parentMaps
}