FROM node:18-alpine
WORKDIR /app/front-end
EXPOSE 5173

ADD ./package.json ./
RUN npm install

COPY . .

ENTRYPOINT [ "npm", "run" ]
CMD [ "dev" ]
