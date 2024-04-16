// Web Scraping Script for Google's Home Page

// Load the Google France home page into 'page'
page = PAGE "http://www.google.fr";

// Select up to 3 <div> elements from the loaded page and assign them to 'x'
SELECT "div" AS x FROM page LIMIT 3;

// Return the text content of each selected <div> element
RETURN TEXT x;