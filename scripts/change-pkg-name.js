import { promises as fs } from 'fs';
import path from 'path';

async function start() {
  let myArgs = process.argv.slice(2);
  let filePath = myArgs[0];
  let buffer = await fs.readFile(path.resolve(filePath));
  let pkg = JSON.parse(buffer.toString());
  pkg.name += '.test';
  await fs.writeFile(filePath, JSON.stringify(pkg, null, 2));
  console.info(`[node] updated ${filePath}`);
}

start();