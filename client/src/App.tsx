import React from "react";
import { ChatServiceClient } from "./proto/ChatServiceClientPb";
// import Chat from "./pages/ChatPage";
import { Routes, Route } from "react-router-dom";
import JoinPage from "./pages/JoinPage";

const client = new ChatServiceClient("http://localhost:8080", null, null);

const App: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<JoinPage client={client} />} />
    </Routes>
  );
};

export default App;
