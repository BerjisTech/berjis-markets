# syntax=docker/dockerfile:1.7

ARG ANGULAR_APP_NAME=markets-ui
ARG SKIP_TESTS=false

FROM node:20-bookworm AS builder
WORKDIR /workspace

RUN apt-get update && \
    apt-get install -y chromium && \
    rm -rf /var/lib/apt/lists/*
ENV CHROME_BIN=/usr/bin/chromium

RUN corepack enable && corepack prepare yarn@1.22.22 --activate

COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .

RUN if [ "${SKIP_TESTS}" = "false" ]; then yarn lint && yarn test --browsers=ChromeHeadless --code-coverage; fi
RUN yarn build

FROM nginx:1.27-alpine AS runtime
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /workspace/dist/markets-ui/browser /usr/share/nginx/html
EXPOSE 80
