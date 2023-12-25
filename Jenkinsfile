pipeline {
  agent any
  options {
    buildDiscarder(logRotator(numToKeepStr: '5'))
  }
  environment {
    DOCKERHUB_CREDENTIALS = credentials('dockerhub')
  }
  stages {
    stage('Build') {
      steps {
        sh 'docker build -t 25029/go-service .'
      }
    }
    stage('Login') {
      steps {
        sh 'docker login -u $DOCKERHUB_CREDENTIALS_USR -p $DOCKERHUB_CREDENTIALS_PSW'
      }
    }
    stage('Push') {
      steps {
        sh 'docker push 25029/go-service'
      }
    }
    stage('Cleaning Up'){
        steps{
            sh "docker rmi -f 25029/go-service"
        }
    }
  }
  post {
    always {
      sh 'docker logout'
    }
  }
}