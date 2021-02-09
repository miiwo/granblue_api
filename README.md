Uses scss for styling
uses webpack
uses babel (for transpiling between the different esx versions)
htmlwebpackplugin generates an html file and hooks it up with our index.js file
htmlloader is used to export html as a string to be read by webpack

run `yarn build` to get a production build made into the folder: `dist/`. Copy that over to my backend server so it gets served. If I want the backend to run with the frontend, run from the backend copy with the `dist` folder copied over.
