const express = require("express")
const app = express()
const port = 3000
const bodyParser = require("body-parser")
const { Sequelize, Model, DataTypes } = require("sequelize")
const morgan = require("morgan")

app.use(morgan("combined"))

app.use(bodyParser.urlencoded({ extended: true }))
app.use(bodyParser.json())

var postgresDB = process.env.POSTGRES_DB
var postgresUser = process.env.POSTGRES_USER
var postgresPass = process.env.POSTGRES_PASSWORD
var namespace = process.env.NAMESPACE

const sequelize = new Sequelize(`postgres://${postgresUser}:${postgresPass}@dwk-postgres-svc.${namespace}:5432/${postgresDB}`)

class Todo extends Model {}
Todo.init({
	text: DataTypes.STRING(140),
	done: DataTypes.BOOLEAN
}, { sequelize, modelName: "todo" });

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

	if (todoName.length <= 140) {
		await Todo.create({ text: todoName, done: false })
	}

	res.redirect("back")
})

app.get("/healthz", async (req, res) => {
	try {
		await sequelize.authenticate()
		res.sendStatus(200)
	} catch {
		res.sendStatus(500)
	}
})

app.put("/:id", async (req, res) => {
	var id = req.params.id
	var todo = await Todo.findByPk(id)
	console.log(id, todo)
	if(todo) {
		console.log("testi")
		console.log(req.body)
		todo.update({
			done: req.body["done"]
		  }).then((response) => {
		  	res.sendStatus(200)
		  }).catch((reason) => {
		  	console.log("Failed to update todo:", reason)
			res.sendStatus(418)
		  })
	}
})

app.listen(port, () => {
	console.log(`Server started in port ${port}`)
})