const cheerio = require("cheerio")
const fs = require("fs")
const path = require("path")

// Define array of file locations and load in their data
let files = ["scraper_files/page1.htm", "scraper_files/page2.htm"]
files = files.map(e => fs.readFileSync(path.join(__dirname, e)))

files.forEach(e => {
    const $ = cheerio.load(e.toString())
    const transactionElements = $(".in-transit-record,.record")
    console.log(transactionElements.length)
})