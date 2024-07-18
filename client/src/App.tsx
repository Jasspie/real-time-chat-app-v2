import React, { useState, useRef } from "react";
import { ChatServiceClient } from "./ChatServiceClientPb";
// import Chat from "./pages/ChatPage";
import JoinPage from "./pages/JoinPage";

const client = new ChatServiceClient("http://localhost:8080", null, null);

const App: React.FC = () => {
  const inputRef = useRef<HTMLInputElement>(null); // Specify the input element type
  const [submitted, setSubmitted] = useState<boolean | null>(null); // Specify the state type

  const renderChatPage = () => {
    return <></>;
    // return <ChatPage client={client} />;
  };

  const renderJoinPage = () => {
    return <JoinPage client={client} />;
  };

  return (
    <>
      <div className="container">
        <main className="main">
          {submitted ? renderChatPage() : renderJoinPage()}
        </main>
      </div>
    </>
  );
};

export default App;
