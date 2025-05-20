import axios from 'axios';
import { useEffect, useState } from 'react';
import TcpClient from './TcpClient';

interface ApiResponse {
    status?: string;
    env?: string;
    message?: string;
}

function App() {
    const [data, setData] = useState<ApiResponse | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get<ApiResponse>('/api');
                console.log('Server response:', response.data);
                setData(response.data);
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();
    }, []);

    return (
        <div className="container mx-auto p-4">
            <div className="mb-8 border-2 border-red-200 text-2xl text-red-500 text-center p-4">
                {data?.message ? (
                    <p>Message: {data.message}</p>
                ) : (
                    'No data received'
                )}
            </div>
            <TcpClient />
        </div>
    );
}

export default App;

