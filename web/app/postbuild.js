const fs = require('fs-extra');
const path = require('path');

const buildPath = path.join(__dirname, 'build');
const gitKeepPath = path.join(buildPath, '.gitkeep');
fs.closeSync(fs.openSync(gitKeepPath, 'w'))

const publicPath = path.join(__dirname, '..', 'static');
fs.move(buildPath, publicPath, { overwrite: true });
