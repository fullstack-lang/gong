
Port ot ng 20

same, when in a sub directory stack

```bash
mv ../../../package.json ../../../package.json.tmp
find ../../.. -type d -name "node_modules" -prune -exec rm -rf '{}' +
find ../../.. -type d -name "dist" -prune -exec rm -rf '{}' +
find ../../.. -type d -name ".angular" -prune -exec rm -rf '{}' +
find ../../.. -name "package-lock.json" -prune -exec rm -rf '{}' +
npm i
ng update @angular/core@20 @angular/cli@20 --allow-dirty
ng update @angular/material@20 --allow-dirty
ng update angular-split@20 --allow-dirty
ng update ngx-markdown@20 --allow-dirty
mv ../../../package.json.tmp ../../../package.json
find ../../.. -type d -name "node_modules" -prune -exec rm -rf '{}' +
find ../../.. -type d -name "dist" -prune -exec rm -rf '{}' +
find ../../.. -type d -name ".angular" -prune -exec rm -rf '{}' +
find ../../.. -name "package-lock.json" -prune -exec rm -rf '{}' +
npm i
ng build
```

To be performed from the lib/<lib name>/ng<path> directory
```bash
find ../../.. -type d -name "node_modules" -prune -exec rm -rf '{}' +
find ../../.. -type d -name "dist" -prune -exec rm -rf '{}' +
find ../../.. -type d -name ".angular" -prune -exec rm -rf '{}' +
find ../../.. -name "package-lock.json" -prune -exec rm -rf '{}' +
npm i
ng build
```

```bash
find .. -type d -name "node_modules" -prune -exec rm -rf '{}' +
find .. -type d -name "dist" -prune -exec rm -rf '{}' +
find .. -type d -name ".angular" -prune -exec rm -rf '{}' +
find .. -name "package-lock.json" -prune -exec rm -rf '{}' +
npm i
ng build
```

```bash
find . -type f -name "embed_ng_dist_ng.go" -exec dirname {} \; | sort -u | while IFS= read -r dir; do
  (cd "$dir" && gongc -skipGoModCommands -skipNpmWorkspaces go/models)
done
```

```bash
START_DIR=$(pwd)

# Change directory to the gong repository
cd ../../fullstack-lang/gong || { echo "Directory ../../fullstack-lang/gong not found."; exit 1; }

# Fetch the latest commit ID
COMMIT_ID=$(git rev-parse HEAD)
echo "Latest commit ID: $COMMIT_ID"

# Return to the starting directory
cd "$START_DIR" || exit

# Run go get with the commit ID
go get github.com/fullstack-lang/gong@"$COMMIT_ID"
```