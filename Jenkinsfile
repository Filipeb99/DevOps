node {
    stages {
        stage('Clone repository') {
            checkout scm
        }
        stage('Build image') {
            app = docker.build("Filipeb99/DevOps")
        }
        stage('Test image') {
            app.inside {
                sh 'echo "Tests passed!'
            }
        }
        stage('Push image') {
            sh 'echo "Tests passed!'
        }
    }
}