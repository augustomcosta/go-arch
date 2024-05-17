@echo off

cd ..

for /f "tokens=*" %%i in ('type .env') do set %%i

set MIGRATIONS_DIR=internal\infrastructure\sql\migration

goose -dir "%MIGRATIONS_DIR%" %*