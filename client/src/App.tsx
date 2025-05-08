import axios from 'axios'
import {useEffect, useState} from 'react'

interface ApiResponse {
    status?: string;
    env?: string;
    message?: string;
}

function App() {
    const [data, setData] = useState<ApiResponse | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    // Make a request to backend
    useEffect(() => {
        const fetchData = async () => {
            try {
                setLoading(true);
                const response = await axios.get<ApiResponse>('/api/health');
                console.log('Server response:', response.data);
                setData(response.data);
                setError(null);
            } catch (err: unknown) {
                const errorMessage = err instanceof Error ? err.message : 'An error occurred while fetching data';
                setError(errorMessage);
                setData(null);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    
        return () => {
            // Cleanup if needed
        }
    }, []);

    if (loading) return <div className="m-10 text-center">Loading...</div>;
    if (error) return <div className="m-10 text-center text-red-500">Error: {error}</div>;

    return (
        <div className="m-10 border-2 border-red-200 text-2xl text-red-500 text-center">
            {data ? (
                <div>
                    <p>Status: {data.status}</p>
                    <p>Environment: {data.env}</p>
                    {data.message && <p>Message: {data.message}</p>}
                </div>
            ) : 'No data received'}
        </div>
    )
}

export default App

