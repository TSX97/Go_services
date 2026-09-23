.PHONY: init


GREEN=\033[0;32m
YELLOW=\033[0;33m
CYAN=\033[0;36m
NC=\033[0m


init:
	@echo "\$(CYAN)==== Setup local environment ====\$(NC)"
	
	
	@chmod +x .git-config/hooks/pre-commit
	@chmod +x .git-config/hooks/commit-msg
	
	@git config --local include.path ../.git-config/aliases
	@echo "\$(GREEN)+ aliases are set up succesfuly\$(NC)"
	
	@git config --local core.hooksPath ./.git-config/hooks
	@echo "\$(GREEN)+ hooks are set up succesfuly\$(NC)"
	
	@echo "\$(YELLOW)Ready for development, test & use. Try 'docker compose up --build'\$(NC)"
	
	@echo "===-==-==-=- TODO -=-==-==-==="
	@cat .TODO
	
