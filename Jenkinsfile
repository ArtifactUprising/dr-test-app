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
            when {
                changeRequest target: 'master'
            }
            environment {
                DT_TARGET_ENV="ephemeral"
            }
            steps {
                sh '''
                  export DT_DOCKER_TAGS="${DT_DOCKER_TAGS} ${BUILD_TAG#jenkins-ArtifactUprising-}"
                  docker_build
                  docker_push
                '''
            }
        }
        /*
        stage('Deploy PR') {
            when {
                changeRequest target: 'master'
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
//                equals expected: 'master', actual: env.BRANCH_NAME
                branch 'master'
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
//                equals expected: 'master', actual: env.BRANCH_NAME
                branch "master"
            }
            environment {
                DT_TARGET_ENV = "staging"
            }
            steps {
                sh '''
                  env | sort
                '''
            }
        }
        stage('publish production release') {
            when { tag "v*" }
//            environment {}
            steps {
                sh '''
                  env | sort
                '''
            }
        }
        stage('Deploy Production') {
            when { tag "v*" }
            environment {
                DT_TARGET_ENV = "production"
            }
            steps {
                sh '''
                  env | sort
                '''
            }
        }
        */
    }
}