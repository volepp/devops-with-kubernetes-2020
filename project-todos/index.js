const express = require("express")
const app = express()
const port = 3000
const bodyParser = require("body-parser")
const { Sequelize, Model, DataTypes } = require("sequelize")

app.use(bodyParser.urlencoded({ extended: true }))

var postgresDB = process.env.POSTGRES_DB
var postgresUser = process.env.POSTGRES_USER
var postgresPass = process.env.POSTGRES_PASSWORD

const sequelize = new Sequelize(`postgres://${postgresUser}:${postgresPass}@dwk-postgres-svc.project:5432/${postgresDB}`)

class Todo extends Model {}
Todo.init({
	text: DataTypes.STRING
}, { sequelize, modelNam: "todo" });

app.get("/", async (req, res) => {
	var todos = await Todo.findAll()
	todosJson = JSON.stringify(todos)

	res.json(todosJson)
});

(async () => {
	await sequelize.sync()
})()

app.post("/", async (req, res) => {
	todoName = req.body["todo-name"]

	await Todo.create({ text: todoName })

	res.redirect("back")
})

app.listen(port, () => {
	console.log(`Server started in port ${port}`)
})