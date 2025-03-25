#!/bin/bash

USER=$(whoami)
GROUP=$(id -gn "$USER")
SERVICE_NAME="gocloud"
BUILD_DIR="$(pwd)/bin"
BINARY="${SERVICE_NAME}_$(go env GOOS)_$(go env GOARCH)"
SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"
WORKING_DIR="/opt/yet-another-cloud/gocloud"
EXEC_PATH="${WORKING_DIR}"

if ! command -v go &>/dev/null; then
  echo "You need to install Golang v1.21.1+: https://go.dev/doc/install"
  exit 0
fi

if systemctl is-active --quiet "${SERVICE_NAME}"; then
  echo "Stopping the service"
  sudo systemctl stop "${SERVICE_NAME}"

  if ! systemctl is-active --quiet "${SERVICE_NAME}"; then
    echo "Stopped successfully"
  else
    echo "Failed to stop the service"
    exit 1
  fi
fi

# Rebuild every time
echo "Building binaries"
./build.sh "$(go env GOOS)/$(go env GOARCH)"

echo "Copying '${BUILD_DIR}/${BINARY}' into '${EXEC_PATH}'"
sudo mkdir -p "${EXEC_PATH}"
sudo chown -R "${USER}:${GROUP}" "${EXEC_PATH}"
sudo cp "${BUILD_DIR}/${BINARY}" "${EXEC_PATH}"

echo "[Unit]
Description=gocloud server
After=network.target

[Service]
ExecStart=${EXEC_PATH}/${BINARY}
WorkingDirectory=${WORKING_DIR}
Restart=always
User=${USER}
Group=${GROUP}
Environment=PORT=8080
ExecReload=/bin/kill -HUP \$MAINPID

[Install]
WantedBy=multi-user.target" | sudo tee "$SERVICE_FILE" >/dev/null

echo "Enabling and starting the service"
sudo systemctl enable "$SERVICE_NAME"
sudo systemctl start "$SERVICE_NAME"

if systemctl is-active --quiet "${SERVICE_NAME}"; then
  echo "Successfully started the service"
else
  echo "Failed to start the service"
fi

echo "Installed and started as '${SERVICE_NAME}'"
