pipeline {
    environment {
    imagename = "dev-sumica-imgproxy"
    ecrurl = "http://10.123.31.214:5000/v2/" //docker private registry
    // ecrcredentials = "canh-docker-private" //authenticator
    dockerImage = ''
    }
    agent any
    stages {
      stage('Cloning Git') {  // clone code from github
        steps {
          git branch: 'develop', credentialsId: 'github-canh-sumica-develop', url: 'https://github.com/wakumo/imgproxy.git'
        }
      }
      stage('Building image') { // build image
        steps{
          script {
            dockerImage = docker.build imagename
          }
        }
      }
      stage('Deploy Image') { //deploy image to docker registry private 
        steps{
          script {
            docker.withRegistry(ecrurl) {
              dockerImage.push("$BUILD_NUMBER")
              dockerImage.push('latest')
            }
          }
        }
      }
      stage('speak') { //notifucation to slack
        steps{
          slackSend channel: 'sumica-jenkins-build-images', message: '[Develop Sumica] [imgproxy] Build image has been completed!', color: '#1ddb46'
        }
      }
    }
  }