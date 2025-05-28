const X_CLASS = "x";
const CIRCLE_CLASS = "circle";

const board = document.getElementById("board");
const winningMessageElement = document.getElementById("winningMessage");
const restartButton = document.getElementById("restartButton");
const winningMessageTextElement = document.querySelector(
  "[data-winning-message-text]"
);

let currentGameId = null;

(async () => {
  const response = await fetch("/api/game", {
    method: "POST",
  });
  const data = await response.json();
  currentGameId = data.id;
  startGame();
})();

async function getData(gameId) {
  const url = `/api/game/${gameId}`;
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }
    const json = await response.json();

    console.log(json);
  } catch (error) {
    console.error(error.message);
  }
}

function renderBoard(board) {
  const boardDiv = document.getElementById("board");
  boardDiv.innerHTML = "";

  board.forEach((row, rowIndex) => {
    const rowDiv = document.createElement("div");
    rowDiv.className = "row";

    row.forEach((cell, colIndex) => {
      const cellDiv = document.createElement("div");
      cellDiv.className = "cell";
      cellDiv.textContent = cell === " " ? "" : cell;

      cellDiv.setAttribute("data-row", rowIndex);
      cellDiv.setAttribute("data-col", colIndex);

      if (cell === " ") {
        cellDiv.addEventListener("click", handleClick);
      }

      rowDiv.appendChild(cellDiv);
    });

    boardDiv.appendChild(rowDiv); 
  });
}


restartButton.addEventListener("click", startGame);

function startGame() {
  fetch(`/api/game/${currentGameId}`)
    .then((res) => res.json())
    .then((data) => {
      renderBoard(data.board);
    });

  const cellElements = document.querySelectorAll("[data-cell]");

  cellElements.forEach((cell) => {
    cell.classList.remove(X_CLASS);
    cell.classList.remove(CIRCLE_CLASS);
    cell.removeEventListener("click", handleClick);
    cell.addEventListener("click", handleClick);
  });
  winningMessageElement.classList.remove("show");
}

function placeMark(cell, currentClass) {
  cell.classList.add(currentClass);
}

async function handleClick(e) {
  const cell = e.target;
  const row = parseInt(cell.getAttribute("data-row"));
  const col = parseInt(cell.getAttribute("data-col"));

  try {
    const res = await fetch(`/api/game/${currentGameId}/move`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        gameID: currentGameId,
        row,
        col,
      }),
    });

    if (!res.ok) {
      const errText = await res.text();
      throw new Error(errText || "Move failed");
    }

    const data = await res.json();
    renderBoard(data.board);
  } catch (err) {
    console.error("Move error:", err.message);
  }
}

function endGame(draw, winner) {
  if (draw) {
    winningMessageTextElement.innerText = `It's a tie 😔`;
  } else {
    winningMessageTextElement.innerText = `${winner} wins 🥳🥳🎊🎂`;
  }
  winningMessageElement.classList.add("show");
}
