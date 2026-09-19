const fs = require('fs');
const path = require('path');
const esbuild = require('esbuild');

const browserDir = path.join(__dirname, 'dist', 'ng-github.com-fullstack-lang-gong-lib-splitlite', 'browser');

if (!fs.existsSync(browserDir)) {
  console.log('browser directory does not exist, skipping bundling.');
  process.exit(0);
}

const files = fs.readdirSync(browserDir);
const mainFile = files.find(f => f.startsWith('main-') && f.endsWith('.js') && !f.endsWith('.bundled.js'));

if (!mainFile) {
  console.log('No main-*.js file found to bundle.');
  process.exit(0);
}

const chunkFiles = files.filter(f => f.startsWith('chunk-') && f.endsWith('.js'));
if (chunkFiles.length === 0) {
  console.log('No chunk-*.js files found, nothing to bundle.');
  process.exit(0);
}

console.log(`📦 Bundling ${chunkFiles.length} chunks into ${mainFile}...`);

const mainPath = path.join(browserDir, mainFile);
const tempOutPath = path.join(browserDir, `${mainFile}.bundled.js`);

esbuild.buildSync({
  entryPoints: [mainPath],
  bundle: true,
  format: 'esm',
  outfile: tempOutPath,
  allowOverwrite: true,
  logLevel: 'warning',
});

// Overwrite original main-*.js with bundled version
fs.renameSync(tempOutPath, mainPath);

// Delete the merged chunks
for (const chunk of chunkFiles) {
  fs.unlinkSync(path.join(browserDir, chunk));
}

console.log('✅ Bundling complete, chunks removed.');
