
const sentiment = require('multilang-sentiment');
const fs = require('fs');

const content = 'los gatos son tontos';
//console.dir(sentiment(content, 'es'));

const path = "../data/processed/processed_news_suicide_data/resultsSuicide.json";
const pathout = "../data/processed/processed_news_suicide_data/resultsSuicideSentiment.json";

const newsResult = require(path);

for (const new_item of newsResult.news_scraped_result){
    const sentimentAnalysis = sentiment(new_item.content, 'es')
    new_item.sentiment_analysis = sentimentAnalysis;
    console.log(sentimentAnalysis);
}

fs.writeFileSync(pathout, JSON.stringify(newsResult));




