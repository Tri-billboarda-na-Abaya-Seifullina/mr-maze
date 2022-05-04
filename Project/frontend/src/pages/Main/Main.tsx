import React from "react";
import { Link } from "react-router-dom";
import Header from "../../components/Header/Header";
import "./Main.scss";

const Main = () => {
    return (
        <div className="main">
            <Header/>
            <div className="content">
                <div className="left">
                    <h1>
                        Multiplayer Mr. Maze  game. 
                    </h1>
                    <p>
                        Play with your friends and compete in getting out of the maze.
                    </p>
                    <div>
                        <button className="start">
                            <Link to="/game">
                                Start
                            </Link>
                        </button>
                        <button className="join">
                            Join
                        </button>
                    </div>
                </div>
                <div className="right">
                    <div className="dot dot-blue"/>
                    <div className="dot dot-pink"/>
                    <img src="/images/main-rocket.png" alt="" />
                </div>
            </div>
        </div>
    )
}

export default Main;