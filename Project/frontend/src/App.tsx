import React from 'react';
import { Routes, Route } from 'react-router-dom';
import './App.css';
import GamePage from './pages/GamePage/GamePage';
import Main from './pages/Main/Main';

function App() {
  return (
    <Routes>
      <Route path="/" element={<Main/>}/>
      <Route path="/game" element={<GamePage/>}/>
    </Routes>
  );
}

export default App;
