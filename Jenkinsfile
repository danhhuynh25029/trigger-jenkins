pipeline {
  agent any
  options {
    buildDiscarder(logRotator(numToKeepStr: '5'))
  }
  environment {
    // DOCKERHUB_CREDENTIALS = credentials('dockerhub')
    CONTAINER_REGISTRY = credentials('azure-container-registry')
    IMAGE_NAME = 'myregistry25029.azurecr.io/go-service'
    LOGIN_SERVER = 'myregistry25029.azurecr.io'
  }
  stages {
    stage('Build') {
      steps {
        sh 'docker build -t $IMAGE_NAME .'
      }
    }
    stage('Login') {
      steps {
        sh 'echo $CONTAINER_REGISTRY_USR | docker login -u $CONTAINER_REGISTRY_USR -p $CONTAINER_REGISTRY_PSW $LOGIN_SERVER'
      }
    }
    stage('Push') {
      steps {
        sh 'docker push $IMAGE_NAME'
      }
    }
    stage('Cleaning Up'){
        steps{
            sh 'docker rmi -f $IMAGE_NAME'
        }
    }
  }
  post {
    always {
      sh 'docker logout'
    }
  }
}