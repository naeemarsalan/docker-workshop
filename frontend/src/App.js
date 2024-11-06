import React, { useEffect, useState } from "react";
import axios from "axios";

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080/api/messages";

function App() {
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState([]);  // Initialize messages as an empty array

  useEffect(() => {
    axios.get(API_URL)
      .then((response) => {
        setMessages(response.data);
      })
      .catch((error) => {
        console.error("Error fetching messages:", error);
        setMessages([]);  // Set to an empty array if fetching fails
      });
  }, []);

  const submitMessage = () => {
    if (message.trim()) {
      axios.post(API_URL, { text: message })
        .then(() => {
          setMessages([...messages, { text: message }]);
          setMessage("");
        })
        .catch((error) => {
          console.error("Error posting message:", error);
        });
    }
  };

  return (
    <div>
      <h1>Message Board</h1>
      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        placeholder="Enter your message"
      />
      <button onClick={submitMessage}>Submit</button>
      <ul>
        {messages.map((msg, index) => (
          <li key={index}>{msg.text}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;


