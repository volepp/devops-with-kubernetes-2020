const express = require("express")
const app = express()
const port = 3000
const axios = require("axios")
const fs = require("fs")
var path = require("path")

const imageDir = path.join("/", "usr", "src", "app", "public")
const imagePath = path.join(imageDir, "image.jpg")
const dateFilePath = path.join(imageDir, "imageDate.txt")

app.use(express.static(imageDir))

if(!fs.existsSync(dateFilePath)) {
	fs.writeFile(dateFilePath, "", () => {})
}

app.get("/", (req, res) => {
	res.send("Health check")
})

app.get("/images", (req, res) => {
	var dateString = getDateAsString()
	var savedDate = fs.readFileSync(dateFilePath, "utf8")

	if(savedDate !== dateString) {
		fs.writeFile(dateFilePath, `${dateString}`, () => {})

		downloadImage(imagePath)
	}

	res.sendFile(imagePath)
})

app.listen(port, () => {
	console.log(`Server started in port ${port}`)
})

getDateAsString = () => {
	var today = new Date()
	var day = String(today.getDate()).padStart(2, "0")
	var month = String(today.getMonth() + 1).padStart(2, "0")
	var year = today.getFullYear()

	return day + "-" + month + "-" + year
}

downloadImage = async (filepath) => {
	url = "https://picsum.photos/1200"
	const response = await axios.get(url, { responseType: "stream" })
	response.data.pipe(fs.createWriteStream(filepath))
}
