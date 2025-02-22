# Go 개발 환경 설정

## 1. Go 설치 확인
```bash
$ go version
go version go1.23.5 darwin/arm64
```

## 2. 주요 환경 변수 확인
```bash
# GOPATH: Go 작업 공간 위치
$ echo $GOPATH
/Users/xdream/go

# GOROOT: Go 설치 위치
$ echo $GOROOT
/usr/local/go
```

## 3. Go 작업 공간 구조
Go 작업 공간은 다음과 같은 구조를 가집니다:
```
$GOPATH/
    ├── bin/    # 컴파일된 실행 파일
    ├── pkg/    # 컴파일된 패키지 파일
    └── src/    # 소스 코드
```

## 4. 모듈 지원 확인
Go 모듈은 1.11 버전부터 도입되었으며, 의존성 관리를 위해 사용됩니다.
```bash
# go.mod 파일 생성
$ go mod init example
```

## 5. IDE/편집기 설정
Go 개발을 위한 추천 IDE/편집기:
- VS Code + Go 확장
- GoLand
- Vim/Neovim + Go 플러그인

### VS Code Go 확장 기능
1. 코드 자동 완성
2. 코드 네비게이션
3. 리팩토링 도구
4. 디버깅 지원
5. 테스트 지원

## 6. 유용한 Go 도구들
- `go fmt`: 코드 포맷팅
- `go vet`: 코드 정적 분석
- `go test`: 테스트 실행
- `go doc`: 문서 확인
- `go get`: 패키지 다운로드
- `go install`: 패키지 설치

## 다음 단계
- [ ] Go 작업 공간 구조 확인
- [ ] go.mod 파일 생성
- [ ] IDE/편집기 설정
- [ ] Go 도구 사용법 익히기
