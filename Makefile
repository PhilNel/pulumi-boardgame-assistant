GAMES_DEPLOY_DIR=deployments/games

.PHONY: games-stack
games-stack:
	cd $(GAMES_DEPLOY_DIR) && pulumi up

.PHONY: fmt
fmt:
	go fmt ./...
