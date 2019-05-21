setwd("/home/hugojose.bello/Documentos/git_repos/news-scrapers/data-tunning-scripts/plot-time-series-data")
ine_data <- read.csv("procesado_suicidio_edad_ambos_hasta_2010.csv",  header=TRUE, sep=";")
scraping_data <- read.csv("scraped_news_suicide_data_sorted_2017_2010.csv",  header=TRUE, sep=";")
scraping_data_filtered <- read.csv("scraped_news_suicide_data_sorted_2017_2010_filtrado.csv",  header=TRUE, sep=";")
all_data <- read.csv("all_2011-2017.csv",  header=TRUE, sep=";")
# Installation
# install.packages('ggplot2')
# Loading
library(ggplot2)
#ine_data$month_code<- as.Date(ine_data$month_code , "%Y/%m/%d")
#install.packages('lubridate')
library(lubridate)
ine_data$month_code <- ymd(ine_data$month_code)
scraping_data$month_code <- ymd(scraping_data$month_code)
scraping_data_filtered$month_code <- ymd(scraping_data_filtered$month_code)
  
ggplot(ine_data, aes(month_code, todas_edades)) + geom_line() + xlab("") + ylab("Suicidios")
ggplot(ine_data, aes(month_code, de_15_a_29 + menores_15)) + geom_line() + xlab("") + ylab("Suicidios menores")

ggplot(scraping_data, aes(month_code, number_of_suicide_news)) + geom_line() + xlab("") + ylab("Numero noticias suicidio")
ggplot(scraping_data_filtered, aes(month_code, number_of_suicide_news)) + geom_line() + xlab("") + ylab("Numero noticias suicidio")


plot(all_data$noticias_suicidio, all_data$suicidios, col = "blue", main = "", xlab = "noticias suicidio", ylab = "suicidios")
lm(all_data$noticias_suicidio ~ all_data$suicidios)
abline(lm(all_data$noticias_suicidio ~ all_data$suicidios))

cor(all_data$noticias_suicidio, all_data$suicidios)
  