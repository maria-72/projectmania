import { useEffect, useState } from 'react'

export function useFetch<T>(service: () => Promise<T>) {
    const [isLoading, setIsLoading] = useState(false)
    const [response, setResponse] = useState<T>()
    const [error, setError] = useState<string>()

    useEffect(() => {
        const fetchData = async () => {
            try {
                setIsLoading(true)
                const response = await service()
                setResponse(response)
            } catch (e) {
                if (typeof e === 'string') {
                    setError(e)
                }
            } finally {
                setIsLoading(false)
            }
        }
        fetchData()
    }, [service])

    return { response, error, isLoading }
}
