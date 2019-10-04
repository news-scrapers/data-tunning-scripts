import os
import pandas as pd
import json

sentiment = pd.read_csv("../data/processed/processed_news_suicide_data/scraped_news_suicide_data_sentiment_sorted.csv",sep=";",error_bad_lines=False, encoding="utf-8")
suicide_ine = pd.read_csv("../data/processed/processed_ine_data/procesado_suicidio_edad_ambos.csv",sep=";",error_bad_lines=False, encoding="utf-8")


sentiment["average_sentiment_result"] = sentiment["average_sentiment_result"]/sentiment["number_of_suicide_news"]

sentiment_clean = sentiment[["month_code", "number_of_suicide_news", "average_sentiment_result"]]
suicide_ine_clean = suicide_ine[["mes", "descrip", "Todas las edades"]]
result = sentiment_clean.merge(suicide_ine_clean, left_on='month_code', right_on='mes',left_index=True, how='left')

del result["mes"]

result = result.rename(index=str, columns={"Todas las edades": "suicidios"})

result = result.rename(index=str, columns={"number_of_suicide_news": "noticias_suicidio"})


result = result.rename(index=str, columns={"average_sentiment_result": "media_sentimiento_noticias"})


result.to_csv("../data/processed/all.csv", sep=";", index=False)

print(result)