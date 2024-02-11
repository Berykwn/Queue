import React, { useState, useEffect } from "react";
import axios from "axios";

function App() {
  const [queues, setQueues] = useState([]);
  const [inputData, setInputData] = useState("");

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

  const handleAddData = async () => {
    try {
      const response = await axios.post("http://localhost:8080/produce", {
        data: "testdong",
      });
      console.log("Response:", response);
      setInputData("");
      alert("Data added successfully");
      fetchQueues();
    } catch (error) {
      console.error("Error adding data:", error.response);
      alert("Failed to add data: " + error.message);
    }
  };

  return (
    <div>
      <h1>Queue Monitor</h1>
      <div>
        <button onClick={handleAddData}>Add Data</button>
      </div>
      <table>
        <thead>
          <tr>
            <th>List</th>
            <th>processed at</th>
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
