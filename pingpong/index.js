const express = require("express")
const app = express()
const port = 3000
const { Sequelize, Model, DataTypes } = require("sequelize")
const { isNullOrUndefined } = require("util")

var postgresDB = process.env.POSTGRES_DB
var postgresUser = process.env.POSTGRES_USER
var postgresPass = process.env.POSTGRES_PASSWORD
var namespace = process.env.NAMESPACE

const sequelize = new Sequelize(`postgres://${postgresUser}:${postgresPass}@dwk-postgres-svc.main-application:5432/${postgresDB}`)

class Pong extends Model {}
Pong.init({
	amount: DataTypes.INTEGER
}, { sequelize, modelName: "pong" });

(async () => {
	await sequelize.sync()
	const allPongs = await Pong.findAll()
	if (allPongs.length == 0) {
		const pong = await Pong.create({
			amount: 0
		})
	}
})()

app.get("/", async (req, res) => {
	var pongs = await Pong.findAll()
	var amount = pongs[0].amount+1
	await Pong.update({ amount: amount }, {
		where: {
			id: pongs[0].id
		}
	})
	res.send(`pong ${amount}`)
})

app.get("/pongs", async (req, res) => {
	var pongs = await Pong.findAll()
	var amount = pongs[0].amount

	res.send(`${amount}`)
})

app.get("/healthz", async (req, res) => {
	try {
		await sequelize.authenticate()
		res.sendStatus(200)
	} catch {
		res.sendStatus(500)
	}
})

app.listen(port, () => {
	console.log(`Server started in port ${port}`)
})
