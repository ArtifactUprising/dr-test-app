pipeline {
    agent {
        docker {
            image "966326147313.dkr.ecr.us-west-2.amazonaws.com/deploy-tools:latest"
            args '-v /var/run/docker.sock:/var/run/docker.sock \
                  -u root:root'
        }
    }
    stages {
        stage('debug') {
            steps {
                sh '''
                  env | sort
                '''
            }
        }
        stage('build') {
            when {
                changeRequest target: 'master'
            }
            environment {
                DT_TARGET_ENV="ephemeral"
                AWS_DEFAULT_REGION="us-west-2"
                AWS_ACCESS_KEY_ID=credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_ACCESS_KEY=credentials('AWS_SECRET_ACCESS_KEY')
            }
            steps {
                sh '''
                  . /root/.ashrc
                  read_config

                  export DT_DOCKER_TAGS="${DT_DOCKER_TAGS} $(echo ${JOB_NAME#*/} | tr / -)"
                  docker_build
                  docker_push
                '''
            }
        }
        stage('Deploy PR') {
            when {
                changeRequest target: 'master'
            }
            environment {
                DT_TARGET_ENV="ephemeral"
                AWS_DEFAULT_REGION="us-west-2"
                AWS_ACCESS_KEY_ID=credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_ACCESS_KEY=credentials('AWS_SECRET_ACCESS_KEY')
            }
            steps {
                sh '''
                  . /root/.ashrc
                  read_config

                  export DT_HELM_IMAGETAG=$(echo ${JOB_NAME#*/} | tr / -)
                  export DT_HELM_domainName="${DT_HELM_IMAGETAG}.app.artifactstaging.com"
                  env | sort
                '''
            }
        }
        /*
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