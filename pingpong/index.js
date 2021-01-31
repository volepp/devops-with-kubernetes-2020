const express = require("express")
const app = express()
const port = 3000
const fs = require("fs")
var path = require("path")

var counter = 0

const directory = path.join("/", "usr", "src", "app", "files")
const filePath = path.join(directory, "pingpong.txt")

fs.writeFile(filePath, `${counter}`, () => {
	console.log(`Wrote to file ${filePath}`)
})

app.get("/", (req, res) => {
	counter++
	fs.writeFile(filePath, `${counter}`, () => {})
	res.send(`pong ${counter}`)
})

app.listen(port, () => {
	console.log(`Server started in port ${port}`)
})
