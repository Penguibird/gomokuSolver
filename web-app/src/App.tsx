import { useState } from 'react'

import Board, { url } from './Board'

const DEFAULTCOLUMNS = 4;
const DEFAULTROWS = 4;

function App() {
  return <div>
    <h1>Piškvorky</h1>

    <Board
      columns={DEFAULTCOLUMNS}
      rows={DEFAULTROWS}
    />
  </div>
}

export default App

