migrateup:
	goose postgres postgres://fredel:lyfgoes1@localhost:5432/rssag up

sqlc:
	sqlc generate