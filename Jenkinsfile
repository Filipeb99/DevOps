pipeline {
    
    agent any

    stages {
        stage('Checkout Source') {
            steps {
                git 'https://github.com/Filipeb99/go-kubernetes.git'
            }
        }

        stage('Create Module') {
            steps {
                sh '''#!/bin/bash
                    if [ ! test -f "go.mod" ]
                    then
                        go mod init go-kubernetes
                    fi
                '''
            }
        }
        
        stage('Build Image') {
            steps{
                sh 'docker build -t go-kubernetes .'
            }
        }

        stage('Tag Image') {
            steps{
                sh 'docker tag go-kubernetes filipeb99/go-hello-world:latest'
            }
        }

        stage('Push Image') {
            environment {
                registryCredential = 'dockerhub-credentials'
            }
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', registryCredential) {
                        sh 'docker push filipeb99/go-hello-world:latest'
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
