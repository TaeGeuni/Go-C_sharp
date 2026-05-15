# 🚀 Algorithm Problem Solving

이 저장소는 LeetCode의 다양한 알고리즘 문제를 Go와 Rust 두 가지 언어로 교차 검증하며 풀이한 코드를 모아두는 공간입니다. 언어별 특성과 성능을 비교하고, 핵심 알고리즘 개념을 `note.md`에 정리하여 보관합니다.

## 🛠️ Tech Stack
![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Rust](https://img.shields.io/badge/Rust-%23000000.svg?style=for-the-badge&logo=rust&logoColor=white)

## 📁 Repository Structure
안정적인 개발 환경(LSP, 자동완성) 유지를 위해 하이브리드 워크스페이스 구조를 사용합니다.

```text
AlgorithmProblem/
├── go.mod                 # Go 전체 모듈 루트
├── Cargo.toml             # Rust 전체 워크스페이스 루트
├── create_problem.sh      # 문제 자동 생성 스크립트
│
└── 📂 [algorithm-category]/   (예: hash-table)
    └── 📂 [problem-name]/     (예: 0001-two-sum)
        ├── Cargo.toml     # 독립적인 Rust 패키지 설정
        ├── 📂 src/
        │   └── main.rs    # Rust 코드
        ├── main.go        # Go 코드
        └── note.md        # 문제 접근법 및 복잡도 분석 노트