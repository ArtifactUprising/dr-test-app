pipeline {
    agent any
    stages {
        stage('build') {
            agent {
                docker {
                    image "966326147313.dkr.ecr.us-west-2.amazonaws.com/deploy-tools:latest"
                    args '-v /var/run/docker.sock:/var/run/docker.sock \
                          -u root:root'
                }
            }
            when { equals expecte: 'master', actual: env.CHANGE_TARGET }
            steps {
                sh '''
                env | sort
                echo 'foo'
                docker build -t dr-test-app .
                '''
            }
        }
    }
}