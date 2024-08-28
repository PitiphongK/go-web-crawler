# Mini Web Crawler🌐

A web crawler cli for small-sites. It checks how many internal links exists in a website.

## Quick Intallation
clone this repository
```
git clone https://github.com/RealNai/go-web-crawler.git
cd go-web-crawler
```
then run
```
./crawler <website> <max-go-routine> <max-page>
```
Example:
```
./crawler https://en.wikipedia.org/wiki/Main_Page 5 50
```
*max-go-routine* is the maximum number of concurent go routine that can run at one.

*max-page* is the maximum number of pages to crawl. Big website like wikipedia takes so much time to crawl. This allows the program to exit prematurely.

## Reference
A guided project from https://www.boot.dev/