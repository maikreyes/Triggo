'use client'

import { useState } from "react";
import AppHeader from "./appheader";
import LabeledInput from "./labeledInput";
import SelectedInput from "./selectInput";
import { useSearchParams } from "next/navigation";
import { useFetchRepositories } from "../customHooks/useFetchRepositories";
import { useSubmitWebhook } from "../customHooks/useSubmitWebhook";

export default function WebhookForm() {

    const searchParams = useSearchParams()
    const installationId = searchParams.get("installation_id")

    const {repositories, isLoading, error} = useFetchRepositories(installationId)

    const { submitWebhook, isSubmitting, submitError, isSuccess } = useSubmitWebhook()

    const [repositoryText, setRepositoryText] = useState("")
    const [discordUrl, setDiscordUrl] = useState("")

    

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        
        if(!installationId || !repositoryText || !discordUrl) {
            alert("Please, complete all fields")
            return;
        }

        await submitWebhook({
            installation_id: parseInt(installationId, 10),
            repository: repositoryText,
            discord_url: discordUrl
        })

        if(isSuccess){
            alert("Configuration send!!");
            setDiscordUrl("")
            setRepositoryText("")
        }

    };

    return (
        <form onSubmit={handleSubmit} className="flex flex-col items-center justify-center bg-blue-300 rounded-2xl m-24 p-12 shadow-lg">
            
            <AppHeader 
                imgSrc="/Triggo.png"
            />

            {error && <p className="text-red-600 font-bold">{error}</p>}
            {submitError && <p className="text-red-600 font-bold">{submitError}</p>}

            <div className="flex flex-col gap-6 mt-12 mb-12">
                <SelectedInput
                label="Repository"
                id="Repository"
                value={repositoryText}
                onChange={setRepositoryText}
                options={isLoading ? ["Loading Repositories..."] : repositories}
                />
                <LabeledInput 
                label="Discord Webhook Url"
                type="text"
                id="Url"
                value={discordUrl}
                onChange={setDiscordUrl}
                />
            </div>
            <button 
                type="submit"
                disabled={isLoading}
                className="px-6 py-2 
                bg-white 
                text-black 
                font-semibold 
                rounded-lg 
                shadow-md 
                hover:bg-gray-100 
                transition-colors"
            >
                {isSubmitting ? "Enviando..." : "Submit"}
            </button>
        </form>
    )
}