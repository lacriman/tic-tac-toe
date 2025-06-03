const X_CLASS = "x";
const O_CLASS = "o";

const gameInfo = document.getElementById("gameInfo");
const joinButton = document.getElementById("joinButton");

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

const storedName = localStorage.getItem("username");
if (!storedName) {
  const popup = new Popup({
    id: "username-popup",
    title: "Welcome to Tic Tac Toe",
    content: `
    <label id="usernameLabel" for="username">Please enter your name:</label><br>
    <input type="text" id="usernameInput" placeholder="Your name" /><br><br>
    <button id="submitNameBtn">Submit</button>      
  `,
    showImmediately: true,
    closeButton: false,
    overlay: true,
  });
} else {
}

/* --------------- Popup ------------------------- */

document.addEventListener("DOMContentLoaded", () => {
  setTimeout(() => {
    const submitNameBtn = document.getElementById("submitNameBtn");

    submitNameBtn.addEventListener("click", () => {
      const username = document.getElementById("usernameInput").value.trim();
      if (!username) {
        document.getElementById("usernameLabel").textContent =
          "You have to write your username 👇";
      } else {
        localStorage.setItem("username", username);
        document.getElementsByClassName("popup")[0].style.display = "none";
      }
    });
  }, 100);
});

/* --------------- Join Game ------------------------- */

joinButton.addEventListener("click", async () => {
  const gameId = document.getElementById("gameIdInput").value;
  const response = await fetch(`/join/${gameId}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
  });

  if (response.ok) {
    const data = await response.json();
    window.location.href = `/game${gameId}`;
  } else {
    const error = await response.text();
    gameInfo.textContent = `Error: ${error}`;
  }
});

/* --------------- Board ------------------------- */
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
  const res = await fetch("/api/game", { method: "POST" });
  const data = await res.json();
  currentGameId = data.id;
  renderBoard(data.board);
  winningMessageElement.classList.remove("show");
});
