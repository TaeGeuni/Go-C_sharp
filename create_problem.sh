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

# 3. 문제 대상 폴더 생성 및 이동
mkdir -p "$TARGET_DIR"
cd "$TARGET_DIR" || exit

# 4. Python 기본 템플릿 생성
cat <<EOF > main.py
# $PROBLEM

def main():
    print("Hello, $PROBLEM!")

if __name__ == "__main__":
    main()
EOF

# 5. Go 기본 템플릿 생성
cat <<EOF > main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, $PROBLEM!")
}
EOF

# 6. 마크다운 노트 템플릿 생성 (실행 명령어 Python으로 변경)
cat <<EOF > note.md
# $PROBLEM

## 💡 문제 링크
* [LeetCode - 문제 이름](URL)

## 🧠 문제 접근 및 풀이
* **알고리즘 분류:** $CATEGORY
* **시간 복잡도:** O()
* **공간 복잡도:** O()

### 핵심 아이디어
* ## 🏃‍♂️ 코드 실행
* **Go:** \`go run main.go\`
* **Python:** \`python main.py\` (또는 \`python3 main.py\`)
EOF

echo "✅ 세팅 완료! ($TARGET_DIR)"