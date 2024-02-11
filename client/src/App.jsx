import React, { useState, useEffect } from "react";
import axios from "axios";

function App() {
  const [queues, setQueues] = useState([]);

  useEffect(() => {
    fetchQueues();
  }, []);

  const fetchQueues = async () => {
    try {
      const response = await axios.get("http://localhost:8080/monitor");
      setQueues(response.data);
    } catch (error) {
      console.error("Error fetching queues:", error);
    }
  };

  return (
    <div>
      <h1>Queue Monitor</h1>
      <table>
        <thead>
          <tr>
            <th>List</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          {queues.map((queue) => (
            <tr key={queue.id}>
              <td>{queue.data}</td>
              <td>{queue.processed_at}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default App;
