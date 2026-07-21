import { useEffect, useState } from "react";


export function useFetchRepositories(installationId:string | null) {
    
    const [repositories, setRepositroies] = useState<string[]>([])
    const [isLoading, setIsLoading] = useState(false)
    const [error, setError] = useState<string | null>(null)

    interface RepositoryResponse {
        full_name:string
    }

    useEffect(() => {
        if(!installationId) return

        const fetchRepositories = async () => {
            setIsLoading(true)
            try {
                const response = await fetch(`${process.env.NEXT_PUBLIC_WEBHOOK_URL}/api/repositories?installation_id=${installationId}`)
                
                if (!response.ok) throw new Error("Error to try get repositories")

                const data : RepositoryResponse[] = await response.json()

                const repoNames = data.map((repo) => repo.full_name)
                setRepositroies(repoNames)
            }catch(err) {
                console.error(err)
                setError("Failed Server Conection")
            }finally{
                setIsLoading(false)
            }
        }

        fetchRepositories()

    }, [installationId])

    return { repositories, isLoading, error}

}
