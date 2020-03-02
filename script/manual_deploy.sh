#!/bin/bash

# The content for these files can be found at:
# https://gitlab.com/vega-protocol/trading-core/-/settings/ci_cd
for filename in \
	"$HOME/.devnet-deploy-hosts" \
	"$HOME/.devnet-deploy-ssh-known-hosts" \
	"$HOME/.devnet-deploy-ssh-private-key" \
	"$HOME/.devnet-deploy-slack-hook-url" \
	"$HOME/.devnet-deploy-veganet-auth-mastertoken" \
	"$HOME/.devnet-deploy-veganet-users"
do
	if ! test -r "$filename" ; then
		echo "File not found/readable: $filename"
		exit 1
	fi
done

# TODO: add '-u "$(id -u):$(id -g)"' and get the build running as non-root user.

docker run --rm -ti \
 	-v "$HOME/containergo/pkg:/go/pkg" \
 	-v "$HOME/containergo/cache:/go/cache" \
 	-v "$HOME/go/src/code.vegaprotocol.io/vega:/go/src/project" \
 	-w /go/src/project \
 	-e CI=true \
	-e DEVNET_DEPLOY_HOSTS="$(cat "$HOME/.devnet-deploy-hosts")" \
	-e DEVNET_DEPLOY_SSH_KNOWN_HOSTS="$(cat "$HOME/.devnet-deploy-ssh-known-hosts")" \
	-e DEVNET_DEPLOY_SSH_PRIVATE_KEY="$(cat "$HOME/.devnet-deploy-ssh-private-key")" \
	-e DRONE=true \
	-e DRONE_COMMIT_MESSAGE="$(git log -n1 --pretty=oneline | cut -d ' ' -f 2-)" \
	-e CI_COMMIT_SHA="$(git log -n1 --pretty=oneline |cut -b1-8)" \
	-e DRONE_COMMIT_SHA="$(git log -n1 --pretty=oneline |cut -b1-8)" \
	-e GOCACHE='/go/cache' \
	-e SLACK_HOOK_URL="$(cat "$HOME/.devnet-deploy-slack-hook-url")" \
	-e VEGA_AUTH_MASTERTOKEN="$(cat "$HOME/.devnet-deploy-veganet-auth-mastertoken")" \
	-e VEGANET_USERS="$(cat "$HOME/.devnet-deploy-veganet-users")" \
 	--entrypoint /bin/bash \
 	registry.gitlab.com/vega-protocol/devops-infra/cipipeline:1.11.13 \
 	-c 'make deps && make install && ./script/deploy.sh devnet vega "/go/bin/vega:/home/vega/current/:0755" && for u in $VEGANET_USERS ; do vegaccount -addr geo.d.vega.xyz:3002 -traderid "$u" ; curl -s -XPOST -H "Content-type: application/json" -H "Authorization: Bearer $VEGA_AUTH_MASTERTOKEN" -d "{\"id\": \"$u\", \"password\": \"123\"}" https://auth.d.vega.xyz 1>/dev/null ; done'