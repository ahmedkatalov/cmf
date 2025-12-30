# Build stage
FROM node:20-alpine AS builder

WORKDIR /app
COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build

# Run stage (nginx)
FROM nginx:alpine

# SPA routing (чтобы /login, /app и т.д. работали)
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Vite build output обычно в /dist
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
