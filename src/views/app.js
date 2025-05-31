const apiBase = "/api";

async function fetchUsers() {
  const res = await fetch(`${apiBase}/users/`);
  const users = await res.json();
  const list = document.getElementById("usersList");
  list.innerHTML = users.map(u => `<li>ID: ${u.id}, Name: ${u.name}</li>`).join('');
}

async function addUser() {
  const name = document.getElementById("newUserName").value;
  await fetch(`${apiBase}/users/`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name })
  });
  fetchUsers();
}

async function deleteUser() {
  const id = document.getElementById("deleteUserId").value;
  await fetch(`${apiBase}/users/${id}`, { method: "DELETE" });
  fetchUsers();
}

async function fetchCards() {
  const res = await fetch(`${apiBase}/cards/`);
  const cards = await res.json();
  console.log(cards)
  const list = document.getElementById("cardsList");
  list.innerHTML = cards.map(c => `<li>ID: ${c.card_number}</li>`).join('');
}

async function addCard() {
  const card_number = document.getElementById("newCardNumber").value;
  await fetch(`${apiBase}/cards/`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ card_number })
  });
  fetchCards();
}

async function deleteCard() {
  const card_number = document.getElementById("deleteCardNumber").value;
  await fetch(`${apiBase}/cards/${card_number}`, { method: "DELETE" });
  fetchCards();
}

async function fetchTransactions() {
  const res = await fetch(`${apiBase}/transactions/`);
  const transactions = await res.json();
  const list = document.getElementById("transactionsList");
  list.innerHTML = transactions.map(t => `<li>ID: ${t.id}, Amount: ${t.amount}, Card: ${t.card_number}, Merchant: ${t.id_merchant}</li>`).join('');
}

async function fetchTransactionsByMerchant() {
  const id = document.getElementById("merchantId").value;
  const res = await fetch(`${apiBase}/transactions/${id}`);
  const transactions = await res.json();
  const list = document.getElementById("transactionsList");
  list.innerHTML = transactions.map(t => `<li>ID: ${t.id}, Amount: ${t.amount}, Card: ${t.card_number}, Merchant: ${t.id_merchant}</li>`).join('');
}

async function addTransaction() {
  const amount = parseFloat(document.getElementById("transactionAmount").value);
  const card_number = document.getElementById("transactionCardNumber").value;
  const id_merchant = parseInt(document.getElementById("transactionMerchantId").value);

  const res = await fetch(`${apiBase}/transactions/`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ amount, card_number, id_merchant })
  });

  const data = await res.json();

  if (data.error === "Fraudulent transaction detected!") {
    alert(data.error);
  } else {
    fetchTransactions();
  }
}
