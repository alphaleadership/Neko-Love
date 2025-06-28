const {GOOGLE_IMG_SCRAP} = require('google-img-scrap');
const path = require('path');
const fs = require('fs');
const { exec } = require('child_process');
module.exports =async ({query,nomBaseQuery})=> {
    console.log(nomBaseQuery)
    console.log(`Recherche et enregistrement de ${query}...`);
    try {
      const results = await GOOGLE_IMG_SCRAP({
        search: ` ${query}  `,
        safe:true
      });
      // Création du fichier n.json pour enregistrer les résultats
     
  
      // Lecture du fichier n.json pour obtenir le nombre total de résultats
      let nombreTotal = 0;
  
  
      // Création de la base de données JSON si elle n'existe pas
   
      console.log(results.result.length)
      // Enregistrement des résultats
      for (const [index, result] of results.result.entries()) {
        //scan()
        const url = result.url;
        // Vérification si le lien existe déjà dans la base de données
  
        if (true) {
            if(!fs.existsSync(path.dirname(`./assets/${nomBaseQuery}`))){
                fs.mkdirSync(`./assets/${nomBaseQuery}`)
              }
          // Utilisation de path.basename pour extraire le nom de base de la query
          const filename = `./assets/${nomBaseQuery}/${fs.readdirSync(`./assets/${nomBaseQuery}`).length}.jpg`;
          const filepath = path.join(__dirname, filename);
          console.log(filepath)
         
          // Téléchargement de l'image en utilisant curl avec un délai pour éviter les limitations de taux
          const curlCommand = `curl -o ${filepath} "${url}"`;
          exec(curlCommand, (error, stdout, stderr) => {
            if (error) {
              //console.log(`Erreur lors du téléchargement de l'image: ${error}`);
              if(fs.existsSync(filepath)){
                fs.unlinkSync(filepath)
              }
            } else {
              // Ajout de l'entrée dans la base de données JSON
             /* baseDonnees.push({
                nomFichier: filename,
                ...result
              });
              fs.writeFileSync(cheminBaseDonnees, JSON.stringify(baseDonnees, null, 2));*/
            }
          });
          // Ajout d'un délai de 5 secondes entre chaque téléchargement pour éviter les limitations de taux
          await new Promise(resolve => setTimeout(resolve, 1000));
        } else {
          //    console.log(`Le lien ${url} existe déjà dans la base de données.`);
        }
      }
  
      // Écriture dans le fichier n.json le nombre total de résultats
    
      //fs.writeFileSync(cheminJsonFichier, newContent);
    } catch (error) {
      console.error(`Erreur lors de la recherche et de l'enregistrement: ${error}`);
    }
  }