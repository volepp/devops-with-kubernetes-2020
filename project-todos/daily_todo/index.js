const https = require("https")

const options = {
    hostname: "en.wikipedia.org",
    path: "/wiki/Special:Random",
    method: "GET"
}

const { Sequelize, Model, DataTypes } = require("sequelize")

var postgresDB = process.env.POSTGRES_DB
var postgresUser = process.env.POSTGRES_USER
var postgresPass = process.env.POSTGRES_PASSWORD

const sequelize = new Sequelize(`postgres://${postgresUser}:${postgresPass}@dwk-postgres-svc.project:5432/${postgresDB}`)

class Todo extends Model {}
Todo.init({
	text: DataTypes.STRING
}, { sequelize, modelNam: "todo" });

(async () => {
	await sequelize.sync()

    const req = https.request(options, res => {
        var url = res.headers["location"]

        Todo.create({ text: `Read ${url}` })
    })

    req.end()
})()