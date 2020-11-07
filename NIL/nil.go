// nil study

package main

import (
	"fmt"
)

func main() {

	/*
		tmpVarA, tmpVarB 에 각각 "",0을 직접 할당해줌.
		tmpVarA, tmpVarB 는 로그를 찍어보면 "", 0 이 나오겠지.
		tmpVarC, tmpVarD 는 type만 선언 해놓음.
		tmpVarC, tmpVarD 는 로그를 찍으면 nil 이라고 나올까 ?
			result >
					 tmpVarA is
					 tmpVarB is 0
					 tmpVarC is
					 tmpVarD is 0

		놉! tmpVarC 는 "" , tmpVarD는 0으로 나온다. 타입을 명시해줬기 때문인듯.
		그렇다면 tmpVarA == tmpVarC , tmpVarB == tmpVarD 인가? 아닐듯?
			result >
					 is same ? tmpVarA,tmpVarC  true
					 is same ? tmpVarB,tmpVarD  true

		놉! 각각 모두 true를 뱉는다. 음.. tmpVarC,tmpVarD == nil 은 true인가..? 그럴듯.
		노옵. 애초에 컴파일 에러가 생긴다. int, string 타입엔 자동으로 값을 할당하는것 같다.
		그러면 배열로 바꿔보자.
			result >
					tmpVarC is []
					tmpVarD is []
					tmpVarC,tmpVarD are nil ?   true true

		결론 : nil 이란 명시적인 초기값이 없는 경우라 보면 되곘다.



	*/

	tmpVarA := [3]string{"a", "b", "c"}
	tmpVarB := [3]int{1, 2, 3}
	var tmpVarC []string
	var tmpVarD []int

	fmt.Println(`tmpVarA is`, tmpVarA)
	fmt.Println(`tmpVarB is`, tmpVarB)
	fmt.Println(`tmpVarC is`, tmpVarC)
	fmt.Println(`tmpVarD is`, tmpVarD)

	fmt.Println(`tmpVarC,tmpVarD are nil ?  `, tmpVarC == nil, tmpVarD == nil)

}
