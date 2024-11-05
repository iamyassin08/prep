FROM node:lts-alpine
# WORKDIR /app
RUN npm install -g pnpm
RUN npm install -g vite
RUN mkdir -p /home/node/app
# COPY frontend/package.json /app
WORKDIR /home/node/app
# RUN npm install
COPY package*.json pnpm-lock.yaml ./
RUN pnpm install
COPY ./ /home/node/app
EXPOSE 5173
CMD ["pnpm", "run", "staging", "--host"]
