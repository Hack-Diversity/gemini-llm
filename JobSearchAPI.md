
# Job Searching with Large Language Models 

### TL;DR

Many job boards have deprecated their job search APIs, likely to protect their data from being used to train models. The only option we have is using API endpoitns that have scrapped the data. Here’s a quick overview of the current state:

- **LinkedIn:** No job search API available.
- **Indeed:** Job search API closed; RSS feed deprecated.
- **ZipRecruiter:** No job search API available.
- **Glassdoor:** Free API closed; job search functionality available only to specific partners.
- **RapidAPI:** Offers job search APIs from various sources, though many may be scraped and have request limits. 

# Find below the details and links to APIs explored

## Rapid API

**Job Search:** ✅ 

Rapid API is a API markplace that has endpoints that connects to a variety of data sources. There is a [Jobs API](https://rapidapi.com/search/Jobs?sortBy=ByRelevance) page that has plenty of endpoints to search for jobs, however I suspect most of them are scrapped. Most of them are free but heavly capped on the amount of requests you can do. 


### Jobs APIs found in Rapid API

All of the APIs below have endpoints for job seraching.

- [Jobs API](https://rapidapi.com/Pat92/api/jobs-api14)
- [Active Jobs DB](https://rapidapi.com/fantastic-jobs-fantastic-jobs-default/api/active-jobs-db/playground/apiendpoint_bbaf2569-9650-4b39-bb27-ff69f7916a4b)
- [Indeed](https://rapidapi.com/datazen/api/indeed46/playground/apiendpoint_320b074f-0009-4acb-9f46-bbf161eb46fc)
- [Apijob](https://rapidapi.com/apijobs-apijobs-default/api/apijob-job-searching-api/playground/apiendpoint_a24b8bfe-2db6-44c3-9d62-5c27e5eecf310)



## Linkedin 

**Job Search:** ⛔


Linkedin API can be found [here](https://developer.linkedin.com/). In order to have programatic access to we need to [create an app](https://www.linkedin.com/developers/apps/new?src=or-search&veh=www.google.com) which requires a Linkedin Company Page and a Privacy Policy URL. A personal linked account wont work. The API terms of Use can be found [here] (https://www.linkedin.com/legal/l/api-terms-of-use?src=li-other&veh=www.linkedin.com). There is no API endpoint to serach for job listings on the platform. 


## Indeed 

**Job Search:** ⛔

Closed their API for job seraching. This is not available. There API documentation can be found [here] (https://docs.indeed.com/). The application for their API is [here](https://partners.indeed.com/). They had a RSS feed with jobs that is depracated too. 


## ZipRecruiter

**Job Search:** ⛔


ZipRecrruiter does not have API endpoint to obtain job listings. Find the documentation [here](https://www.ziprecruiter.com/partner/documentation/?python#ziprecruiter-partner-platform).

## Glassdoor

**Job Search:** ⛔

Free API [closed](https://help.glassdoor.com/s/article/Glassdoor-API?language=en_US) in 2021. They offer the API buy only for specific partners and approved on a case by case basis. I believe they have the job serach functionaity but it's definitly behind a the registration wall. Docuemntation [here](https://www.glassdoor.com/developer/companiesApiActions.htm#CompanySearch)


