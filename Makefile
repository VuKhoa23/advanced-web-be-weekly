swagger:
	swag init --parseDependency --parseInternal
wire:
	wire ./internal
logdy:
	tail -f $(name) | logdy --port=8081