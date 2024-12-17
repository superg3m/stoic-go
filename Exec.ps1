# Requires ANSI color support in the terminal
[CmdletBinding()]
param (
    [string] $command,
    [string] $projectName = "",
    [bool] $useUiFrontDocker = $true,
    [bool] $useUiAdminDocker = $true,
    [bool] $useSmtpDocker = $true,
    [array] $commands = @(),
    [hashtable] $envVariables = @{},
    [bool] $isInteractive = $false,

    [string] $dbEngine = "",
    [string] $dbHost = "localhost",
    [string] $dbPort = "3306",
    [string] $dbUser = "root",
    [string] $dbName = "zsf"
)

$webContainerName = ""

$hasEnvFile = Test-Path -Path "./docker/.env"

$COLOR_GREEN = "`e[32m"
$COLOR_RED = "`e[31m"
$COLOR_BLUE = "`e[34m"
$COLOR_RESET = "`e[0m"

function GetDSN([string] $engine, [string] $host, [string] $port, [string] $dbName) {
    switch ($engine) {
        "mysql" {
            return "mysql:host=$host;port=$port;dbname=$dbName"
        }
        "pgsql" {
            return "pgsql:host=$host;port=$port;dbname=$dbName"
        }
        "sqlsrv" {
            return "sqlsrv:Server=$host,$port;Database=$dbName;Encrypt=0"
        }
        default {
            throw "$COLOR_RED Unsupported database engine: $engine $COLOR_RESET"
        }
    }
}

function PromptForYesNo([string] $promptMessage) {
    do {
        $response = (Read-Host -Prompt "$COLOR_BLUE $promptMessage (Y/N) $COLOR_RESET").Trim().ToUpper()
        if ($response -in @("Y", "N")) {
            return $response -eq "Y"
        }
        Write-Host "$COLOR_RED Invalid input. Please enter 'Y' or 'N'. $COLOR_RESET"
    } while ($true)
}

function PromptForProjectName() {
	$projectName = (Read-Host -Prompt "$COLOR_BLUE Project Name$COLOR_RESET").Trim().ToLower()
	return $projectName
}

function PromptForEngine() {
    do {
        $engine = (Read-Host -Prompt "$COLOR_BLUE Enter database engine (mysql/pgsql/sqlsrv) $COLOR_RESET").Trim().ToLower()
        if ($engine -in @("mysql", "pgsql", "sqlsrv")) {
            return $engine
        }
        Write-Host "$COLOR_RED Invalid input. Please enter 'mysql', 'pgsql', or 'sqlsrv'. $COLOR_RESET"
    } while ($true)
}

function DockerStart([string] $projectName, [string] $webContainer) {
    Push-Location ./docker
    docker compose -f docker-compose.yml -p $projectName up -d
    Pop-Location
    Write-Host "$COLOR_GREEN Docker containers have been started. $COLOR_RESET"
}

function DockerStop([string] $projectName) {
    Push-Location ./docker
    docker compose -f docker-compose.yml -p $projectName down
    Pop-Location
    Write-Host "$COLOR_GREEN Docker containers have been stopped. $COLOR_RESET"
}

function CreateComposeFile([string] $dbEngine, [bool] $useUiAdminDocker, [bool] $useUiFrontDocker, [bool] $useSmtpDocker) {
    Push-Location ./docker
    $composeFiles = @("-f docker-compose.yml")

    if ($useUiAdminDocker) {
        $composeFiles += "-f docker-compose-ui-admin.yml"
    }
    if ($useUiFrontDocker) {
        $composeFiles += "-f docker-compose-ui-front.yml"
    }
    if ($useSmtpDocker) {
        $composeFiles += "-f docker-compose-smtp.yml"
    }
	
    try {
        docker compose -f docker-compose.yml up -d
    } catch {
        Write-Host "$COLOR_RED An error occurred while starting Docker containers: $_ $COLOR_RESET"
        exit 1
    }

    Pop-Location
    Write-Host "Docker containers have been started for project: $projectName"
}

function InitializeProject([string] $projectName) {
    Write-Host "$COLOR_GREEN Initializing the project: $projectName $COLOR_RESET"

    if (-not $hasEnvFile) {
		$projectName = PromptForProjectName
        $useUiFrontDocker = PromptForYesNo "Enable UI Front Docker?"
        $useUiAdminDocker = PromptForYesNo "Enable UI Admin Docker?"
        $useSmtpDocker = PromptForYesNo "Enable SMTP Docker?"
		$dbEngine = PromptForEngine

        $envContent = @"
PROJECT_NAME=$projectName
UI_FRONT_DOCKER=$useUiFrontDocker
UI_ADMIN_DOCKER=$useUiAdminDocker
DB_ENGINE=$dbEngine
SMTP_DOCKER=$useSmtpDocker
"@
        $envContent | Out-File -FilePath "./docker/.env"
        Write-Host "$COLOR_GREEN .env file created successfully. $COLOR_RESET"
    }


    $webContainerName = "$projectName-web"

    CreateComposeFile -dbEngine $dbEngine -useUiAdminDocker $useUiAdminDocker -useUiFrontDocker $useUiFrontDocker -useSmtpDocker $useSmtpDocker
    DockerStart -projectName $projectName -webContainer $webContainerName
    Write-Host "$COLOR_GREEN Project initialization completed. $COLOR_RESET"
}

function MigrateDatabaseUp() {
    Write-Host "$COLOR_BLUE Applying database migrations... $COLOR_RESET"
    docker exec -it $webContainerName php vendor/bin/stoic-migrate up
    Write-Host "$COLOR_GREEN Database migrations applied. $COLOR_RESET"
}

function MigrateDatabaseDown() {
    Write-Host "$COLOR_BLUE Reverting database migrations... $COLOR_RESET"
    docker exec -it $webContainerName php vendor/bin/stoic-migrate down
    Write-Host "$COLOR_GREEN Database migrations reverted. $COLOR_RESET"
}

switch ($command) {
    "init" {
        InitializeProject -projectName $projectName
    }
    "down" {
        DockerStop -projectName $projectName
    }
    "update" {
        MigrateDatabaseUp
    }
    default {
        Write-Host "$COLOR_RED Unknown command: $command. Available commands: init, down, update. $COLOR_RESET"
    }
}

if ($dbEngine -ne "") {
    Write-Host "$COLOR_GREEN DSN: $(GetDSN -engine $dbEngine -host $dbHost -port $dbPort -dbName $dbName) $COLOR_RESET"
}
