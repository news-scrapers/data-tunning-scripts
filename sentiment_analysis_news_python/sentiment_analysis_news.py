# https://www.pingshiuanchua.com/blog/post/simple-sentiment-analysis-python?utm_campaign=News&utm_medium=Community&utm_source=DataCamp.com
import nltk
import json
#nltk.download('vader_lexicon')
nltk.download('perluniprops')

def analyze_new(newContent):
    result = nltk_sentiment(newContent)
    return result

def nltk_sentiment(sentence):
    from nltk.sentiment.vader import SentimentIntensityAnalyzer

    nltk_sentiment = SentimentIntensityAnalyzer()
    score = nltk_sentiment.polarity_scores(sentence)
    return score

def load_news():
    with open('resultsSuicide.json') as json_file:  
        return json.load(json_file)["news_scraped_result"]

if __name__== "__main__":
    print(analyze_new("que bueno es todo me encanta si si"))
    news = load_news()
    print(analyze_new(news[0]["content"]))