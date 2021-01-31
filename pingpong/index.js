const express = require("express")
const app = express()
const port = 3000
var path = require("path")

var counter = 0

app.get("/", (req, res) => {
	res.send(`pong ${counter}`)
	counter++
})

app.listen(port, () => {
	console.log(`Server started in port ${port}`)
})
