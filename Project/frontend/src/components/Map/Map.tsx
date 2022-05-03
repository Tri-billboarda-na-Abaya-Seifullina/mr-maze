import React from "react";
import MovingAbility from "../../interfaces/MovingAbility";
import Position from "../../interfaces/Position";
import Cell from "../Cell/Cell";
import "./Map.scss";

interface Map {
    map: MovingAbility[][],
    playerPosition: Position;
}

const Map: React.FC<Map> = ({ map, playerPosition }) => {

    const {x,y} = playerPosition;
    const rowsCount = map.length;
    const cellsCount = map[0].length;

    return (
        <div className="map">
            {
                map.map((row: MovingAbility[], rowIndex) => (
                    <div className="cell-row" key={rowIndex}>
                        {
                            row.map((cell: MovingAbility, cellIndex) => (
                                <Cell key={cellIndex} movingAbility={cell} playerPosition={x === rowIndex && y === cellIndex} finish={rowIndex === rowsCount-1 && cellIndex === cellsCount - 1}/>
                            ))
                        }
                    </div>
                ))
            }
            <div className="player" style={{ transform: `translate(${playerPosition.y*30+5}px, ${playerPosition.x*30+5}px)` }}>

            </div>
        </div>
    );
}

export default Map;