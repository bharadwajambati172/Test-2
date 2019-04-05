
const fs = require('fs');
const util = require('util');

const readFile = util.promisify(fs.readFile);
const writeFile = util.promisify(fs.writeFile);

let aes256 = require('aes256');
let key = 'myKey';
let cipher = aes256.createCipher(key);

async function writeEncryptedData() {
    try {
        let data = await readFile("./input.txt");
        var encryptedData = cipher.encrypt(data.toString());
        console.log("encryptedData :",encryptedData)
        await writeFile("./output.txt", encryptedData)
    } catch (err) {
        console.log(err)
    }
}

async function readEncryptedData() {
    try {
        let data = await readFile("./output.txt");
        var decryptedData = cipher.decrypt(data.toString());
        console.log("decryptedData :",decryptedData)
    } catch (err) {
        console.log(err)
    }
}

writeEncryptedData()
readEncryptedData()