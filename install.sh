#!/bin/bash

# Function to print instructions for adding ~/.local/bin to PATH
print_instructions() {
    echo "To have access to clx from the command line, you need to add ~/.local/bin to your PATH."
    echo "To add ~/.local/bin to your PATH, follow the instructions for your shell:"
    echo ""
    echo "Bash"
    echo 'echo "export PATH=\$PATH:\$HOME/.local/bin" >> ~/.bashrc'
    echo 'source ~/.bashrc'
    echo ""
    echo "Zsh:"
    echo 'echo "export PATH=\$PATH:\$HOME/.local/bin" >> ~/.zshrc'
    echo 'source ~/.zshrc'
    echo ""
    echo "Fish:"
    echo 'echo "set -Ux PATH \$PATH \$HOME/.local/bin" >> ~/.config/fish/config.fish'
}

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

# Define the installation directory
install_dir="$HOME/.local/bin"

# Make sure the installation directory exists
sudo mkdir -p "$install_dir"

# Download the archive and extract it
curl -L "$download_url" --output /tmp/clx.tar.gz 

sudo tar xf /tmp/clx.tar.gz --directory=$install_dir

echo "Installation completed successfully."
if [[ ":$PATH:" == *":$HOME/.local/bin:"* ]]; then
    echo "~/.local/bin is already in your PATH. you can use clx freely."
    echo "Usage: clx <prompt>"
else
    echo "~/.local/bin is not in your PATH."
    print_instructions
fi