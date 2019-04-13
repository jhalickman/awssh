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
            sh 'go build'
            echo 'go build'
          }
        }
        stage('Test') {
          steps {
            echo 'Testing'
          }
        }
      }
    }
    stage('error') {
      steps {
        input 'Resdy To Deploy'
      }
    }
  }
}