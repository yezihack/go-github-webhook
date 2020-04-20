package main

import (
	"testing"
)

func TestHashing(t *testing.T) {
	// Hash
	hash := HashPayload("secret", []byte(TEST_HASH_PAYLOAD))
	// Check that it matches up with what GitHub computed
	if hash != "178bc5c43ed5fab45dd958d7f057697b9a7a1278" {
		t.Errorf("Unexpected hash: %s", hash)
	}
}

// A real payload from GitHub
const TEST_HASH_PAYLOAD = `{"zen":"Keep it logically awesome.","hook_id":6508402,"hook":{"url":"https://api.github.com/repos/AaronO/go-git-http/hooks/6508402","test_url":"https://api.github.com/repos/AaronO/go-git-http/hooks/6508402/test","ping_url":"https://api.github.com/repos/AaronO/go-git-http/hooks/6508402/pings","id":6508402,"name":"web","active":true,"events":["*"],"config":{"url":"http://dummy.com","content_type":"json","insecure_ssl":"0","secret":"********"},"last_response":{"code":null,"status":"unused","message":null},"updated_at":"2015-11-25T16:41:29Z","created_at":"2015-11-25T16:41:29Z"},"repository":{"id":22745825,"name":"go-git-http","full_name":"AaronO/go-git-http","owner":{"login":"AaronO","id":949223,"avatar_url":"https://avatars.githubusercontent.com/u/949223?v=3","gravatar_id":"","url":"https://api.github.com/users/AaronO","html_url":"https://github.com/AaronO","followers_url":"https://api.github.com/users/AaronO/followers","following_url":"https://api.github.com/users/AaronO/following{/other_user}","gists_url":"https://api.github.com/users/AaronO/gists{/gist_id}","starred_url":"https://api.github.com/users/AaronO/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/AaronO/subscriptions","organizations_url":"https://api.github.com/users/AaronO/orgs","repos_url":"https://api.github.com/users/AaronO/repos","events_url":"https://api.github.com/users/AaronO/events{/privacy}","received_events_url":"https://api.github.com/users/AaronO/received_events","type":"User","site_admin":false},"private":false,"html_url":"https://github.com/AaronO/go-git-http","description":"A Smart Git Http server library in Go (golang)","fork":false,"url":"https://api.github.com/repos/AaronO/go-git-http","forks_url":"https://api.github.com/repos/AaronO/go-git-http/forks","keys_url":"https://api.github.com/repos/AaronO/go-git-http/keys{/key_id}","collaborators_url":"https://api.github.com/repos/AaronO/go-git-http/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/AaronO/go-git-http/teams","hooks_url":"https://api.github.com/repos/AaronO/go-git-http/hooks","issue_events_url":"https://api.github.com/repos/AaronO/go-git-http/issues/events{/number}","events_url":"https://api.github.com/repos/AaronO/go-git-http/events","assignees_url":"https://api.github.com/repos/AaronO/go-git-http/assignees{/user}","branches_url":"https://api.github.com/repos/AaronO/go-git-http/branches{/branch}","tags_url":"https://api.github.com/repos/AaronO/go-git-http/tags","blobs_url":"https://api.github.com/repos/AaronO/go-git-http/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/AaronO/go-git-http/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/AaronO/go-git-http/git/refs{/sha}","trees_url":"https://api.github.com/repos/AaronO/go-git-http/git/trees{/sha}","statuses_url":"https://api.github.com/repos/AaronO/go-git-http/statuses/{sha}","languages_url":"https://api.github.com/repos/AaronO/go-git-http/languages","stargazers_url":"https://api.github.com/repos/AaronO/go-git-http/stargazers","contributors_url":"https://api.github.com/repos/AaronO/go-git-http/contributors","subscribers_url":"https://api.github.com/repos/AaronO/go-git-http/subscribers","subscription_url":"https://api.github.com/repos/AaronO/go-git-http/subscription","commits_url":"https://api.github.com/repos/AaronO/go-git-http/commits{/sha}","git_commits_url":"https://api.github.com/repos/AaronO/go-git-http/git/commits{/sha}","comments_url":"https://api.github.com/repos/AaronO/go-git-http/comments{/number}","issue_comment_url":"https://api.github.com/repos/AaronO/go-git-http/issues/comments{/number}","contents_url":"https://api.github.com/repos/AaronO/go-git-http/contents/{+path}","compare_url":"https://api.github.com/repos/AaronO/go-git-http/compare/{base}...{head}","merges_url":"https://api.github.com/repos/AaronO/go-git-http/merges","archive_url":"https://api.github.com/repos/AaronO/go-git-http/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/AaronO/go-git-http/downloads","issues_url":"https://api.github.com/repos/AaronO/go-git-http/issues{/number}","pulls_url":"https://api.github.com/repos/AaronO/go-git-http/pulls{/number}","milestones_url":"https://api.github.com/repos/AaronO/go-git-http/milestones{/number}","notifications_url":"https://api.github.com/repos/AaronO/go-git-http/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/AaronO/go-git-http/labels{/name}","releases_url":"https://api.github.com/repos/AaronO/go-git-http/releases{/id}","created_at":"2014-08-08T04:18:00Z","updated_at":"2015-11-04T16:44:09Z","pushed_at":"2015-05-04T18:42:16Z","git_url":"git://github.com/AaronO/go-git-http.git","ssh_url":"git@github.com:AaronO/go-git-http.git","clone_url":"https://github.com/AaronO/go-git-http.git","svn_url":"https://github.com/AaronO/go-git-http","homepage":null,"size":439,"stargazers_count":43,"watchers_count":43,"language":"Go","has_issues":true,"has_downloads":true,"has_wiki":true,"has_pages":false,"forks_count":3,"mirror_url":null,"open_issues_count":1,"forks":3,"open_issues":1,"watchers":43,"default_branch":"master"},"sender":{"login":"AaronO","id":949223,"avatar_url":"https://avatars.githubusercontent.com/u/949223?v=3","gravatar_id":"","url":"https://api.github.com/users/AaronO","html_url":"https://github.com/AaronO","followers_url":"https://api.github.com/users/AaronO/followers","following_url":"https://api.github.com/users/AaronO/following{/other_user}","gists_url":"https://api.github.com/users/AaronO/gists{/gist_id}","starred_url":"https://api.github.com/users/AaronO/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/AaronO/subscriptions","organizations_url":"https://api.github.com/users/AaronO/orgs","repos_url":"https://api.github.com/users/AaronO/repos","events_url":"https://api.github.com/users/AaronO/events{/privacy}","received_events_url":"https://api.github.com/users/AaronO/received_events","type":"User","site_admin":false}}`
