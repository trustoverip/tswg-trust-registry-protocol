#!/bin/bash
# Renders both the Approved and Working Draft specs sequentially.
# SpecUp only reads specs.json and exits after rendering one spec,
# so we swap configs between renders.

set -e

# Run image/asset setup
bash ./move_images.sh

# Render approved spec (already in specs.json)
echo "=== Rendering Approved spec ==="
node -e "require('spec-up')({ nowatch: true })"

# Swap in draft config, render, then restore
echo "=== Rendering Working Draft spec ==="
cp specs.json specs-approved.json
cp specs-draft.json specs.json
node -e "require('spec-up')({ nowatch: true })"
cp specs-approved.json specs.json
rm specs-approved.json

# Create root redirect to approved version
cat > ./dist/index.html << 'REDIRECT'
<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="refresh" content="0; url=approved/">
  <title>Redirecting...</title>
</head>
<body>
  <p>Redirecting to <a href="approved/">approved specification</a>...</p>
</body>
</html>
REDIRECT

echo "=== Done. Output: dist/approved/index.html and dist/draft/index.html ==="
