#!/bin/bash

# This script finds all gong model directories and runs gong generate conditionally.

# Exit immediately if a command exits with a non-zero status.
set -e

echo "🚀 Starting conditional gong build..."

# Corrected line: We apply dirname a third time to get the actual project root.
find . -path '*/go/models/docs.go' -exec dirname {} \; | xargs -I {} dirname {} | xargs -I {} dirname {} | sort -u | while IFS= read -r dir; do
  (
    # Change into the target directory in a subshell
    cd "$dir"
    
    if [ -f "embed_ng_dist_ng.go" ]; then
      echo "==> Found embed_ng_dist_ng.go in '$dir', running standard command..."
      gong generate --skipGoModCommands --skipNpmWorkspaces --skipNg --skipGoModCommands go/models
    else
      echo "==> Did not find embed_ng_dist_ng.go in '$dir', running level1 command..."
      gong generate --level1 --skipNpmWorkspaces --skipGoModCommands go/models
    fi
  )
done

echo "✅ Conditional build finished."