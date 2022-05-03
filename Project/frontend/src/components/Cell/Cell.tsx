import React, { StyleHTMLAttributes } from "react";
import MovingAbility from "../../interfaces/MovingAbility";
import "./Cell.scss";

interface CellProps {
    movingAbility: MovingAbility;
    playerPosition: boolean;
    finish: boolean;
}

const Cell: React.FC<CellProps> = ({ movingAbility, playerPosition = false, finish }) => {

    const styles: React.CSSProperties = {}
    let height = 30;
    let width = 30;

    if (!movingAbility.left) {
        styles.borderLeft = "2px solid #000";
        width-=2;
    };
    if (!movingAbility.right) {
        styles.borderRight = "2px solid #000";
        width-=2;
    };
    if (!movingAbility.up) {
        styles.borderTop = "2px solid #000";
        height-=2;
    };
    if (!movingAbility.down) {
        styles.borderBottom = "2px solid #000";
        height-=2;
    };

    styles.minHeight = `${height}px`
    styles.minWidth = `${width}px`

    return (
        <div className="cell" style={styles}>
            {
                // playerPosition &&
                // <div className="player">

                // </div>
            }
            {
                finish &&
                <img src="/images/finish.png" alt="Finish" />
            }
        </div>
    )
}

export default Cell;