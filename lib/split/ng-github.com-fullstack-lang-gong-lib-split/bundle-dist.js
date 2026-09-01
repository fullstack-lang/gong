const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const browserDir = path.join(__dirname, 'dist', 'ng-github.com-fullstack-lang-gong-lib-split', 'browser');

if (!fs.existsSync(browserDir)) {
  console.log('Browser dist directory does not exist:', browserDir);
  process.exit(0);
}

const files = fs.readdirSync(browserDir);
const mainFile = files.find(f => f.startsWith('main-') && f.endsWith('.js'));
const chunkFiles = files.filter(f => f.startsWith('chunk-') && f.endsWith('.js'));

if (mainFile && chunkFiles.length > 0) {
  console.log(`📦 Bundling ${chunkFiles.length} chunks into ${mainFile}...`);
  const mainPath = path.join(browserDir, mainFile);
  const bundledPath = mainPath + '.bundled.js';

  execSync(`npx -y esbuild "${mainPath}" --bundle --outfile="${bundledPath}" --format=esm`, { stdio: 'inherit' });
  fs.renameSync(bundledPath, mainPath);

  for (const cf of chunkFiles) {
    fs.unlinkSync(path.join(browserDir, cf));
  }

  const indexPath = path.join(browserDir, 'index.html');
  if (fs.existsSync(indexPath)) {
    let indexHtml = fs.readFileSync(indexPath, 'utf-8');
    indexHtml = indexHtml.replace(/<link[^>]*rel="modulepreload"[^>]*>/gi, '');
    fs.writeFileSync(indexPath, indexHtml, 'utf-8');
  }

  console.log('✅ Bundling complete, chunks removed.');
} else {
  console.log('No chunks found to bundle.');
}
