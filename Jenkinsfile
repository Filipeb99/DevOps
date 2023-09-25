pipeline {
    agent any

    tools { go 'go1.21' }

    stages {
        stage('Build App') {
            steps {
                sh 'go mod init mainpkg'
            }
        }
        stage('Build Image') {
            steps {
                sh 'docker build . -t simple-server'
            }
        }
    }
}