#!/bin/bash

# Define the base directories
SOURCE_BASE="./specification"
DEST_BASE="./dist"

# Find all 'images' directories under the source base
IMAGES_DIRS=$(find "$SOURCE_BASE" -type d -name "images")

for DIR in $IMAGES_DIRS; do
    # Compute the destination path by replacing the source base with the destination base
    DEST_PATH="${DIR/$SOURCE_BASE/$DEST_BASE}"

    # Ensure the destination directory exists
    mkdir -p "$DEST_PATH"

    # Move all files within the images directory to the destination
    cp -r "$DIR"/* "$DEST_PATH"
done

mkdir -p dist/v1/bindings/restful/
mkdir -p dist/images

cp specification/v1/bindings/restful/swagger.yaml dist/v1/bindings/restful/swagger.yaml
cp specification/v1/bindings/restful/swagger.yaml dist/swagger.yaml
cp -r specification/v1/images dist/
