// tar -cJf resultsSuicideSentiment.tar.xz resultsSuicide
// 7z a -mm=Deflate -mfb=258 -mpass=15 -r foo.zip resultsSuicide
const sentiment = require('multilang-sentiment');
const fs = require('fs');

const content = 'los gatos son tontos';
//console.dir(sentiment(content, 'es'));

const path = "../data/processed/processed_news_suicide_data/resultsSuicide";
const pathout = "../data/processed/processed_news_suicide_data/resultsSuicideSentiment/";


    //const sentimentAnalysis = sentiment(new_item.content, 'es')
if (!fs.existsSync(pathout)){
    fs.mkdirSync(pathout);
}

const startingPoint = "2007-07-03T14:20:12.798Z"

fs.readdir(path, function(err, list) {
    if (err) return done(err);
    list.forEach(element => {
        const filecontents = require(path + "/" + element);
        if (element.indexOf("sent_")===-1){
            sentimentAnalysis = sentiment(filecontents.content, 'es');
            filecontents.sentiment_analysis = sentimentAnalysis;
            fs.writeFileSync(path+"/"  + element, JSON.stringify(filecontents));    
            fs.renameSync( path+"/"  + element, path+"/sent_"  + element)
            console.log(element)
        }
        
        
    });

}
)





