#!/usr/bin/env bash
set -euo pipefail

echo "Checking prerequisites..."
command -v go >/dev/null || { echo "Go not found. Install Go first."; exit 1; }

echo "Installing CLI tools and deps..."


# App dependencies
go get github.com/lib/pq
go get github.com/google/uuid
go get github.com/joho/godotenv

echo "Ensuring .env exists..."
if [ ! -f ".env" ]; then
  if [ -f ".env.example" ]; then
    cp .env.example .env
    echo "Created .env from .env.example"
  else
    cat > .env <<'EOF'
DB_URL="postgres://USER:PASSWORD@localhost:5432/chirpy?sslmode=disable"
EOF
    echo "Created a minimal .env (edit DB_URL)"
  fi
fi

echo "Generating sqlc code..."
sqlc generate

echo "Verifying build..."
go build ./...

echo "Done!"
