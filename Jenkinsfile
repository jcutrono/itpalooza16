#!groovy
import groovy.json.JsonOutput

// Loads the standardBuild function/step from workflowLibs.git/vars/standardBuild.groovy
// and invokes it.
standardBuild {
    environment = 'golang:1.7.0'
    mainScript = '''
go version
go build -v hello-world.go
'''
    postScript = '''
ls -l
./hello-world
'''
}

// Add whichever params you think you'd most want to have
// replace the slackURL below with the hook url provided by
// slack when you configure the webhook
def notifySlack(text) {
    def slackURL = 'https://hooks.slack.com/services/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'
    def payload = JsonOutput.toJson([text      : text,
                                     channel   : "#ci",
                                     username  : "jenkins",
                                     icon_emoji: ":jenkins:"])
    sh "curl -X POST --data-urlencode \'payload=${payload}\' ${slackURL}"
}