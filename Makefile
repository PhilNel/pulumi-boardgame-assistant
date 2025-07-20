GAMES_DEPLOY_DIR=deployments/games
REFERENCES_DEPLOY_DIR=deployments/references

.PHONY: games-stack
games-stack:
	cd $(GAMES_DEPLOY_DIR) && pulumi up

.PHONY: references-stack
references-stack:
	cd $(REFERENCES_DEPLOY_DIR) && pulumi up

.PHONY: fmt
fmt:
	go fmt ./...
