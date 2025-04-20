# Use the slim variant of the node image
FROM node:slim@sha256:dfb18d8011c0b3a112214a32e772d9c6752131ffee512e974e59367e46fcee52

# Set the working directory to /app
WORKDIR /app

# Install required libraries, including musl and other dependencies for Next.js
RUN apt-get update && apt-get install -y \
  libc6-dev \
  && rm -rf /var/lib/apt/lists/*

# Copy package.json and package-lock.json for dependency installation
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code into the container
COPY . .

# Build the application
RUN npm run build

# Expose the application port
EXPOSE 3000

# Start the application
CMD ["npm", "start"]
