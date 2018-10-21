#!/usr/bin/env groovy
pipeline {
agent {
        docker {
            image 'golang'
        }
    }
    stages {
       stage('Test') {
            steps {
                echo 'Testing...'
                sh 'go get github.com/franela/goblin'
                sh 'go test'
            }
        }
    }
}