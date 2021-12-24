import React from "react";
import "./App.css";
import Feed from "./views/Feed";
import Settings from "./views/Settings";
import { Routes, Route, Link, useLocation } from "react-router-dom";
import WatchVideo from "./views/WatchVideo";

function App() {
  const location = useLocation();
  return (
    <div id="app" className="App">
      {/* <header className="App-header">
        <Link to="/settings">Settings</Link> | <Link to="/feed">Feed</Link>
      </header> */}

      <div className="all-content-container">
        <Routes>
          <Route path="/" element={<Feed />} />
          <Route path="/feed" element={<Feed />} />
          <Route path="/settings" element={<Settings />} />
          <Route path="/watch" element={<WatchVideo location={location} />} />
        </Routes>
      </div>
    </div>
  );
}

export default App;
