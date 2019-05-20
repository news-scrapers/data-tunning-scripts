setwd("/home/hugojose.bello/Documentos/git_repos/news-scrapers/data-tunning-scripts/plot-time-series-data")
ine_data <- read.csv("procesado_suicidio_edad_ambos_hasta_2010.csv",  header=TRUE, sep=";")
scraping_data <- read.csv("scraped_news_suicide_data_sorted_2017_2010.csv",  header=TRUE, sep=";")
# Installation
# install.packages('ggplot2')
# Loading
library(ggplot2)
#ine_data$month_code<- as.Date(ine_data$month_code , "%Y/%m/%d")
#install.packages('lubridate')
library(lubridate)
ine_data$month_code <- ymd(ine_data$month_code)
scraping_data$month_code <- ymd(scraping_data$month_code)

ggplot(ine_data, aes(month_code, todas_edades)) + geom_line() + xlab("") + ylab("Suicidios")
ggplot(scraping_data, aes(month_code, number_of_suicide_news)) + geom_line() + xlab("") + ylab("Numero noticias suicidio")
