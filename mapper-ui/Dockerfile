FROM node:20 AS mapper-build

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./
RUN npm install
COPY . ./
RUN npm run build

FROM nginx:1.20-alpine
COPY --from=mapper-build /app/build /usr/share/nginx/html
