# Paper Trading MVP – Backend-First Plan

## ✅ Goal
Build a backend trading simulation system using the Alpaca Paper Trading API, with endpoints for market data, portfolio tracking, and order execution. Future-ready for user accounts and frontend integration.

---

## 🧠 Stack

- Language: Go
- Framework: Gin (API), Resty (external requests)
- Auth (later): JWT or basic session management
- DB (later): Postgres for user data, trades, leaderboard

---

## ✅ MVP Features (Backend)

### 1. Market Data
- [x] GET `/quote?symbols=AAPL,TSLA`  
  → Fetch latest stock quotes

### 2. Order Execution
- [] POST `/order`  
  Body: `{ symbol, qty, side, type }`  
  → Places a paper trade via Alpaca

### 3. Portfolio Tracking
- [] GET `/portfolio`  
  → Returns current positions, buying power, and account info

### 4. Environment Config
- [] `.env.example` for user-supplied API keys

---

## 🕹️ Optional Backend Extensions (Post-MVP)

- [ ] WebSocket stream for live quote updates
- [ ] Trade history storage (in DB)
- [ ] Leaderboard endpoint `/leaderboard`
- [ ] Simulated users (basic ID param → fake auth)
- [ ] Real user accounts (Postgres + JWT auth)

---

## 🎨 Frontend (Later)

- Basic dashboard: quotes, portfolio, buy/sell form
- User leaderboard
- Trade history view

---

## 📦 Deployment (Final phase)

- [ ] Dockerfile + deployment script
- [ ] Hosted API (Railway, Fly.io, Render)
- [ ] Live demo frontend (Vercel)
