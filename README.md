# Go lang toy project

```
고언어 연습용 토이 프로젝트 레파지토리
```

# Concept

## nil ?

> 참고 1. <https://2kindsofcs.tistory.com/3>  
> 참고 2. <https://stackoverflow.com/questions/35983118/what-does-nil-mean-in-golang>

## main ?

Main 패키지
일반적으로 패키지는 라이브러리로서 사용되지만, "main" 이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식된다. 패키지명이 main 인 경우, 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다. 그리고 이 main 패키지 안의 main() 함수가 프로그램의 시작점 즉 Entry Point가 된다. 패키지를 공유 라이브러리로 만들 때에는, main 패키지나 main 함수를 사용해서는 안된다.

> 참고 1. <http://golang.site/go/article/15-Go-%ED%8C%A8%ED%82%A4%EC%A7%80>

## 자동완성 ?

아래 참고 1 을 따라 VSCode를 세팅해주면 자동완성 기능을 편하게 쓸 수 있다.
근데.. 내가 만든 패키지에 선언된 함수는 왜 자동완성이 안될까..

> 참고 1. <https://medium.com/backend-habit/setting-golang-plugin-on-vscode-for-autocomplete-and-auto-import-30bf5c58138a>

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
