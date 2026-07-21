import { useState } from "react"

interface WebhookPayload {
    installation_id: number
    repository: string
    discord_url: string
}

export function useSubmitWebhook() {

    const [isSubmitting, setIsSubmitting] = useState(false)
    const [submitError, setSusbmitError] = useState<string | null>(null)
    const [isSuccess, setIsSuccess] = useState(false)


    const submitWebhook = async (payload: WebhookPayload) => {
        
        setIsSubmitting(true)
        setSusbmitError(null)
        setIsSuccess(false)

        try {

            const data = {
                installation_id: payload.installation_id,
                repository: payload.repository,
                discord_url: payload.discord_url
            }

            const response = await fetch(`${process.env.NEXT_PUBLIC_WEBHOOK_URL}/api/setup`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(data),
            })
            
            if(!response.ok){
                throw new Error("Server reject the configuration")
            }

            setIsSuccess(true)
            
        } catch (err) {
        console.error("Error to send configuration", err)
        } finally {
            setIsSubmitting(false)
        }

    }

    return {submitWebhook, isSubmitting, submitError, isSuccess}

}



