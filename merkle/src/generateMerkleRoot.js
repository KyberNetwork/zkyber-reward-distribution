const path = require('path');
const fs = require('fs');
const argv = require('yargs/yargs')(process.argv.slice(2)).argv;
const { parseRewards } = require('./parseRewards');
const configPath = argv.f;

let json = JSON.parse(fs.readFileSync(path.join(__dirname, configPath), { encoding: 'utf8' }));
let cycle = json['cycle'];
if (typeof json !== 'object') {
  throw new Error('Invalid JSON');
}

json = JSON.stringify(parseRewards(json), null, 2);
fs.writeFileSync(path.join(__dirname, `../../results/cycle_${cycle}/merkle_data.json`), json);
console.log(`Writing merkle data to ./results/cycle_${cycle}/merkle_data.json ... \nDone.`);
fs.writeFileSync(path.join(__dirname, `../../results/latest_merkle_data.json`), json);
