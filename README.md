# Go 언어 학습

## 개발환경 설정

- 설치
  - 다운로드 : https://golang.org/
  - 정상 설치 확인
    - 폴더 생성 확인 : /usr/local/go (mac OS) 
    - 커맨드 확인 : go 명령어 입력

**Go PATH : /usr/local/go는 go 언어에서 중요함**

- Go 에서는 무조건 코드가 Go PATH 디렉토리에 있어야 함!
  - go는 코드를 받는 곳이 정해져 있음 (npm, pip와는 다름)
  - 키포인트 : 모든 것을 디렉토리에 보관한다!
  - 코드를 잘 정리하는 방법은 도메인별로 분류해서 저장하는 것임!
  - go modules라는 것이 있는데.. 이건 나중에 살펴봄

```sh
cd /usr/local/go
mkdir src
cd src
mkdir github.com
mkdir comdori-web
cd comdori-web
git clone git@github.com:comdori-web/learngo.git
cd learngo
touch main.go
```

- main.go 가 있는 폴더를 vscode로 열어서 vscode에서 설치하라는 것들을 모두 설치한다.

## Main Package

- 프로젝트를 컴파일 하길 원한다면 "main.go" 패키지이름을 사용해야만 함.
  - 컴파일러가 main 이름을 가진 패키지를 찾고, main function을 찾아 실행시킴.
- go에서는 어떤 패키지를 사용하는지 작성해주어야 함.

```go
package main

import "fmt"

func main() {
  fmt.Pringln("Hello World!");
}
```

- 실행 : go run main.go

## Packages and Imports

- 위의 예제에서 fmt는 go가 가지고 있는 패키지 중 하나임
- vscode에서는 저장시 패키지를 자동으로 import하거나 해제해줌
- fmt
  - formmatting을 위한 package
  - 키포인트 : 왜 Println의 'P'는 대문자인가?
    - func를 export하려면 대문자로 시작하면 됨. (public/private 개념)

**cmd + click 하면 해당 패키지 내용을 볼 수 있음.**

## Variables and Constants

```go
const name string = "kyuha" // 상수 : 변경 안됨.
var name string = "kyuha" // 변수 : 변경 가능
name := "kyuha" // 축약형, 함수 안에서만 다음과 같이 축약하여 선언 가능, 타입은 go가 찾아줌.
```

- go는 type 언어이다.
