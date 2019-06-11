pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd moviecontrol && go build'
                sh 'cd roomcontrol && go build'
                sh 'cd showcontrol && go build'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd moviecontrol && go test -cover'
                sh 'cd roomcontrol && go test -cover'
                sh 'cd showcontrol && go test -cover'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'golangci-lint run --deadline 20m --enable-all'
            }
        }
        stage('Build Docker Image') {
            agent any
            steps {
                sh 'echo skip docker'
            }
        }
    }
}
