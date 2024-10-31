swagger:
	swag init --parseDependency --parseInternal
wire:
	wire ./internal
logdy:
	tail -f logs/$(name).log | logdy --port=8081 --max-message-count=0