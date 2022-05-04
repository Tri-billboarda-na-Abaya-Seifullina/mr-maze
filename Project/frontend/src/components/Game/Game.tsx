import React, { SyntheticEvent, useCallback, useEffect, useState } from "react";
import MovingAbility from "../../interfaces/MovingAbility";
import Position from "../../interfaces/Position";
import map from "../../utils/map";
import { Link, useNavigate } from "react-router-dom";
import { useEventListener } from "../../utils/useEventListener";
import Map from "../Map/Map";
import "./Game.scss";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContent from "@mui/material/DialogContent";
import DialogActions from "@mui/material/DialogActions";
import MazeService from "../../services/maze.service";

const Game = () => {

    const [playerPosition, setPlayerPosition] = useState<Position>({x: 0, y: 0})
    const [movingEnabled, setMovingEnabled] = useState<boolean>(true)
    const [showWinnerModal, setShowWinnerModal] = useState<boolean>(false);
    const [map, setMap] = useState<MovingAbility[][]>([[]]);
    const navigate = useNavigate();
    console.log(map)

    useEffect(() => {
        MazeService.generate()
        .then(res => setMap(res.data.map));
    }, [])

    const reload = (_e: SyntheticEvent) => {
        navigate(0);
    }

    const handler = ({ key }: KeyboardEvent) => {
        if (!movingEnabled) return;
        if (["ArrowUp", "ArrowDown", "ArrowLeft", "ArrowRight"].includes(key)) {
            setMovingEnabled(false);
            let {x,y} = playerPosition;
            let newX = x;
            let newY = y;
            if (key === "ArrowUp" && map[x][y].up) {
                newX--;
            } else if (key === "ArrowDown"  && map[x][y].down) {
                newX++;
            } else if (key === "ArrowLeft"  && map[x][y].left) {
                newY--;
            } else if (key === "ArrowRight" && map[x][y].right) {
                newY++;
            }
            if (newX == map.length - 1 && newY === map[0].length - 1) {
                setShowWinnerModal(true);
            }
            setPlayerPosition(({x,y}) => ({ x: newX, y: newY }))
            setTimeout(() => setMovingEnabled(true), 80);
        }
      }

    useEventListener('keydown', handler);
    return (
        <div className="game">
            <div className="players-board">
                <div className="players-board-player">
                    <span className="player-name">
                        {`Player 1: `}
                    </span>
                    <span className="player-color">

                    </span>
                </div>
            </div>
            <Map map={map} playerPosition={playerPosition}/>
            <Dialog open={showWinnerModal} fullWidth>
                <DialogTitle>
                    Congratulations!
                </DialogTitle>
                <DialogContent >
                    Player 1 is a winner!
                </DialogContent>
                <DialogActions>
                    <button className="custom-button primary" onClick={reload}>
                        New Game
                    </button>
                    <button className="custom-button secondary">
                        <Link to="/">
                            Back to main
                        </Link>
                    </button>
                </DialogActions>
            </Dialog>
        </div>
    )
}

export default Game;