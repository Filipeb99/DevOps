pipeline {
    
    agent any

    stages {
        stage('Checkout Source') {
            steps {
                git 'https://github.com/Filipeb99/DevOps.git'
            }
        }

        stage('Create Module') {
            steps {
                sh '''#!/bin/bash

                    if ! [ -f "go.mod" ]; then
                        go mod init devops
                    fi
                '''
            }
        }
        
        stage('Build Image') {
            steps{
                sh 'docker build -t devops .'
            }
        }

        stage('Tag Image') {
            steps{
                sh 'docker tag devops filipeb99/go-app:latest'
            }
        }

        stage('Push Image') {
            environment {
                registryCredential = 'dockerhub-credentials'
            }
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', registryCredential) {
                        sh 'docker push filipeb99/go-app:latest'
                    }
                }
            }
        }

        stage('Deploy Go Container to Kubernetes') {
            steps {
                sh 'kubectl apply -f k8s-deployment.yml'
            }
        }
    }
}
