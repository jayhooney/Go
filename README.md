# Go lang toy project

```
고언어 연습용 토이 프로젝트 레파지토리
```

# Concept

## 1. nil ?

> 참고 1. <https://2kindsofcs.tistory.com/3>  
> 참고 2. <https://stackoverflow.com/questions/35983118/what-does-nil-mean-in-golang>

## 2. main ?

Main 패키지
일반적으로 패키지는 라이브러리로서 사용되지만, "main" 이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식된다. 패키지명이 main 인 경우, 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다. 그리고 이 main 패키지 안의 main() 함수가 프로그램의 시작점 즉 Entry Point가 된다. 패키지를 공유 라이브러리로 만들 때에는, main 패키지나 main 함수를 사용해서는 안된다.

> 참고 1. <http://golang.site/go/article/15-Go-%ED%8C%A8%ED%82%A4%EC%A7%80>

## 3. 자동완성 ?

아래 참고 1 을 따라 VSCode를 세팅해주면 자동완성 기능을 편하게 쓸 수 있다.
근데.. 내가 만든 패키지에 선언된 함수는 왜 자동완성이 안될까..

> 참고 1. <https://medium.com/backend-habit/setting-golang-plugin-on-vscode-for-autocomplete-and-auto-import-30bf5c58138a>
