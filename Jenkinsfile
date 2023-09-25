pipeline {
    agent any

    tools { go 'go1.21' }

    stages {
        stage('Build App') {
            sh 'go mod init mainpkg'
        }
        stage('Build Image') {
            app = docker.build("Filipeb99/DevOps")
        }
        stage('Test image') {
            app.inside {
                sh 'echo "Tests passed!'
            }
        }
        stage('Push image') {
            sh 'echo "docker push ..."'
        }
    }
}