# 🚀 Algorithm Problem Solving

이 저장소는 LeetCode의 다양한 알고리즘 문제를 Go와 C# 두 가지 언어로 교차 검증하며 풀이한 코드를 모아두는 공간입니다. 언어별 특성과 성능을 비교하고, 핵심 알고리즘 개념을 `note.md`에 정리하여 보관합니다.

## 🛠️ Tech Stack
![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![C#](https://img.shields.io/badge/C%23-%23239120.svg?style=for-the-badge&logo=csharp&logoColor=white)

## 📁 Repository Structure
안정적인 개발 환경(LSP, 자동완성) 유지를 위해 하이브리드 워크스페이스 구조를 사용합니다.

```text
AlgorithmProblem/
├── go.mod                     # Go 전체 모듈 루트
├── AlgorithmProblem.sln       # C# 전체 솔루션 루트 (선택 사항이나 통합 관리에 유리)
├── create_problem.sh          # 문제 자동 생성 스크립트
│
└── 📂 [algorithm-category]/       (예: hash-table)
    └── 📂 [problem-name]/         (예: 0001-two-sum)
        ├── [problem-name].csproj  # 독립적인 C# 프로젝트 설정 파일
        ├── Program.cs             # C# 코드 (또는 Solution.cs)
        ├── main.go                # Go 코드
        └── note.md                # 문제 접근법 및 복잡도 분석 노트
```
---

## ⚙️ 스크립트 사용법 (문제 세팅 자동화)
새로운 문제를 풀 때마다 일일이 폴더와 템플릿 파일을 만들 필요 없이 `create_problem.sh` 스크립트를 사용해 환경을 자동 구성합니다.

1. 스크립트 실행 권한 부여 (최초 1회만)
```bash
chmod +x create_problem.sh
```

2. 스크립트 실행
분류(카테고리) 이름과 문제 이름을 인자로 넘겨줍니다.
```bash
./create_problem.sh <분류명> <문제명>
```

3. 사용 예시
```bash
./create_problem.sh hash-table 0001-two-sum
```

위 명령어를 실행하면 `hash-table/0001-two-sum` 폴더가 생성되며, 그 안에 실행 가능한 C# 프로젝트(`.csproj`, `Program.cs`), Go 템플릿(`main.go`), 그리고 문제 풀이를 기록할 마크다운 노트(`note.md`)가 즉시 세팅됩니다.