pipeline {
  agent {
    docker {
      image 'golang:1.8'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh 'make build'
      }
    }
  }
}