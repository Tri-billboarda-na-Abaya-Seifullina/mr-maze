import React from "react";
import Game from "../../components/Game/Game";
import Header from "../../components/Header/Header";
import "./GamePage.scss";

const GamePage = () => {
    return (
        <div className="game-page">
            <Header/>
            <Game/>
        </div>
    )
}

export default GamePage;