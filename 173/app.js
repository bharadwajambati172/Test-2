let words = ["war", "eat","donut", "pizza", "raw", "ate", "apizz"]

function sortWord(word){
    let result=word.split("")
    result=result.sort()
    return result.join("")
}

let output=[]
for(let i in words){
        let values=words.filter((x)=>{
            if(sortWord(x)===sortWord(words[i])){
                return x
            }
        })
        output.push("("+values.toString()+")")
        for(let j in values){
            if(words[i] !== values[j]){
            let index=words.indexOf(values[j])
            words.splice(index,1)
            }
        }
}
output=output.toString()
console.log(output)