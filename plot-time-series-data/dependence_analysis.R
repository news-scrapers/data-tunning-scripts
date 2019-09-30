setwd("/home/hugo/Documents/hugo_documentos/github/data-tunning-scripts/plot-time-series-data")
ine_data <- read.csv("procesado_suicidio_edad_ambos_hasta_2007.csv",  header=TRUE, sep=";")
scraping_data_filtered <- read.csv("scraped_news_suicide_data_sorted_2017_2007_filtrado_no_outliers.csv",  header=TRUE, sep=";")
all_data <- read.csv("all_2007-2017_no_outliers.csv",  header=TRUE, sep=";")
# Installation
# install.packages('ggplot2')
# Loading
library(ggplot2)
#ine_data$month_code<- as.Date(ine_data$month_code , "%Y/%m/%d")
#install.packages('lubridate')
library(lubridate)
ine_data$month_code <- ymd(ine_data$month_code)
scraping_data_filtered$month_code <- ymd(scraping_data_filtered$month_code)
  
ggplot(ine_data, aes(month_code, todas_edades)) + geom_line() + xlab("") + ylab("Suicidios")

ggplot(scraping_data_filtered, aes(month_code, number_of_suicide_news)) + geom_line() + xlab("") + ylab("Numero noticias suicidio")


plot(all_data$noticias_suicidio, all_data$suicidios, col = "blue", main = "", xlab = "noticias suicidio", ylab = "suicidios")
lm(all_data$noticias_suicidio ~ all_data$suicidios)
abline(lm(all_data$noticias_suicidio ~ all_data$suicidios))

cor(all_data$noticias_suicidio, all_data$suicidios)


# chi-squared test of independence NOT VALID, categorized data only
tbl = table(all_data$suicidios, all_data$noticias_suicidio) 
chisq.test(tbl) 

# https://github.com/AnaBPazos/AlterCorr/blob/master/R/AlterCorrM.R

# chi-squared test of independence NOT VALID, categorized data only
# Kendall test
cor.test(all_data$suicidios, all_data$noticias_suicidio, method = "kendall")
