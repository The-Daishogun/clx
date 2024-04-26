#!/bin/bash
# Check if running with root privileges
if [[ $EUID -ne 0 ]]; then
  echo "This script requires root privileges. Please run with sudo."
  exit 1
fi

# Define the URL to download from
url="https://api.github.com/repos/The-Daishogun/clx/releases/latest"

# Get the OS and architecture
os=$(uname -s | tr '[:upper:]' '[:lower:]')
arch=$(uname -m)

# Adjust the architecture name to match the format in the URL
if [[ "$arch" == "x86_64" ]]; then
  arch="amd64"
elif [[ "$arch" == "aarch64" ]]; then
  arch="arm64"
fi
# Get the download link for the latest release
download_url=$(curl -s "$url" | grep "browser_download_url.*$os-$arch.*\.tar\.gz[^.]" | cut -d '"' -f 4)

# Check if download link was found
if [[ -z "$download_url" ]]; then
  echo "Failed to find download link for clx for $os-$arch"
  exit 1
fi

# Define the installation directory (you can change this to your preference)
install_dir="/usr/local/bin"

# Make sure the installation directory exists
sudo mkdir -p "$install_dir"

# Download the archive and extract it
curl -L "$download_url" --output /tmp/clx.tar.gz 

sudo tar xf /tmp/clx.tar.gz --directory=$install_dir

echo "Installation completed successfully."