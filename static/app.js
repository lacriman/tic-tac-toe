const X_CLASS = "x";
const O_CLASS = "o";

const boardElement = document.getElementById("board");
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

function renderBoard(board) {
  boardElement.innerHTML = "";

  board.forEach((row, rowIndex) => {
    const rowDiv = document.createElement("div");
    rowDiv.className = "row";

    row.forEach((cell, colIndex) => {
      const cellDiv = document.createElement("div");
      cellDiv.classList.add("cell");

      const lower = cell.toLowerCase();
      if (lower === "x" || lower === "o") {
        cellDiv.classList.add(lower);
        cellDiv.textContent = cell;
      } else {
        cellDiv.textContent = "";
        cellDiv.addEventListener("click", handleClick);
      }

      cellDiv.dataset.row = rowIndex;
      cellDiv.dataset.col = colIndex;

      rowDiv.appendChild(cellDiv);
    });

    boardElement.appendChild(rowDiv);
  });
}

async function startGame() {
  try {
    const res = await fetch(`/api/game/${currentGameId}`);
    const data = await res.json();
    renderBoard(data.board);
  } catch (err) {
    console.error("Failed to start game: ", err.message);
  }
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
      body: JSON.stringify({ row, col }),
    });

    if (!res.ok) {
      const errText = await res.text();
      throw new Error(errText || "Move failed");
    }

    const data = await res.json();
    renderBoard(data.board);

    if (data.status === "won") {
      endGame(data.winner);
    }
  } catch (err) {
    console.error("Move error:", err.message);
  }
}

function endGame(winner) {
  if (!winner) {
    winningMessageTextElement.innerText = `It's a tie 😔`;
  } else {
    winningMessageTextElement.innerText = `${winner} wins 🥳🥳🎊🎂`;
  }
  winningMessageElement.classList.add("show");
}

restartButton.addEventListener("click", async () => {
  winningMessageElement.classList.remove("show");
  const res = await fetch("/api/game", {method: "POST"});
  const data = await res.json();
  currentGameId = data.id;
  renderBoard(data.board)
});
