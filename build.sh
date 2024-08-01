#!/bin/bash

# Set the output directory
OUTPUT_DIR="builds"

# Create the output directory if it doesn't exist
mkdir -p $OUTPUT_DIR

# Define the platforms
PLATFORMS=("linux/amd64" "darwin/amd64" "windows/amd64")

# Loop over each platform
for PLATFORM in "${PLATFORMS[@]}"; do
  OS=$(echo $PLATFORM | cut -d'/' -f1)
  ARCH=$(echo $PLATFORM | cut -d'/' -f2)
  
  OUTPUT_NAME="$OUTPUT_DIR/claybook-$OS-$ARCH"
  
  if [ $OS = "windows" ]; then
    OUTPUT_NAME+='.exe'
  fi

  echo "Building for $OS/$ARCH..."
  GOOS=$OS GOARCH=$ARCH go build -o $OUTPUT_NAME ./cmd/web/
  if [ $? -ne 0 ]; then
    echo "An error occurred while building for $OS/$ARCH."
    exit 1
  fi
done

echo "Build complete!"

