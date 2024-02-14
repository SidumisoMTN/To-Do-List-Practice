# define the image base of the application
FROM node:latest

# set the working diretory
WORKDIR /

# Copy application files from local machine
COPY main.html .

# clean the NPM Cache
RUN npm cache clean --force

# install the dependencies
# RUN npm install
RUN RUN apt-get update && apt-get install -y npm

# Expose the application to port 8080
EXPOSE 8080

# specify the default command,
# set up the server and the port number
CMD ["http-server", "-p", "8080"]

