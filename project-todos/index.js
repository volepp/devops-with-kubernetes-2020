const express = require("express")
const app = express()
const port = 3000
const bodyParser = require("body-parser")

app.use(bodyParser.urlencoded({ extended: true }))

var todos = []

app.get("/", (req, res) => {
	todosJson = JSON.stringify(todos)

	res.json(todosJson)
})

app.post("/", (req, res) => {
	todoName = req.body["todo-name"]

	todos.push(todoName)

	res.redirect("back")
})

app.listen(port, () => {
	console.log(`Server started in port ${port}`)
})