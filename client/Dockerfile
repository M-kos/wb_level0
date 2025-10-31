FROM node:20 AS builder
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:1.25-alpine
COPY --from=builder /app/build /var/www/html
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
EXPOSE 3001
CMD ["nginx", "-g", "daemon off;"]
