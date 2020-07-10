pipeline {
    agent {
        docker {
            image "966326147313.dkr.ecr.us-west-2.amazonaws.com/deploy-tools:latest"
            args '-v /var/run/docker.sock:/var/run/docker.sock \
                  -u root:root'
        }
    }
//    agent any
    stages {
        stage('build') {
//            agent {
//                docker {
//                    image "966326147313.dkr.ecr.us-west-2.amazonaws.com/deploy-tools:latest"
//                    args '-v /var/run/docker.sock:/var/run/docker.sock \
//                          -u root:root'
//                }
//            }
            when {
                equals expected: 'master', actual: env.CHANGE_TARGET
            }
//            environment {}
            steps {
                sh '''
                  env | sort
                  echo 'foo'
                  docker build -t dr-test-app .
                '''
            }
        }
        stage('Deploy PR') {
            when {
                equals expected: 'master', actual: env.CHANGE_TARGET
            }
            environment {
                DT_TARGET_ENV = "ephemeral"
            }
            steps {
                sh '''
                  env | sort
                '''
            }
        }
        stage('publish staging rc') {
            when {
                equals expected: 'master', actual: env.BRANCH_NAME
            }
//            environment {}
            steps {
                sh '''
                  env | sort
                '''
            }
        }
        stage('Deploy Staging') {
            when {
                equals expected: 'master', actual: env.BRANCH_NAME
            }
            environment {
                DT_TARGET_ENV = "staging"
            }
            steps {
                sh '''
                  env | sort
                '''
            }
            input {
                message "Deploy build to staging?"
                ok "Deploy to Staging"
            }
            options {
                timeout(time: 20, unit: 'MINUTES')
            }
        }
        stage('publish production release') {
//            when {}
//            environment {}
            steps {
                sh '''
                  env | sort
                '''
            }
        }
        stage('Deploy Production') {
//            when {}
            environment {
                DT_TARGET_ENV = "production"
            }
            steps {
                sh '''
                  env | sort
                '''
            }
            input {
                message "Deploy build to production?"
                ok "Deploy to Production"
            }
            options {
                timeout(time: 20, unit: 'MINUTES')
            }
        }
    }
}