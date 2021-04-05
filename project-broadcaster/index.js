const express = require("express")
const app = express()
const port = 3000
const Discord = require("discord.js")
const nats = require("nats")

var natsName = process.env.NATS_NAME
var namespace = process.env.NAMESPACE
var discordToken = process.env.DISCORD_TOKEN

const client = new Discord.Client()
client.login(discordToken)
var clientReady = false
var dwkChannel = null

client.on("ready", () => {
    clientReady = true
    console.log("Discord client connected")

    dwkChannel = client.channels.cache.find(channel => channel.name == "dwk-broadcast")
})

const nc = nats.connect({ url: `nats://${natsName}.${namespace}:4222` }) || null

const setReadyToProcess = () => {
    const sub = nc.subscribe("todos", { queue: "broadcaster.workers" }, (msg) => {
        nc.unsubscribe(sub)
        sendDiscordMessage(msg)
    })
};

const sendDiscordMessage = (msg) => {
    const message = JSON.parse(msg)
    const msgString = JSON.stringify(message, null, 2)
    if (message.createdAt == message.updatedAt) {
        const discordMsg = `New todo created:\n\`${msgString}\``
        console.log(discordMsg)
        dwkChannel.send(discordMsg)
    } else {
        // Updated
        const discordMsg = `Todo updated:\n\`${msgString}\``
        console.log(discordMsg)

        dwkChannel.send(discordMsg)
    }
    setReadyToProcess()
};

app.get("/healthz", async (req, res) => {
    if (nc == null || !clientReady) {
        res.sendStatus(500)
    } else {
        res.sendStatus(200)
    }
})

app.listen(port, () => {
	console.log(`Server started in port ${port}`)
})

setReadyToProcess()