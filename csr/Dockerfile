FROM node:18-alpine3.17 as build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . /app
RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html
COPY --from=build /app/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80