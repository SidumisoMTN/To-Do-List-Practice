# define the image base of the application
FROM node:latest

# set the working diretory
WORKDIR /

# Copy application files from local machine
COPY main.html .

# install the dependencies
RUN npm install

# Expose the application to port 80
EXPOSE 8080

# specify the default command,
# set up the server and the port number
CMD ["http-server", "-p", "8080"]

