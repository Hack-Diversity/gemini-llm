import { RequestHandler } from "express"
import { filterSocialMediaPosts } from "../gemini/gemini"

export const filterJobPosts: RequestHandler = async (req, res) => {
    const inputPosts: string[] | undefined = req.body["posts"]
    const title: string | undefined = req.body["title"]

    if(!inputPosts?.length || !title) {
        res.status(400).send("Bad request")
        return
    }

    const posts = await filterSocialMediaPosts(inputPosts, title)
    res.json({posts})
}