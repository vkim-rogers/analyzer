def String determineRepoName() {
    return scm.getUserRemoteConfigs()[0].getUrl().tokenize('/').last().split("\\.")[0]
}

def stageSwitcher = [
    printEnvironmentVariables : true,
    qualityGates : false,
    dockerImage : true,
    updateWorkloadsInK8s : true,
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

        stage('Update workloads in k8s'){
            when { expression { "${stageSwitcher.updateWorkloadsInK8s}" == "true" } }
            steps {
                container('kubectl') {
                    sh "apk --no-cache --update add grep"

                    sh '''
                        kubectl get pods --all-namespaces > /tmp/pods.info
                        cat /tmp/pods.info
                        grep -E analyzer-[a-z0-9]{8,12}-[a-z0-9]{4,6} /tmp/pods.info > /tmp/pods.info.filtered || :
                        cat /tmp/pods.info.filtered
                        if [ `cat /tmp/pods.info.filtered | wc -l` = "0" ]
                        then
                            echo "No pods for update identified. Exiting from the container."
                            exit 0
                        fi
                    '''

                    sh '''
                        cat /tmp/pods.info.filtered | awk '{print "kubectl label pods" " " $2 " " "updateTimestamp=`date +%Y%m%d-%H%M`" " "  "--overwrite"  " " "--namespace" " "  $1}' > /tmp/kubectl.cmds
                    '''
                    sh "cat /tmp/kubectl.cmds"
                    sh '''
                        cat /tmp/kubectl.cmds | xargs -I {} sh -c '{}'
                    '''
                }

            }
        }


    }
}
