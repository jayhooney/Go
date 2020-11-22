package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var logger = fmt.Println

type ResponseA struct {
	Page   int
	Fruits []string
}

// JSON 키 이름을 커스터마이징하려면 구조체 필드 선언부에 태그를 사용할 수 있다.
// 근데 이걸 언제쓰지..?
type ResponseB struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	logger("===== JSON =====")

	// bool encoding
	isB, _ := json.Marshal(true)
	logger(string(isB))

	// int encoding
	intB, _ := json.Marshal(1)
	logger(string(intB))

	// float encoding
	fltB, _ := json.Marshal(7.31)
	logger(string(fltB))

	// string encoding
	strB, _ := json.Marshal("jay")
	logger(string(strB))

	// slice encoding
	sliceD := []string{"apple", "google", "ms"}
	sliceB, _ := json.Marshal(sliceD)
	logger(string(sliceB))

	// map encoding
	mapD := map[string]int{"apple": 1, "google": 2, "ms": 3}
	mapB, _ := json.Marshal(mapD)
	logger(string(mapB))

	resAD := &ResponseA{
		Page:   1,
		Fruits: []string{"apple", "google", "ms"},
	}
	resAB, _ := json.Marshal(resAD)
	logger(string(resAB))

	resBD := &ResponseB{
		Page:   1,
		Fruits: []string{"apple", "google", "ms"},
	}
	resBB, _ := json.Marshal(resBD)
	logger(string(resBB))

	// decoding
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	logger(dat)

	// 디코딩된 맵을 사용하기 위해선 적절한 타입으로 캐스팅해야한다. (캐스팅 문법이 낯설군,,)
	num := dat["num"].(float64)
	logger(num)

	strs := dat["strs"].([]interface{})
	str2 := strs[1].(string)
	logger(str2)

	// JSON을 커스텀 데이터 타입으로 디코딩
	// 프로그램에 추가적인 타입 안전성을 추가할 수 있다는 장점이 있다. 데이터에 접근하고 디코딩할 때 타입 단언 `assertion` 을 할 필요성을 없애준다.
	jsonStr := `{"page":1, "fruits":["orange", "melon", "tomato"]}`
	res := ResponseB{}
	json.Unmarshal([]byte(jsonStr), &res)

	logger(res)
	logger(res.Page)
	logger(res.Fruits[1])

	// 위 예제들은 바이트와 문자열을 사용했지만, os.stdout과 같은 os.Writer나 http response body에 직접 json 인코딩을 스트리밍할 수 있다.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 731, "melon": 137}
	enc.Encode(d)

}
