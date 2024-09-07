import { GoogleGenerativeAI } from '@google/generative-ai'

export async function filterSocialMediaPosts(inputPosts: string[], title: string) : Promise<string[]> {
	const placeholder = `
		You are a social media post filter. I am a ${title} who is looking for a job. 
		I want to remove social media posts that are not relevant to my job search and title.
		When providing a response ONLY return a JSON list of the relevant social media posts.
		DO NOT return code blocks or any other information, ONLY the JSON.
		Below are the social media posts that I would like you to filter:

	`

    // Generate the prompt by appending the posts to the placeholder
    const prompt = placeholder + inputPosts.map(post => "> " + post).join("\n")

	// Generate the content
    try {
        const res = await generate(prompt)
        return JSON.parse(res)
    } catch (e) {
        console.error(e)
        return []
    }
}

/**
 * Generate content using the provided prompt.
 */
async function generate(prompt: string) : Promise<string> {
	const client = getClient()
    const model = client.getGenerativeModel({ model: "gemini-1.5-pro" });

	const res =  await model.generateContent(prompt);

    return res.response.text()
}

/**
* Get a new instance of GoogleGenerativeAI with the provided API key.
*/
function getClient() : GoogleGenerativeAI {
	return new GoogleGenerativeAI(process.env.GEMINI_API_KEY || "");
}

