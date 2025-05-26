const X_CLASS = 'x'
const CIRCLE_CLASS = 'circle'

const cellElements = document.querySelectorAll('[data-cell]')
const board = document.getElementById('board')
const winningMessageElement = document.getElementById('winningMessage')
const restartButton = document.getElementById('restartButton')
const winningMessageTextElement = document.querySelector('[data-winning-message-text]')

const response = await fetch("/api/game", { method: "POST" });
const data = await response.json();
const gameId = data.id;

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

startGame()

restartButton.addEventListener('click', startGame)

function startGame() {
  cellElements.forEach(cell => {
    cell.classList.remove(X_CLASS)
    cell.classList.remove(CIRCLE_CLASS)
    cell.removeEventListener('click', handleClick)
    cell.addEventListener('click', handleClick, { once: true })
  })
  winningMessageElement.classList.remove('show')
}

function handleClick(e) {
  const cell = e.target
  const currentClass = circleTurn ? CIRCLE_CLASS : X_CLASS
  placeMark(cell, currentClass)
  if (checkWin(currentClass)) {
    endGame(false)
  } else if (isDraw()) {
    endGame(true)
  } else {
    swapTurns()
    setBoardHoverClass()
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
