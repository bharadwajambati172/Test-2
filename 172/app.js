
const fs = require('fs');
const util = require('util');

const crypto = require("crypto");

const readFile = util.promisify(fs.readFile);
const writeFile = util.promisify(fs.writeFile);


async function readAndWriteHashData() {
    try {
        let data = await readFile("./input.txt");
        let hash = crypto.createHash('sha256').update(data).digest('base64');
        await writeFile("./output.txt", hash)
    } catch (err) {
        console.log(err)
    }
}
readAndWriteHashData()
