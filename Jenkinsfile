pipeline {
  agent {
    docker {
      image 'golang:1.8'
    }

  }
  stages {
    stage('Build') {
      parallel {
        stage('Build') {
          steps {
            sh 'make build'
          }
        }
        stage('Test') {
          steps {
            echo 'Testing'
          }
        }
      }
    }
    stage('') {
      steps {
        input 'Resdy To Deploy'
      }
    }
  }
}