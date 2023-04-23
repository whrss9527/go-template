node  {
    // jenkins web hook
    properties([
      pipelineTriggers([
       [$class: 'GenericTrigger',
        genericVariables: [
         [key: 'ref', value: '$.ref'],
         [
          key: 'before',
          value: '$.before',
          expressionType: 'JSONPath', //Optional, defaults to JSONPath
          regexpFilter: '', //Optional, defaults to empty string
          defaultValue: '' //Optional, defaults to empty string
         ]
        ],
        genericRequestVariables: [
         [key: 'requestWithNumber', regexpFilter: '[^0-9]'],
         [key: 'requestWithString', regexpFilter: '']
        ],
        genericHeaderVariables: [
         [key: 'headerWithNumber', regexpFilter: '[^0-9]'],
         [key: 'headerWithString', regexpFilter: '']
        ],

        causeString: 'Triggered on $ref',

        token: 'go-template',
        tokenCredentialId: '',

        printContributedVariables: true,
        printPostContent: true,

        silentResponse: false,
        regexpFilterText: '$ref',
        regexpFilterExpression: 'refs/heads/'+env.BRANCH_NAME
       ]
      ])
     ])

    def appimage
    // img repo
    def registry = ''
    def registryCredential = 'aliyun-docker-image-repository'

    // gitlab credentials
    stage("Checkout"){
        checkout([$class: 'GitSCM', 
            branches: [[name: '*/'+env.BRANCH_NAME]], 
            extensions: [[$class: 'SubmoduleOption', disableSubmodules: false, parentCredentials: false, recursiveSubmodules: true, reference: '', trackingSubmodules: false]], 
            userRemoteConfigs: [[credentialsId: '', url: '']]])
    }

    stage('Build') {
        appimage = docker.build registry + ":" + env.BRANCH_NAME + "-$BUILD_NUMBER"
    }

    stage('Publish') {
        docker.withRegistry('https://registry-vpc.cn-beijing.aliyuncs.com', registryCredential ) {
            appimage.push()
            appimage.push(env.BRANCH_NAME)
        }
    }

    stage('Deploy') {
        def img = "{img-repo-url}" + env.BRANCH_NAME + "-$BUILD_NUMBER"

        def namespace
        if (env.BRANCH_NAME == 'master' ) {
            namespace = '{namespace-dev}'
            sh "kubectl set image deployment/{deploymentName} {appName}="+img+" --record --namespace "+namespace
        } else if (env.BRANCH_NAME == 'staging' ) {
            namespace = '{namespace-dev-staging}'
        } else if (env.BRANCH_NAME == 'production' ) {
            namespace = '{namespace-dev-production}'
        }


    }
}
