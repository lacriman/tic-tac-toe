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

async function init() {
  try {
    // check session
    const sessionRes = await fetch("/api/session");
    const sessionData = sessionRes.ok ? await sessionRes.json() : null;

    if (!sessionData) {
      showUsernamePopup();
      return;
    }

    const urlParams = new URLSearchParams(window.location.search); // returns "?gameId=abc123" https://localhost:3000/?gameId=abc123
    const joinGameId = urlParams.get("gameId"); // returns "abc123"

    // create new game
    const gameRes = await fetch("/api/game", { method: "POST" });
    const gameData = await gameRes.json();
    currentGameId = gameData.id;

    // start game
    startGame();
  } catch (err) {
    console.error("Initialization failed:", err.message);
  }
}

init();

/* --------------- Popup ------------------------- */
function showUsernamePopup() {
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
}

document.addEventListener("DOMContentLoaded", () => {
  document.body.addEventListener("click", async (event) => {
    if (event.target && event.target.id === "submitNameBtn") {
      const username = document.getElementById("usernameInput").value.trim();
      if (!username) {
        document.getElementById("usernameLabel").textContent =
          "You have to write your username 👇";
      } else {
        try {
          const response = await fetch("/api/session", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({ name: username }),
          });
          if (!response.ok) {
            throw new Error("Request failed");
          }

          document.getElementsByClassName("popup")[0].style.display = "none";
        } catch (err) {}
      }
    }
  });
});

/* --------------- Join Game ------------------------- */
async function joinGame(gameId, username) {
  try {
    const response = await fetch(
      `/api/game/${gameId}/join?name=${encodeURIComponent(username)}`, //encodeURIComponent("John Doe") => "John%20Doe"
      {
        method: "POST",
      }
    );

    if (!response.ok) {
      throw new Error(await response.text());
    }

    const data = await response.json();
    currentGameId = gameId;
    gameInfo.textContent = `${username} game as ${data.symbol}`;

    startPolling();
  } catch (err) {
    gameInfo.textContent = `Error: ${err.message}`;
    console.error("Join failed: ", err.message);
  }
}

/* --------------- Join Button ------------------------- */
joinButton.addEventListener("click", async () => {
  const gameId = document.getElementById("gameIdInput").value.trim();
  if (!gameId) {
    gameInfo.textContent = "Please enter a game ID";
    return;
  }

  try {
    const sessionRes = await fetch(`/api/session`);
    const sessionData = sessionRes.ok ? await sessionRes.json() : null;

    if (!sessionData) {
      gameInfo.textContent = "Please set your username first";
      setTimeout(() => {
        showUsernamePopup();
      }, 550);
      return;
    }

    await joinGame(gameId, sessionData.name);
  } catch (err) {
    gameInfo.textContent = `Error: ${err.message}`;
  }
});

/* --------------- Polling ------------------------- */
function startPolling() {
  setInterval(async () => {
    if (!currentGameId) return;

    try {
      const res = await fetch(`/api/game/${currentGameId}`);
      const data = await res.json();

      renderBoard(data.board);

      if (data.status === "won" || data.status === "draw") {
        endGame(data.winner);
      }
    } catch (err) {
      console.error("Polling error: ", err);
    }
  }, 1000);
}

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
