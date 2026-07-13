# 🚀 Algorithm Problem Solving

이 저장소는 LeetCode의 다양한 알고리즘 문제를 해결하고 코드를 모아두는 공간입니다. 
기존에는 Go와 C#을 사용했으나, 현재는 Go와 Python 두 가지 언어로 교차 검증하며 풀이하고 있습니다. 언어별 특성과 성능을 비교하고, 핵심 알고리즘 개념을 `note.md`에 정리하여 보관합니다.

## 🛠️ Tech Stack
![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Python](https://img.shields.io/badge/Python-3776AB?style=for-the-badge&logo=python&logoColor=white)
![C#](https://img.shields.io/badge/C%23-%23239120.svg?style=for-the-badge&logo=csharp&logoColor=white) *(Legacy)*

## 📁 Repository Structure
안정적인 개발 환경 유지를 위해 카테고리 및 문제별로 폴더를 관리합니다.

```text
AlgorithmProblem/
├── go.mod                     # Go 전체 모듈 루트
├── AlgorithmProblem.sln       # C# 전체 솔루션 루트 (과거 풀이용)
├── create_problem.sh          # 문제 자동 생성 스크립트
│
└── 📂 [algorithm-category]/       (예: hash-table)
    └── 📂 [problem-name]/         (예: 0001-two-sum)
        ├── main.py                # Python 코드
        ├── main.go                # Go 코드
        └── note.md                # 문제 접근법 및 복잡도 분석 노트
        (※ 이전 문제들의 경우 .csproj 및 Program.cs 가 존재할 수 있습니다.)

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

위 명령어를 실행하면 `hash-table/0001-two-sum` 폴더가 생성되며, 그 안에 `main.py`, `main.go` 파일과 문제 풀이를 기록할 마크다운 노트(`note.md`)가 즉시 세팅됩니다.
