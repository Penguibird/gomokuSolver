// /** @jsxImportSource @emotion/react */ 
// import { jsx, css } from '@emotion/react'
import * as React from 'react';
import styled from '@emotion/styled';
import { css } from '@emotion/react';

interface BoardProps extends React.ComponentProps<typeof Wrapper> {

};

interface SizeProps {
  columns: number
  rows: number
  squareSize?: number
}

const DEFAULTSQUARESIZE = '20px';



const Wrapper = styled.div<SizeProps>`
  display: grid;
  grid-template-columns: repeat(${props => props.columns}, ${props => props.squareSize ?? DEFAULTSQUARESIZE});
  grid-template-rows: repeat(${props => props.rows}, ${props => props.squareSize ?? DEFAULTSQUARESIZE});
`;

type SQUARESTATE = 'empty' | 'circle' | 'cross'
const Square = styled.button<{ state: SQUARESTATE }>`
border: 1px solid black;
  ${props => props.state == 'circle' && css`
    background-color: blue;
  `}
  ${props => props.state == 'cross' && css`
    background-color: red;
  `}
`;

type Point = `${number},${number}`

interface BoardState {
  Circles: Point[],
  Crosses: Point[],
}

const currentPlayerReducer = (s: 'Circles' | 'Crosses', a: 'toggle' | 'Circles' | 'Crosses') => {
  if (a == 'toggle') {
    if (s == 'Circles')
      return 'Crosses';

    else
      return 'Circles';
  }
  else
    return a;
};
export const url = 'http://localhost:8000'

const Board: React.FC<BoardProps> = ({ children, ...props }) => {
  const { columns, rows, squareSize } = props;
  const [boardState, setBoardState] = React.useState<BoardState>({ Circles: [], Crosses: [] })
  const [currentlyPlaying, updateCurrentlyPlaying] = React.useReducer(currentPlayerReducer, 'Circles');
  const togglePlayer = () => updateCurrentlyPlaying('toggle');

  const [fetching, setFetching] = React.useState(false)

  const handleSquareUpdate = (coord: Point) => {
    if (fetching)
      return
    const newBoardState = { ...boardState }
    newBoardState[currentlyPlaying].push(coord)
    setBoardState(newBoardState)
    // togglePlayer();
    setFetching(true);
    console.log(newBoardState)
    const [Row, Column] = coord.split(',').map(n => parseInt(n, 10))
    fetch(url + '/turn', {
      method: 'POST',
      body: JSON.stringify({
        Row, Column
      }),
      // headers: {
      //   'Content-Type': "application/json", 
      // },
      // mode: ''
    })
      .then(res => res.json())
      .then((boardState: { Circles: number[][], Crosses: number[][] }) => {
        setFetching(false)
        console.log(boardState)
        setBoardState({
          Circles: boardState.Circles.map(n => n.join(',') as Point),
          Crosses: boardState.Crosses.map(n => n.join(',') as Point),
        })
      })
      .catch(e => {
        alert(e);
        console.error(e)
      })
  }

  return <>
    <button onClick={() => {
      fetch(url + '/newGame')
        .then(res => res.text())
        .then(res => {
          if (res == 'SUCCESS')
            setBoardState({ Circles: [], Crosses: [] })
        })
    }}>
      New Game
    </button>
    <Wrapper {...props}>
      {Array(columns * rows).fill(null).map((_, i) => {
        const colNum = i % columns;
        const rowNum = Math.floor(i / columns)
        const coord = `${colNum},${rowNum}` as Point
        const squareState = boardState.Circles.includes(coord)
          ? 'circle'
          : boardState.Crosses.includes(coord)
            ? 'cross'
            : 'empty';

        return <Square
          key={coord}
          onClick={() => {
            if (squareState != 'empty')
              return;
            else
              handleSquareUpdate(coord)
          }}
          state={squareState}
        />
      })}
    </Wrapper>
  </>
}

export default Board;
