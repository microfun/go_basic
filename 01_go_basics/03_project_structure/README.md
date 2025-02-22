# Go Project Structure

## 실습 목표
- Go 프로젝트의 기본 구조 이해
- 패키지와 모듈 시스템 사용법 학습

## 실습 내용

### 1. 프로젝트 구조 설정
- 프로젝트 루트에 `main.go` 파일과 `mypackage` 디렉토리를 생성
- `mypackage` 디렉토리에 `mypackage.go` 파일 작성

### 2. 모듈 초기화
- 프로젝트 루트에서 `go mod init project_structure` 명령어로 모듈 초기화
- `go.mod` 파일에서 모듈 이름 확인

### 3. 코드 작성 및 실행
- `main.go`에서 `project_structure/mypackage` 패키지를 임포트
- `mypackage`의 `Hello` 함수 호출
- `go run main.go` 명령어로 프로그램 실행

## 문제 해결
- 패키지 경로가 올바르지 않을 경우, `go.mod` 파일의 모듈 이름과 임포트 경로가 일치하는지 확인
- `go mod tidy` 명령어로 의존성 정리

이 실습을 통해 Go 프로젝트의 구조와 모듈 시스템을 이해하고, 패키지를 올바르게 사용하는 방법을 익혔습니다.
