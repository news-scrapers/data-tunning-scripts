# https://www.pingshiuanchua.com/blog/post/simple-sentiment-analysis-python?utm_campaign=News&utm_medium=Community&utm_source=DataCamp.com
import nltk
import json
from textblob import TextBlob

nltk.downloader.download('vader_lexicon')

#nltk.download('vader_lexicon')
#nltk.download('perluniprops')

def analyze_new(newContent):
    result = nltk_sentiment(newContent)
    return result

def nltk_sentiment(sentence):
    from nltk.sentiment.vader import SentimentIntensityAnalyzer

    nltk_sentiment = SentimentIntensityAnalyzer()
    translated = TextBlob(sentence).translate(to="en")
    score = nltk_sentiment.polarity_scores(str(translated))
    return score

def load_news():
    with open('../data/processed/processed_news_suicide_data/resultsSuicide.json') as json_file:  
        return json.load(json_file)

def write_analyzed_news(news):
    with open('../data/processed/processed_news_suicide_data/resultsSuicideWithAnalysis.json') as json_file:  
        son.dump(news, f)

if __name__== "__main__":
    print(analyze_new("que bueno es todo me encanta si si"))
    news = load_news()
    amount = len(news["news_scraped_result"])
    print("analyzing " + str(amount))
    for i in range(0, amount):
        analysis = analyze_new(news["news_scraped_result"][i]["content"])
        news["news_scraped_result"][i]["sentiment_analysis"] = analysis
        print(analysis)
    write_analyzed_news(news)