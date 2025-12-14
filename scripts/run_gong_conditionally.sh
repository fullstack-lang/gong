#!/bin/bash

# This script finds all gong model directories and runs gong generate conditionally.

# Source nvm to use the correct node version
if [ -s "$HOME/.nvm/nvm.sh" ]; then
  . "$HOME/.nvm/nvm.sh"
  nvm use
fi

# Exit immediately if a command exits with a non-zero status.
set -e

echo "🚀 Starting conditional gong build..."

# we perform a first pass to compile the level 4 applications (there is a dependency to the level 1 applications )

# Corrected line: We apply dirname a third time to get the actual project root.
find . -path '*/go/models/docs.go' -exec dirname {} \; | xargs -I {} dirname {} | xargs -I {} dirname {} | sort -u | while IFS= read -r dir; do
  (
    # Change into the target directory in a subshell
    cd "$dir"
    
    if [ -f "embed_ng_dist_ng.go" ]; then
      echo "==> Found embed_ng_dist_ng.go in '$dir', running standard level 4 compilation..."
      gong generate --skipGoModCommands --skipNpmWorkspaces --skipGoModCommands go/models
    fi
  )
done

## we compile the level 1 applications

# Corrected line: We apply dirname a third time to get the actual project root.
find . -path '*/go/models/docs.go' -exec dirname {} \; | xargs -I {} dirname {} | xargs -I {} dirname {} | sort -u | while IFS= read -r dir; do
  (
    # Change into the target directory in a subshell
    cd "$dir"
    
    if [ -f "embed_ng_dist_ng.go" ]; then
      echo "==> Did find embed_ng_dist_ng.go in '$dir', second pass, skipping level 4 compilation..."
    else
      echo "==> Did not find embed_ng_dist_ng.go in '$dir', running level1 command..."
      gong generate --level1 --skipGoModCommands go/models
    fi
  )
done

echo "✅ Conditional build finished."