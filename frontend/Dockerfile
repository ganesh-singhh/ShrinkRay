
FROM node:22 AS build

WORKDIR /app/frontend

COPY . .

RUN npm install && npm run build

FROM node:22-slim

WORKDIR /app/frontend

COPY --from=build /app/frontend .

RUN npm ci --only=production && npm install -g serve

EXPOSE 3000

CMD ["sh", "-c", "serve -s builds/frontend-builds/ -l ${FRONTEND_PORT:-3000}"]