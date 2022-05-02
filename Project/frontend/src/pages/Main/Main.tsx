import React from "react";
import Header from "../../components/Header/Header";
import "./Main.scss";

const Main = () => {
    return (
        <div className="main">
            <Header/>
            <div className="content">
                <div className="left">
                    <h1>
                        In aenean posuere lorem risus nec. 
                    </h1>
                    <p>
                        In aenean posuere lorem risus nec. Tempor tincidunt aenean purus purus vestibulum nibh mi venenatis
                    </p>
                    <div>
                        <button className="start">
                            Start
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