GAMES_DEPLOY_DIR=deployments/games
REFERENCES_DEPLOY_DIR=deployments/references

.PHONY: games
games:
	cd $(GAMES_DEPLOY_DIR) && pulumi up

.PHONY: references
references:
	cd $(REFERENCES_DEPLOY_DIR) && pulumi up

.PHONY: fmt
fmt:
	go fmt ./...
