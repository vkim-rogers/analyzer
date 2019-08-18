def String determineRepoName() {
    return scm.getUserRemoteConfigs()[0].getUrl().tokenize('/').last().split("\\.")[0]
}

def stageSwitcher = [
    printEnvironmentVariables : true,
    qualityGates : false,
    dockerImage : true,
    ]

/* Declarative pipeline must be enclosed within a pipeline block */
pipeline {

    environment {
        IMG_REG_CREDS=credentials("docker-registry-credentials")
        SONAR_HOST_URL="http://placeholder"
    }

    agent {
      kubernetes {
        label "ci-${determineRepoName()}"
        yamlFile 'k8s-ci-pod.yaml'
        defaultContainer 'jnlp'
      }
    }
	triggers {
        pollSCM('H/5 * * * *')
    }

    
    stages {
        stage('Print environment variables') {
            when { expression { "${stageSwitcher.printEnvironmentVariables}" == "true" } }

            steps {
                script {
                    echo sh(script: 'env|sort', returnStdout: true);
                }
            }
        }


		stage('Quality Gates') {
            when { expression { "${stageSwitcher.qualityGates}" == "true" } }

            steps {
                container('dnd'){
                    sh 'placeholder'
                }
            }
        }

        stage('Docker image'){
            when { expression { "${stageSwitcher.dockerImage}" == "true" } }

            steps {
                // BUILD CONTAINER
                container('dnd'){
                    sh "docker build --tag dawnbreather/analyzer-api:latest ."
                    sh "docker login -u'${IMG_REG_CREDS_USR}' -p'${IMG_REG_CREDS_PSW}'"
                    sh "docker push dawnbreather/analyzer-api:latest"
                }
            }
        }

    }
}
