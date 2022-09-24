package GoLib

import (
	"encoding/json"
	"io/ioutil"
)

/**
 *
 * @author yh
 * @date 2022/9/25 06:16
 * @description config：用于读取指定目录的配置的json信息，如果不指定目录则为当前目录，该信息只能为一层配置内容，返回一个map结构
 */

func ReadConfigInfo(path ...string) (error, map[string]string) {
	for _, e := range path {
		println(e)
	}
	if len(path) == 0 {
		res, err := ioutil.ReadFile(GetCurrentDirectory())

		if err != nil {
			return err, nil
		}
		var result = make(map[string]string)
		err = json.Unmarshal(res, &result)
		return nil, result
	} else {
		res, err := ioutil.ReadFile(path[0])

		if err != nil {
			return err, nil
		}
		var result = make(map[string]string)
		err = json.Unmarshal(res, &result)
		return nil, result
	}
}
