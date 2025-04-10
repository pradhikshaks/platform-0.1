FROM node:23-alpine

WORKDIR /app
COPY demo-UI/ .
RUN yarn
EXPOSE 8002
ENTRYPOINT ["yarn","run","dev"]