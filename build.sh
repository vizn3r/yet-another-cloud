#!/bin/bash

TARGETS=(
  "windows/amd64"
  "windows/amd64"
  "windows/386"
  "windows/arm"
  "windows/arm64"
  "linux/amd64"
  "linux/386"
  "linux/arm"
  "linux/arm64"
)

DIRS=(
  "cloud"
)

BUILD_DIR="bin"
BASE_DIR="$(pwd)"
mkdir -p "$BUILD_DIR"

build() {
  local OS="$1"
  local ARCH="$2"

  for DIR in "${DIRS[@]}"; do
    MODULE_NAME=$(basename "$(pwd)")

    NAME="$BUILD_DIR/${MODULE_NAME}_${OS}_${ARCH}"
    if [ "$OS" == "windows" ]; then
      NAME+=".exe"
    fi

    echo "[BUILD] Building '$DIR' for $OS/$ARCH..."

    # Build cloud
    cd "$BASE_DIR/$DIR"
    env GOOS=$OS GOARCH=$ARCH go build -o "$BASE_DIR/$NAME" .

    if [ $? -ne 0 ]; then
      echo "[BUILD] Failed build for $OS/$ARCH"
    else
      echo "[BUILD] Built $OS/$ARCH"
    fi
  done
}

if [ "$1" == "all" ]; then
  for TARGET in "${TARGETS[@]}"; do
    OS=${TARGET%/*}
    ARCH=${TARGET#*/}

    build "$OS" "$ARCH"
  done
  exit 0

elif [ "$#" -gt 0 ]; then
  TARGETS=()
  for arg in "$@"; do
    TARGETS+=("$arg")
  done

  for TARGET in "${TARGETS[@]}"; do
    OS=${TARGET%/*}
    ARCH=${TARGET#*/}
    build "$OS" "$ARCH"
  done

else
  build "$(go env GOOS)" "$(go env GOARCH)"
fi

echo "[BUILD] Build done"
