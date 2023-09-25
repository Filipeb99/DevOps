pipeline {
    agent any

    tools { go 'go1.21' }

    stages {
        stage('Build App') {
            sh 'go mod init mainpkg'
        }
        stage('Build Image') {
            docker build -t simple-server .
        }
        stage('Test image') {
            docker run -p 8000:8000 -t simple-server
        }
        stage('Push image') {
            sh 'echo "docker push ..."'
        }
    }
}