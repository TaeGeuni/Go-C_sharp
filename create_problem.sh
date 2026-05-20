#!/bin/bash

# 인자 개수 확인
if [ "$#" -ne 2 ]; then
    echo "💡 사용법: ./create_problem.sh <분류명> <문제명>"
    echo "💡 예시: ./create_problem.sh hash-table 0001-two-sum"
    exit 1
fi

CATEGORY=$1
PROBLEM=$2
TARGET_DIR="$CATEGORY/$PROBLEM"

# 1. 분류(카테고리) 폴더가 없으면 생성
mkdir -p "$CATEGORY"

# 2. 문제 폴더가 이미 존재하는지 체크
if [ -d "$TARGET_DIR" ]; then
    echo "❌ 에러: 이미 존재하는 문제 폴더입니다 ($TARGET_DIR)"
    exit 1
fi

echo "🚀 [$PROBLEM] 문제 환경 세팅을 시작합니다..."

# 3. C# 콘솔 프로젝트 생성 (.csproj 및 Program.cs 자동 생성)
# C#은 폴더명에 숫자나 하이픈이 들어가도 프로젝트 생성에 문제가 없습니다.
cd "$CATEGORY" || exit
dotnet new console -n "$PROBLEM" -o "$PROBLEM" --quiet
cd "$PROBLEM" || exit

# C# 기본 템플릿 덮어쓰기 (최상위 문 사용)
cat <<EOF > Program.cs
// $PROBLEM
using System;

Console.WriteLine("Hello, $PROBLEM!");
EOF

# 4. Go 기본 템플릿 생성
cat <<EOF > main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, $PROBLEM!")
}
EOF

# 5. 마크다운 노트 템플릿 생성
cat <<EOF > note.md
# $PROBLEM

## 💡 문제 링크
* [LeetCode - 문제 이름](URL)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** $CATEGORY
* **시간 복잡도:** $O()$
* **공간 복잡도:** $O()$

### 핵심 아이디어
* 

## 🏃‍♂️ 코드 실행
* **Go:** \`go run main.go\`
* **C#:** \`dotnet run\` (현재 디렉토리에서)
EOF

echo "✅ 세팅 완료! ($TARGET_DIR)"