import React, { useRef } from "react";
import { useNavigate } from "react-router-dom";

const JoinPage: React.FC = () => {
  const userRef = useRef<HTMLInputElement>(null);
  const roomRef = useRef<HTMLInputElement>(null);
  const navigate = useNavigate();

  const joinHandler = async () => {
    if (!userRef.current || !roomRef.current) return;
    const userName = userRef.current.value;
    const roomName = roomRef.current.value;
    navigate("/chat", { state: { userName, roomName } });
  };

  return (
    <div className="container">
      <div>
        <h1>Chat</h1>
      </div>
      <div style={{ padding: "10px 0" }}>
        <input
          ref={userRef}
          style={{ fontSize: "1.3rem" }}
          type="text"
          id="username"
          placeholder="Your username..."
        />
      </div>
      <div style={{ padding: "10px 0" }}>
        <input
          ref={roomRef}
          style={{ fontSize: "1.3rem" }}
          type="text"
          id="room"
          placeholder="Room name..."
        />
      </div>
      <div>
        <button
          onClick={joinHandler}
          style={{
            margin: "20px",
            padding: "7px 38px",
            fontSize: "1.2em",
            boxSizing: "content-box",
            borderRadius: "4px",
          }}
        >
          Join Room
        </button>
      </div>
    </div>
  );
};

export default JoinPage;
