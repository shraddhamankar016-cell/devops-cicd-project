pipeline {
    agent any

    environment {
        IMAGE_NAME = "yourdockerhubusername/go-app"
        IMAGE_TAG  = "v${BUILD_NUMBER}"
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/yourusername/devops-cicd-project.git'
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t $IMAGE_NAME:$IMAGE_TAG .'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
                    sh 'echo $PASS | docker login -u $USER --password-stdin'
                    sh 'docker push $IMAGE_NAME:$IMAGE_TAG'
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh "kubectl set image deployment/go-app-deployment go-app=$IMAGE_NAME:$IMAGE_TAG"
                sh "kubectl rollout status deployment/go-app-deployment"
            }
        }
    }

    post {
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed. Check logs.'
        }
    }
}
