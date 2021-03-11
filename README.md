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

## Function

```go
// 기본 형태
func add(a int, b int) int {
  return a + b
}

// a, b가 같은 형태이면 아래와 같이 선언 가능
func add(a, b int) int {
  return a + b
}

// go는 multiple return value를 지원함.
func lenAndUpper(name string) (int, string) {
  return len(name), strings.ToUpper(name)
}

func main() {
  totalLength, upperName := lenAndUpper("kyuha")
  totalLength, _ := lenAndUpper("kyuha") // _를 쓰면 해당 값은 무시됨.
}

// 가변인자 : ...
// TODO : 여기는 다시 한 번 찾아볼 필요가 있어보임!!
func repeatMe(words ...string) {
  fmt.Println(words)
}

func main() {
  repeatMe("hwang", "kyu", "ha")
}

// naked return
func lenAndUpper(name string) (length int, uppercase string) { // 리턴내용 작성
  length = len(name) // :=이 아니라 =을 쓰는 이유는, 이미 length가 int로 선언되기 때문임.
  uppercase = strings.ToUpper(name)
  return // return을 따로 쓸 필요 없다.
}

func main() {
  totalLength, upperName := lenAndUpper("kyuha")
}

// defer : 함수가 끝나고 무언가를 할 수 있는 기능
// 파일을 닫거나 하는 등에 사용될 수 있음.
func lenAndUpper(name string) (length int, uppercase string) {
  defer fmt.Println("I'm done") // return 이후에 이 부분이 실행 됨.
  length = len(name)
  uppercase = strings.ToUpper(name)
  return
}
```

## for, range, ...args

- for : go에는 반복이 for 만 있음.

```go
func superAdd(numbers ...int) int {
  total := 0

  fmt.Println(numbers)

  for index, number := range numbers {
    fmt.Println(index, number)
  }

  for i:=0; i< len(numbers) ; i++ {
    fmt.Println(numbers[i])
  }

  for _, number := range numbers {
    fmt.Println(number)
  }
  
  return total
}

func main() {
  total := superAdd(1, 2, 3, 4, 5, 6)
  fmt.Println(total)
}
```

## If with a Twist

```go
func canIDrink(age int) bool {
  if age < 18 {
    return false
  } else {
    return true
  }
}

// variable expression
func canIDrink(age int) bool {
  // if 조건문 전에 ;를 쓰면 왼쪽에 변수를 선언할 수 있음.
  // 이것을 사용하면 koreanAge가 if문에서만 쓰인다는 것을 명시할 수 있다.
  if koreanAge := age; koreanAge < 18 {
    return false
  } else {
    return true
  }
}
```

## switch

```go
func canIDrink(age int) bool {
  switch age {
  case age < 18:
    return false
  case age == 18:
    return true
  case age > 50:
    return false
  }
  return false
}

func canIDrink(age int) bool {
  switch koreanAge := age + 2; koreanAge { // variable expression도 가능
  case age < 18:
    return false
  case age == 18:
    return true
  }
  return false
}
```

## Pointers!

```go
func main() {
  a := 2
  b := &a // &를 붙이면 포인터임
  fmt.Println(*b) // 포인터에 *을 붙이면 해당 주소 참조
}
```
