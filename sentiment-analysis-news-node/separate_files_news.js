
const sentiment = require('multilang-sentiment');
const fs = require('fs');

const content = 'los gatos son tontos';
//console.dir(sentiment(content, 'es'));

const path = "../data/processed/processed_news_suicide_data/resultsSuicide.json";
const pathout = "../data/processed/processed_news_suicide_data/resultsSuicide/";

const newsResult = require(path);

if (!fs.existsSync(pathout)){
    fs.mkdirSync(pathout);
}

for (const new_item of newsResult.news_scraped_result){
    //const sentimentAnalysis = sentiment(new_item.content, 'es')
    fs.writeFileSync(pathout+new_item.date +"__"+ new_item.id + ".json", JSON.stringify(new_item));
}





