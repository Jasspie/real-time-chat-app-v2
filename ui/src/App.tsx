import React from "react";
import { Routes, Route } from "react-router-dom";
import JoinPage from "./pages/JoinPage";
import ChatPage from "./pages/ChatPage";

const App: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<JoinPage />} />
      <Route path="/chat" element={<ChatPage />} />
    </Routes>
  );
};

export default App;
