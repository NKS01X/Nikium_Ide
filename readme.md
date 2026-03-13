# Nikium IDE - Foundation

A monorepo workspace containing barebones foundations for a web IDE. The current setup is a clean slate configured with Turborepo and pnpm for rapid development.

## 🏗️ Architecture Setup

This repository uses a **monorepo** structure powered by `pnpm` workspaces and `turbo`.

### Applications

- **`apps/api`**: A barebones Go HTTP server (`net/http`) listening on port `:8080`.
- **`apps/worker`**: A barebones Go process designed for background tasks.
- **`apps/web`**: A React application built with Vite and Tailwind CSS. 

### Packages

Shared configurations and UI components are (or will be) stored in the `packages/` directory.
- `@nikium/eslint-config`
- `@nikium/tsconfig`

---

## 🚀 Getting Started

### Prerequisites

- Node.js (v20+)
- pnpm (`npm install -g pnpm`)
- Go (v1.22+)

### Installation

To install all dependencies and link the local workspace packages, run from the root directory:

```bash
pnpm install
```

---

## 🛠️ Development

You can start all the development servers simultaneously using Turborepo:

```bash
pnpm run dev
```

### Starting Individual Services

If you want to run a specific application (e.g., just the React frontend):

```bash
pnpm run dev --filter=web
```

---

## 📁 Project Structure

```
nikium_ide/
├── apps/
│   ├── api/                 # Go HTTP Server
│   ├── web/                 # React Vite Frontend
│   └── worker/              # Go background worker
├── packages/
│   ├── eslint-config/       # Shared ESLint configurations
│   └── tsconfig/            # Shared TypeScript configurations
├── pnpm-workspace.yaml      # pnpm workspace definition
├── package.json             # Root workspace scripts and dependencies
└── turbo.json               # Turborepo pipeline configuration
```
