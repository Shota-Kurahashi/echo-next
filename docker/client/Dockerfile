# Pull official base image
FROM node:18

# Set working directory
WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install

# Add app
COPY . .

# Start app
CMD ["npm", "run", "dev"]
