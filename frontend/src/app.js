const express = require("express")
const app = express()
const path = require("path")
const http = require("http")

app.use("/dist", express.static("dist"))
app.get("*", (req, res) => res.sendFile(path.join(__dirname + "/client/index.html")))

app.listen(5000, _ => console.log("Frontend available at :5000"))