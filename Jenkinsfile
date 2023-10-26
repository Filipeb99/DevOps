pipeline {

    environment {
        dockerimagename = "filipeb99/go-app"
        dockerImage = ""
    }
    
    agent any

    stages {
        stage('Checkout Source') {
            steps {
                git 'https://github.com/Filipeb99/DevOps.git'
            }
        }
        
        stage('Build image') {
            steps{
                script {
                    dockerImage = docker.build dockerimagename
                }
            }
        }

        stage('Pushing Image') {
            environment {
                registryCredential = 'dockerhub-credentials'
            }
            steps {
                script {
                    docker.withRegistry( 'https://registry.hub.docker.com', registryCredential ) {
                        dockerImage.push("latest")
                    }
                }
            }
        }

        stage('Deploying Go container to Kubernetes') {
            steps {
                sh 'kubectl apply -f kubernetes.yaml'
                sh 'kubectl get services'
            }
        }
    }
}
