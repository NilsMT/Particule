package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug, StatusDebug       bool
	//valeur par défauts
	DefaultSizeX, DefaultSizeY   float64
	DefaultOpacity               float64
	DefaultR, DefaultG, DefaultB float64
	DefaultSX, DefaultSY         float64
	DefaultRotation              float64
	//spawn
	RandomSpawn, MiddleSpawn bool
	InitNumParticles         int
	SpawnX, SpawnY           int
	SpawnRate                float64
	//comportement
	RandomColor, RandomSizeX, RandomSizeY, RandomOpacity, ContinuousRotation bool
	RandomSpeedX, RandomSpeedY, RandomRotation                               bool
	MaxSizeY, MaxSizeX, MaxSpeedY, MaxSpeedX, MaxRotation                    float64
	//extension
	Bounce, Gravity bool
	GravityForce    float64
}

var General Config
