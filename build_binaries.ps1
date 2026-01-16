Remove-Item -ErrorAction SilentlyContinue -Path "./cmd/bin/StoicMigration.exe"
Remove-Item -ErrorAction SilentlyContinue -Path "./cmd/bin/StoicModelBuilder.exe"
Remove-Item -ErrorAction SilentlyContinue -Path "./cmd/bin/wgo.exe"

cd ./cmd/src/StoicMigration ; go build -o StoicMigration.exe ; move ./StoicMigration.exe ../../bin/StoicMigration.exe ; cd ../../..
cd ./cmd/src/StoicModelBuilder ; go build -o StoicModelBuilder.exe ; move ./StoicModelBuilder.exe ../../bin/StoicModelBuilder.exe ; cd ../../..
cd ./cmd/src/wgo-main ; go build -o wgo.exe ; move ./wgo.exe ../../bin/wgo.exe ; cd ../../..