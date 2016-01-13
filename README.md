Pour ajouter de nouveau constructeur au programme il faut :

	- Ajouter un nouveau constructeur dans le fichier AnalyzeConstructor.go( ex : Seabird = 0, autres = 1)
	- Ajouter dans le switch de choix constructeur un cas dans le fichier oceano2oceansites.go(ex :  case autres )
	- Ajouter dans ce nouveau cas la façon dont les fichiers seront décoder(voir le readSeabird.go)
	
Pour ajouter un nouveau type pour le constructeur Seabird au programme il faut :

	- Ajouter un nouveau Type et son expression régulière dans le fichier AnalyzeTypeSeabird.go( ex : CTD = 3, BTL = 5, Autres = 6)
	- Ajouter dans le switch de choix Type un cas dans le fichier readSeabird.go(ex :  case autres )
	- Ajouter dans ce nouveau cas la façon dont le fichier sera lu 
	- Ajouter les fichiers grammaires et config pour chaque nouveau type pour la lecture
	
Fichier grammaire :
Il se compose des expressions régulière propre au type et des différentes pass nécessaire à son décodage

Fichier Config :
Il permet la configuration de la structure Nc qui stocke les différentes données pour la création du fichier NetCDF