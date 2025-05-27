const X_CLASS = 'x'
const CIRCLE_CLASS = 'circle'

const board = document.getElementById('board')
const winningMessageElement = document.getElementById('winningMessage')
const restartButton = document.getElementById('restartButton')
const winningMessageTextElement = document.querySelector('[data-winning-message-text]')

let currentGameId = null;

(async () => {
  const response = await fetch("/api/game", { method: "POST" });
  const data = await response.json();
  currentGameId = data.id;
})();

async function getData(gameId) {
  const url = `/api/game/${gameId}`;
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`)
    }
    const json = await response.json();

    console.log(json);
  } catch (error) {
    console.error(error.message);
  }
}

function createGrid() {
  const board = document.getElementById('board')
  board.innerHTML = '' // Clear board before regenerating

  for (let row = 0; row < 3; row++) {
    for (let col = 0; col < 3; col++) {
      const cell = document.createElement('div')
      cell.classList.add('cell')
      cell.setAttribute('data-cell', '')
      cell.setAttribute('data-row', row)
      cell.setAttribute('data-col', col)
      board.appendChild(cell)
    }
  }
}

startGame()

restartButton.addEventListener('click', startGame)

function startGame() {
  createGrid()
  const cellElements = document.querySelectorAll('[data-cell]')

  cellElements.forEach(cell => {
    cell.classList.remove(X_CLASS)
    cell.classList.remove(CIRCLE_CLASS)
    cell.removeEventListener('click', handleClick)
    cell.addEventListener('click', handleClick, { once: true })
  })
  winningMessageElement.classList.remove('show')
}

async function handleClick(e) {
  const cell = e.target
  const row = parseInt(cell.getAttribute('data-row'))
  const col = parseInt(cell.getAttribute('data-col'))

  try {
    const res = await fetch(`/api/game/${currentGameId}/move`, {
      method: "POST",
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ gameID: currentGameId, row, col })
    })

    if (!res.ok) {
      const errText = await res.text()
      throw new Error(errText || 'Move failed')
    }

    const data = await res.json()
    updateBoardUI(data.board)
  } catch (err) {
    console.error("Move error:", err.message)
  }
}


function endGame(draw, winner) {
  if (draw) {
    winningMessageTextElement.innerText = `It's a tie 😔`
  } else {
    winningMessageTextElement.innerText = `${winner} wins 🥳🥳🎊🎂`
  }
  winningMessageElement.classList.add('show')
}
