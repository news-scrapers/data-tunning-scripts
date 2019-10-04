setwd("/Users/hugojosebello/Documents/git-repos/data-tunning-scripts/plot-time-series-data")
ine_data <- read.csv("./procesado_suicidio_edad_ambos_hasta_2007.csv",  header=TRUE, sep=";")
scraping_data_filtered <- read.csv("./scraped_news_suicide_data_sorted_2017_2007_filtrado_no_outliers.csv",  header=TRUE, sep=";")
all_data <- read.csv("./all.csv",  header=TRUE, sep=";")
# Installation
# install.packages('ggplot2')
# Loading
library(ggplot2)
#ine_data$month_code<- as.Date(ine_data$month_code , "%Y/%m/%d")
#install.packages('lubridate')
library(lubridate)
ine_data$month_code <- ymd(ine_data$month_code)
scraping_data_filtered$month_code <- ymd(scraping_data_filtered$month_code)
all_data$month_code <- ymd(all_data$month_code)

ggplot(ine_data, aes(month_code, todas_edades)) + geom_line() + xlab("") + ylab("Suicidios")

ggplot(scraping_data_filtered, aes(month_code, number_of_suicide_news)) + geom_line() + xlab("") + ylab("Numero noticias suicidio")

ggplot(all_data, aes(month_code, media_sentimiento_noticias)) + geom_line() + xlab("") + ylab("Numero noticias suicidio")


plot(all_data$noticias_suicidio, all_data$suicidios, col = "blue", main = "", xlab = "noticias suicidio", ylab = "suicidios")
lm(all_data$noticias_suicidio ~ all_data$suicidios)
abline(lm(all_data$noticias_suicidio ~ all_data$suicidios))


plot(all_data$media_sentimiento_noticias, all_data$suicidios, col = "blue", main = "", xlab = "sentimiento noticias suicidio", ylab = "suicidios")
lm(all_data$media_sentimiento_noticias ~ all_data$suicidios)
abline(lm(all_data$media_sentimiento_noticias ~ all_data$suicidios))

# Pearson correlation: not significative, although it is not the best method since we are
# working with temporal series

cor(all_data$noticias_suicidio, all_data$suicidios)
cor(all_data$media_sentimiento_noticias, all_data$suicidios)

# https://github.com/AnaBPazos/AlterCorr/blob/master/R/AlterCorrM.R

# Kendall test (non parametric test, alternative to pearson)
cor.test(all_data$suicidios, all_data$noticias_suicidio, method = "kendall")
# it is significative but we can not fully trust it since we are working with time series


# cross-correlation of TIME SERIES
ccf(all_data$noticias_suicidio, all_data$suicidios)

ccf(all_data$media_sentimiento_noticias, all_data$suicidios)


# interpretation 
# https://support.minitab.com/en-us/minitab/18/help-and-how-to/modeling-statistics/time-series/how-to/cross-correlation/interpret-the-results/all-statistics-and-graphs/

# threadshole
2/sqrt(length(all_data$noticias_suicidio) - 15)

# example ->
# https://nwfsc-timeseries.github.io/atsa-labs/sec-tslab-correlation-within-and-among-time-series.html

# in our case: We find that suicide cases are relatively high after a periode of 10 to 20 months of higher number of news
