import express from 'express'
import { filterJobPosts } from './controllers/socialMediaPosts'

const app = express()
app.use(express.json());


app.get('/', (req, res) => {
    res.send('^_^')
})
app.post('/filter', filterJobPosts)

const PORT = process.env.PORT || 8080
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`)
})
