# Compile contracts for binding generation
FROM node:16-alpine3.17 as contracts-builder

RUN apk add --no-cache bash git python3 make g++ musl-dev linux-headers && corepack enable

WORKDIR /specular
COPY pnpm-lock.yaml ./
RUN pnpm fetch

COPY . /specular
RUN pnpm install --offline --frozen-lockfile

WORKDIR /specular/clients/geth/specular
RUN make contracts

# Build geth binary
FROM golang:1.19-alpine3.17 as geth-builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

WORKDIR /specular/clients/geth/go-ethereum
COPY --from=contracts-builder /specular/clients/geth/go-ethereum/go.mod ./
COPY --from=contracts-builder /specular/clients/geth/go-ethereum/go.sum ./
RUN go mod download

WORKDIR /specular/clients/geth/specular
COPY --from=contracts-builder /specular/clients/geth/specular/go.mod ./
COPY --from=contracts-builder /specular/clients/geth/specular/go.sum ./
RUN go mod download

COPY --from=contracts-builder /specular /specular

WORKDIR /specular/clients/geth/specular
RUN make geth-docker

# Specular hardhat image
FROM contracts-builder

# We need geth for genesis export
COPY --from=geth-builder /specular/clients/geth/specular/build/bin/geth /specular/clients/geth/specular/build/bin/geth

WORKDIR /specular/contracts

EXPOSE 8545
CMD ["npx", "hardhat", "node", "--hostname", "0.0.0.0"]
