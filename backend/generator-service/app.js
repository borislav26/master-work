import express from 'express'
import {getCluster} from './cluster.js'
import bodyParser from 'body-parser'

const app = express()
app.use(bodyParser.text({
    extended: true,
    limit: '50mb'
}))

app.post("/generate-pdf", async (req, res) => {
    const cluster = await getCluster()
    console.log("request processed", req.body)
    res.contentType("application/pdf");
    const pdf = await cluster.execute(req.body)
    console.log("this is done")
    res.send(pdf)
})

app.listen("4005", () => {
    console.log("app is running on port")
})




