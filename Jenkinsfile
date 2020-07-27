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
                DT_TARGET_CLUSTER="app"
                AWS_DEFAULT_REGION="us-west-2"
                AWS_ACCESS_KEY_ID=credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_ACCESS_KEY=credentials('AWS_SECRET_ACCESS_KEY')
            }
            steps {
                sh '''
                  . /root/.ashrc
                  read_config

                  set_eks_auth
                  helm_deploy
                '''
            }
        }
        stage('publish staging rc') {
            when { branch 'master' }
            environment {
                DT_TARGET_ENV="staging"
                AWS_DEFAULT_REGION="us-west-2"
                AWS_ACCESS_KEY_ID=credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_ACCESS_KEY=credentials('AWS_SECRET_ACCESS_KEY')
            }
            steps {
                sh '''
                    . /root/.ashrc

                    read_config
                    docker_promote "${DT_TARGET_ENV}"
                '''
            }
        }
        stage('Deploy Staging') {
            when { branch "master" }
            environment {
                DT_TARGET_ENV = "staging"
                DT_TARGET_CLUSTER="app"
                AWS_DEFAULT_REGION="us-west-2"
                AWS_ACCESS_KEY_ID=credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_ACCESS_KEY=credentials('AWS_SECRET_ACCESS_KEY')
            }
            steps {
                sh '''
                    . /root/.ashrc

                    read_config
                    export DT_HELM_IMAGETAG="${DT_TARGET_ENV}"
                    set_eks_auth
                    helm_deploy
                '''
            }
        }
        stage('Delete PR ENV') {
            when { branch 'master' }
            environment {
                DT_TARGET_CLUSTER="app"
                AWS_DEFAULT_REGION="us-west-2"
                AWS_ACCESS_KEY_ID=credentials('AWS_ACCESS_KEY_ID')
                AWS_SECRET_ACCESS_KEY=credentials('AWS_SECRET_ACCESS_KEY')
            }
            steps {
                sh '''
                    . /root/.ashrc

                    read_config
                    set_eks_auth
                    reap_env
                '''
            }
        }
        stage('publish production release') {
            when { tag "v*" }
            environment {
                AWS_DEFAULT_REGION="us-west-2"
                AWS_ACCESS_KEY_ID=credentials('PROD_AWS_ACCESS_KEY_ID')
                AWS_SECRET_ACCESS_KEY=credentials('PROD_AWS_SECRET_ACCESS_KEY')
            }
            steps {
                sh '''
                    . /root/.ashrc

                    read_config
                    docker_promote "${TAG_NAME}" "staging"
                '''
            }
        }
        stage('Deploy Production') {
            when { tag "v*" }
            environment {
                DT_TARGET_ENV="prod"
                DT_TARGET_CLUSTER="prod-app"
                AWS_DEFAULT_REGION="us-west-2"
                AWS_ACCESS_KEY_ID=credentials('PROD_AWS_ACCESS_KEY_ID')
                AWS_SECRET_ACCESS_KEY=credentials('PROD_AWS_SECRET_ACCESS_KEY')
            }
            steps {
                sh '''
                    . /root/.ashrc

                    read_config
                    export DT_HELM_IMAGETAG="${TAG_NAME}"
                    set_eks_auth
                    kubectl get ns
                    helm_deploy
                '''
            }
        }
    }
}