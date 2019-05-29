
const Sentiment = require('sentiment');
const sentiment = new Sentiment();


const analyze = (content) => {
    const result = sentiment.analyze(content);
    console.dir(result);
    return result;
}


const content = 'cats are stupid';
const newsResult = require("./resultsSuicide.json");
analyze(content)
