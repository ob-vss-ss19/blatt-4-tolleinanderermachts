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
				sh 'cd reservationcontrol && go build'
				sh 'cd usercontrol && go build'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'cd moviecontrol && go test -cover moviecontrol.go moviecontrol_test.go'
                sh 'cd roomcontrol && go test -cover roomcontrol.go roomcontrol_test.go'
                sh 'cd showcontrol && go test -cover showcontrol.go showcontrol_test.go'
                sh 'cd reservationcontrol && go test -cover reservation.go reservationcontrol_test.go'
                sh 'cd usercontrol && go test -cover usercontrol.go usercontrol_test.go'
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
                sh "docker-build-and-push -b ${BRANCH_NAME} -s moviecontrol -f moviecontrol.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s roomcontrol -f roomcontrol.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s showcontrol -f showcontrol.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s usercontrol -f usercontrol.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME} -s reservationcontrol -f reservationcontrol.dockerfile"
            }
        }
    }
}
