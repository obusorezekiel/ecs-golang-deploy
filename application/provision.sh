#!/bin/bash

# Function to update the package list and install dependencies
install_dependencies() {
    sudo apt-get update -y
    sudo apt-get install -y \
        docker.io \
        docker-compose
}

# Function to start and enable Docker
setup_docker() {
    sudo systemctl start docker
    sudo systemctl enable docker
    sudo usermod -aG docker $USER
}

# Function to configure Nginx
setup_nginx() {
    sudo cp /vagrant/nginx.conf /etc/nginx/nginx.conf
}

# Main function to orchestrate the setup
main() {
    install_dependencies
    setup_docker
    setup_nginx
}

# Run the main function
main
