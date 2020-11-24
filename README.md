# Go lang basic practice

```
고언어 기초 연습용 레파지토리
챕터별 진행내용은 아래 링크를 통해 진행
```

출처 : <https://mingrammer.com/gobyexample/>

# Concept

## nil ?

> 참고 1. <https://2kindsofcs.tistory.com/3>  
> 참고 2. <https://stackoverflow.com/questions/35983118/what-does-nil-mean-in-golang>

## main ?

Main 패키지
일반적으로 패키지는 라이브러리로서 사용되지만, "main" 이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식된다. 패키지명이 main 인 경우, 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다. 그리고 이 main 패키지 안의 main() 함수가 프로그램의 시작점 즉 Entry Point가 된다. 패키지를 공유 라이브러리로 만들 때에는, main 패키지나 main 함수를 사용해서는 안된다.

> 참고 1. <http://golang.site/go/article/15-Go-%ED%8C%A8%ED%82%A4%EC%A7%80>

## go run? build? install?

1. run

- 현재 작성한 go 파일을 테스트 하는 용도로 씀.

2. build

- go 파일을 실행명령 파일로 만들어주는 명령어.

3. install

- 직계 상위 디렉토리 코드 전체를 대상으로 한 명령파일을 생성하는 명령어

> 참고 1. <https://medium.com/@whj2013123218/go-run-build-install-fa29cab5bc32>

## Goroutine ?

goroutine은 thread와 비교하여 적은 비용이 든다.
goroutine은 몇 kb이며 goroutine은 thread에 비해 application의 요청에 따라 stack의 용량을 늘렸다 줄일 수 있다.
thread의 경우 stack사이즈가 고정되어 있다.
goroutine들은 몇개의 os thread에 다중송신한다.
몇 천개의 goroutine으로 실행되는 프로그램은 오로지 하나의 thread로 이루어져있다.
만약 goroutine이 입력때문에 blocking되어 있다면 다른 os Thread를 만들어서 남아있는 goroutine을 새로 만든 thread에 옮겨버린다.
goroutine들 사이에는 channel을 통해 의사소통한다.
channel은 의도적으로 goroutine을 사용하여 공유 메모리에 접근할 때 경쟁상태를 예방한다.

> 참고 1. <https://hoony-gunputer.tistory.com/entry/goroutine-and-channel>

## Channel ?

동시에 실행되고 있는 Goroutine을 연결해주는 파이프.
한 고루틴에서 채널에 값을 전달하면 다른 고루틴에서 이 값을 받을 수 있다.
버퍼를 지정하여 사용할 수도 있음!

## Atomic ?

원자성을 보장하는 연산(atomic operation)이란 쪼갤 수 없는 연산을 말한다.

i += 1 같은 단순한 연산이라 해도 최소 세 단계를 거친다.

1. 메모리에서 값을 읽는다.

2. 값을 1 증가시킨다.

3. 새로운 값을 메모리에 다시 쓴다.

고루틴 여러 개를 동시에 실행하면 CPU는 각 고루틴을 시분할하여 번갈아 실행하는 방식으로 병행 처리를 한다. 즉, i += 1 같은 단순한 연산을 처리하는 도중에도 CPU는 해당 연산을 잠시 중단한 후 다른 고루틴을 수행할 수 있고, 이 과정에서 동기화 문제가 발생할 수 있다.

> 참고 1. <https://thebook.io/006806/ch05/03/05/>

## Panic ?

panic은 일반적으로 무언가가 예상치 못하게 잘못되었음을 의미하고
대부분 정상적인 작동 중에는 발생해선 안되는 오류 또는 정상적으로 처리할 준비가 되어있지 않은 오류에 대해 빠르게 실패시키는데 사용.

에러를 핸들링 하기 위해 익셉션을 사용하는 타 언어와는 다르게 Go에서는 에러 가능성이 있는 곳에서 에러를 가리키는 반환값을 사용하는게 일반적.

## Collection Function ?

GO LANG에서는 제네릭을 지원하지 않는다.
따라서 특별히 필요한 경우엔 직접 만들어 사용해한다.
JS 의 프로토타입 함수을 선언해 사용하는것과 비슷한듯

# 기타 지식

## VSCode & 자동완성

아래 참고 1 을 따라 VSCode를 세팅해주면 자동완성 기능을 편하게 쓸 수 있다.
근데.. 내가 만든 패키지에 선언된 함수는 왜 자동완성이 안될까..

> 참고 1. <https://medium.com/backend-habit/setting-golang-plugin-on-vscode-for-autocomplete-and-auto-import-30bf5c58138a>
