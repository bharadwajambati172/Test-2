
const fs = require('fs');
const util = require('util');


const readFile = util.promisify(fs.readFile);
const writeFile = util.promisify(fs.writeFile);


async function readWriteFiles() {
    try{
  let data = await readFile("./input.txt");
  await writeFile("./output.txt",data)
    }catch(err){
        console.log(err)
    }
}

async function compareFiles(){
    try{
    let data1 = await readFile("./input.txt");
    let data2 = await readFile("./output.txt");
   return JSON.stringify(data1)==JSON.stringify(data2)
    }catch(err){
        console.log(err)
    }
}

readWriteFiles()
compareFiles().then((res)=>{
    console.log("Does both files have same data?",res)
}).catch((err)=>{console.log("error :",err)})