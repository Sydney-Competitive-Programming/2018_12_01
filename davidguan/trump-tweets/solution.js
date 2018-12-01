// Case-insensitive, which treate "Hello", "heLLO" the same.
// `node solution.js ../../trump_tweets_18.json` gives:
// The top three used words are:
// "the": 4556 times
// "and": 2691 times
// "to": 2456 times
var fs = require('fs')

var obj = JSON.parse(fs.readFileSync(process.argv[2]))
var words = obj.reduce((words, d) => {
  d.text.split(' ').forEach(w => {
    w = w.toLowerCase()
    words[w] = words[w] ? words[w] + 1 : 1
  })
  return words
}, {})

// Don't have to use arr if we just need the most used word, use arr for fun.
let wordArr = []
Object.keys(words).forEach(word => {
  wordArr.push({ w: word, c: words[word] })
})
wordArr.sort((a, b) => b.c - a.c)
console.log("The top three used words are:")
wordArr.slice(0, 3).forEach(d => {
  console.log(`"${d.w}": ${d.c} times`)
})