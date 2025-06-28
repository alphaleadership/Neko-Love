const { exec } = require('child_process');
const path = require('path');
const fs = require('fs');
const Piscina = require('piscina');
const scan=require("./clean.js").dedup//()


const ps=["neko"]
//searchAndSave('votre_recherche'); // Remplacez 'votre_recherche' par votre recherche

const buildsearch=(base)=>{
    const resultatsCumules = [];
    for (let i = 0; i < (1 << ps.length); i++) {
        let cumule = `"${base}"`;
        for (let j = 0; j < ps.length; j++) {
            if ((i & (1 << j)) !== 0) {
                cumule += ` "${ps[j]}"`;
            }
        }
        resultatsCumules.push(cumule);
    }
    return [... new Set(resultatsCumules)];
}

const build=(b)=>{
    return {
        query:buildsearch(b,ps),
        base:b

    }
}

const buildpath=(query)=>{
  let chemin = query.replaceAll('"','')
  
  if (query.includes(' ')) {
    chemin = query.split(' ').join('_')
  }
  chemin =chemin
  const params = ps;
  if (params) {
    params.forEach(param => {
      const firstLetter = param.charAt(0).toLowerCase();
      chemin = chemin.replace(param.split(' ').join('_'), firstLetter+params.indexOf(param));
    });
  }
  return chemin;
}


const searchesArray = fs.readdirSync("./assets").map((dir)=>{return build(dir)});







const piscina = new Piscina({
  filename: path.resolve(__dirname, 'worker.js')
});
const main =async()=>{
    console.time('download')

    //searchesArray.push(...searches);
    for (const [index, item] of searchesArray.entries()) {
      try {
        scan()
       console.time(item)
       //console.log(item)
       
       await piscina.run({query:item.query, nomBaseQuery:item.base});
        console.timeEnd(item)
        // Ajout d'un délai de 1 minute entre chaque recherche pour éviter les limitations de taux
        //await new Promise(resolve => setTimeout(resolve, 60000)); // Augmentation du délai à 1 minute
      } catch (error) {
        console.error(`Erreur lors de la recherche et de l'enregistrement de ${item}: ${error}`);
      }
    }
    scan(["./assets"])  
    console.timeEnd("download")
    console.log('Toutes les recherches sont terminées.');
} 
  main()





