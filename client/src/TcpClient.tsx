import React, { useState, useEffect, useRef } from 'react';

interface Message {
    text: string;
    type: 'sent' | 'received';
    timestamp: Date;
}

export default function TcpClient(): React.ReactElement {
    const [messages, setMessages] = useState<Message[]>([]);
    const [inputMessage, setInputMessage] = useState('');
    const [isConnected, setIsConnected] = useState(false);
    const wsRef = useRef<WebSocket | null>(null);

    useEffect(() => {
        // Create WebSocket connection through nginx proxy
        const ws = new WebSocket('ws://localhost/ws');
        wsRef.current = ws;

        ws.onopen = () => {
            setIsConnected(true);
            setMessages((prev: Message[]) => [...prev, {
                text: 'Connected to server',
                type: 'received',
                timestamp: new Date()
            }]);
        };

        ws.onmessage = (event) => {
            setMessages((prev: Message[]) => [...prev, {
                text: event.data,
                type: 'received',
                timestamp: new Date()
            }]);
        };

        ws.onclose = () => {
            setIsConnected(false);
            setMessages((prev: Message[]) => [...prev, {
                text: 'Disconnected from server',
                type: 'received',
                timestamp: new Date()
            }]);
        };

        ws.onerror = (error) => {
            console.error('WebSocket error:', error);
            setMessages((prev: Message[]) => [...prev, {
                text: 'Error connecting to server',
                type: 'received',
                timestamp: new Date()
            }]);
        };

        return () => {
            ws.close();
        };
    }, []);

    const sendMessage = () => {
        if (!inputMessage.trim() || !wsRef.current) return;

        wsRef.current.send(inputMessage);
        setMessages((prev: Message[]) => [...prev, {
            text: inputMessage,
            type: 'sent',
            timestamp: new Date()
        }]);
        setInputMessage('');
    };

    return (
        <div className="p-4 max-w-md mx-auto">
            <div className="mb-4">
                <div className="flex items-center gap-2">
                    <div className={`w-3 h-3 rounded-full ${isConnected ? 'bg-green-500' : 'bg-red-500'}`} />
                    <span className="text-sm">{isConnected ? 'Connected' : 'Disconnected'}</span>
                </div>
            </div>

            <div className="h-96 overflow-y-auto border rounded-lg p-4 mb-4 bg-gray-50">
                {messages.map((msg, index) => (
                    <div
                        key={index}
                        className={`mb-2 p-2 rounded-lg ${msg.type === 'sent' ? 'bg-blue-100 ml-auto' : 'bg-gray-100'
                            } max-w-[80%] ${msg.type === 'sent' ? 'ml-auto' : 'mr-auto'}`}
                    >
                        <div className="text-sm">{msg.text}</div>
                        <div className="text-xs text-gray-500 mt-1">
                            {msg.timestamp.toLocaleTimeString()}
                        </div>
                    </div>
                ))}
            </div>

            <div className="flex gap-2">
                <input
                    type="text"
                    value={inputMessage}
                    onChange={(e) => setInputMessage(e.target.value)}
                    onKeyPress={(e) => e.key === 'Enter' && sendMessage()}
                    placeholder="Type a message..."
                    className="flex-1 p-2 border rounded-lg"
                    disabled={!isConnected}
                />
                <button
                    onClick={sendMessage}
                    disabled={!isConnected}
                    className="px-4 py-2 bg-blue-500 text-white rounded-lg disabled:bg-gray-300"
                >
                    Send
                </button>
            </div>
        </div>
    );
} 
