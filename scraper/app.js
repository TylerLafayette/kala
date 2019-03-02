const cheerio = require("cheerio")
const fs = require("fs")
const path = require("path")

// Define array of file locations and load in their data
let files = ["scraper_files/page1.htm", "scraper_files/page2.htm"]
files = files.map(e => fs.readFileSync(path.join(__dirname, e)))

let transactions = []

files.forEach(e => {
    const $ = cheerio.load(e.toString())
    const transactionElements = $(".in-transit-record,.record")
    for(let i = 0; i < transactionElements.length; i++) {
        const transactionElement = transactionElements[i]
        const description = $(transactionElement).find(".transTitleForEditDesc").text().trim()
        const price = parseFloat($(transactionElement).find(".amount").text().trim())
        const finalBalance = parseFloat($(transactionElement).find(".balance").text().trim())
        if(description.length < 3) return
        // description = description.toString().split(" ")[2]
        console.log(`${description} | $${price} | $${finalBalance}`)
        transactions.push({
            description,
            price,
            finalBalance
        })
    }
})

fs.writeFile("output.json", JSON.stringify(transactions), (err, data) => console.log("Saved file."))