# 4장: 패키지 관리 (go mod)

## 실습 내용

1. **`go.mod` 파일 생성 및 이해**
   - `go mod init` 명령어를 사용하여 모듈을 초기화하고 `go.mod` 파일을 생성했습니다.

2. **의존성 관리**
   - `go get`을 사용하여 외부 패키지를 추가하고, `go mod tidy`로 불필요한 의존성을 정리했습니다.

3. **버전 관리**
   - `go get`을 통해 특정 버전으로 의존성을 업그레이드 및 다운그레이드하는 방법을 실습했습니다.

4. **`go.sum` 파일 탐색**
   - `go.sum` 파일의 목적과 역할을 이해하고, `go mod verify`로 의존성을 검증했습니다.

5. **의존성 교체**
   - `replace` 지시어를 사용하여 다른 버전이나 로컬 모듈로 의존성을 교체하는 방법을 실습했습니다.

6. **Semantic Import Versioning**
   - 주요 버전 변경을 처리하는 방법과 import 경로에 버전 번호가 포함된 모듈 생성 방법을 탐색했습니다.

7. **Vendor 디렉토리**
   - `go mod vendor`를 사용하여 `vendor` 디렉토리를 생성하고, 이를 활용하여 프로젝트를 빌드했습니다.

8. **빌드 제약 조건**
   - 빌드 제약 조건을 이해하고 다양한 환경에 대한 의존성 관리 방법을 배웠습니다.

## 버전 관리 실습

1. **패키지 업그레이드**
   - `go get github.com/fatih/color@latest` 명령어를 사용하여 `github.com/fatih/color` 패키지를 최신 버전으로 업그레이드했습니다.
   - 최신 버전에서는 빌드가 성공적으로 완료되었습니다.

2. **패키지 다운그레이드**
   - `go get github.com/fatih/color@v1.9.0` 명령어를 사용하여 `v1.9.0` 버전으로 다운그레이드 시도했습니다.
   - `v1.9.0` 버전에서는 빌드가 실패하였으며, 이는 API 변경이나 의존성 문제로 인한 것일 수 있습니다.

3. **결론**
   - 최신 버전으로 업그레이드하여 문제를 해결하였으며, 버전 간의 호환성을 유지하기 위해 변경 로그를 확인하거나 코드를 수정하는 것이 중요합니다.

## Vendor 디렉토리 사용 실습

1. **Vendor 디렉토리 생성**
   - `go mod vendor` 명령어를 사용하여 `vendor` 디렉토리를 생성하고, 모든 의존성을 로컬에 저장했습니다.

2. **로컬에서 의존성 관리**
   - `vendor` 디렉토리를 통해 외부 네트워크에 접근하지 않고도 프로젝트를 빌드하고 실행할 수 있음을 확인했습니다.

3. **복사 및 이동의 용이성**
   - `vendor` 디렉토리만 복사하여 다른 환경에서도 동일한 의존성을 사용할 수 있음을 실습했습니다.

4. **일관성 유지**
   - `vendor` 디렉토리를 통해 의존성의 버전을 고정하여, 다른 환경에서도 일관된 결과를 얻을 수 있음을 확인했습니다.

## 배운 점
- Go의 모듈 시스템을 통해 의존성을 효율적으로 관리할 수 있음을 배웠습니다.
- 다양한 명령어를 통해 모듈과 의존성을 관리하는 방법을 익혔습니다.

## 학습 정리
- `go.mod`와 `go.sum` 파일을 통한 의존성 관리 방법을 익혔습니다.
- `vendor` 디렉토리를 활용하여 프로젝트의 이동성과 일관성을 높이는 방법을 배웠습니다.
- 의존성 업그레이드/다운그레이드 및 교체 방법을 실습했습니다.
- 다양한 명령어를 통해 Go의 패키지 관리 시스템을 효율적으로 활용할 수 있게 되었습니다.

이 실습을 통해 Go의 패키지 관리에 대한 깊은 이해를 얻을 수 있었습니다.
