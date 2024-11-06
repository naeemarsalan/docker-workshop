import React, { useEffect, useState } from "react";
import axios from "axios";

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080/api/messages";

function App() {
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    axios.get(API_URL).then((response) => {
      setMessages(response.data);
    });
  }, []);

  const submitMessage = () => {
    axios.post(API_URL, { text: message }).then(() => {
      setMessages([...messages, { text: message }]);
      setMessage("");
    });
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

